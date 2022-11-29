package dao

import (
	"asm_platform/infrastructure/config"
	constapicode "asm_platform/infrastructure/pkg/constants/api_code"
	"asm_platform/infrastructure/pkg/slog"
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var (
	// db 将此变量私有，使用时调用 DB() or DBT() 函数获取
	db *gorm.DB
)

type Writer struct {
}

func (w Writer) Printf(format string, args ...interface{}) {
	slog.Debugf(format, args...)
}

// Init mysql init
func Init() {
	if err := NewMysqlDriver(); err != nil {
		slog.Panicf("mysql error on database initialization: %s\n", err)
		return
	}
}

func Close() error {
	sqlDB, _ := db.DB()
	return sqlDB.Close()
}

func NewMysqlDriver() error {
	appConfig := config.GetConfig()
	user := appConfig.GetString("mysql.user")
	password := appConfig.GetString("mysql.password")
	host := appConfig.GetString("mysql.host")
	port := appConfig.GetInt("mysql.port")
	dbName := appConfig.GetString("mysql.db")
	maxIdleConn := appConfig.GetInt("mysql.maxIdleConn")
	maxOpenConn := appConfig.GetInt("mysql.maxOpenConn")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbName)
	mysqlDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}

	sqlDB, _ := mysqlDB.DB()
	sqlDB.SetConnMaxLifetime(10 * time.Minute)
	sqlDB.SetMaxOpenConns(maxOpenConn)
	sqlDB.SetMaxIdleConns(maxIdleConn)
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	db = mysqlDB
	return nil
}

// KEY 开启事务时，存在context中的值
const KEY = "Context_DB_Transaction"

// DB 获取无事务DB
func DB() *gorm.DB {
	return db
}

// DBT 通过传递的Context获取DB，所有执行SQL务必使用这个函数获取DB
func DBT(ctx context.Context) *gorm.DB {
	if ctx != nil {
		//获取上下文
		tx := ctx.Value(KEY)
		if tx != nil {
			//如果不为空，已开启事务，返回事务DB
			return tx.(*gorm.DB)
		}
	}
	//如果为空，说明没有开启事务，默认以非事务执行
	return db
}

// Begin 开启事务，事务一旦开始，必须使用 tx 处理数据!
func Begin(ctx context.Context) context.Context {
	tx := db.Begin()
	//放入上下文
	return context.WithValue(ctx, KEY, tx)
}

// Commit 提交事务
func Commit(ctx context.Context) (code constapicode.SocError) {
	tx := DBT(ctx)
	err := tx.Commit()
	if err != nil {
		return constapicode.DBTxCommitError
	}
	return constapicode.Success
}

// Rollback 回滚事务
func Rollback(ctx context.Context) (code constapicode.SocError) {
	tx := DBT(ctx)
	err := tx.Rollback()
	if err.Error != nil {
		fmt.Println("事务回滚失败", err.Error)
		return constapicode.DBTxRollbackError
	}
	return constapicode.Success
}

// SavePoint 保存点，可以使用 RollbackTo() 回滚到此位置再提交
func SavePoint(ctx context.Context, pointName string) (code constapicode.SocError) {
	tx := DBT(ctx)
	point := tx.SavePoint(pointName)
	if point.Error != nil {
		fmt.Println("事务保存点失败", point.Error)
		return constapicode.DBTxSavePointError
	}

	return constapicode.Success
}

// RollbackTo 回滚至某个保存点
func RollbackTo(ctx context.Context, pointName string) (code constapicode.SocError) {
	tx := DBT(ctx)
	rollback := tx.RollbackTo(pointName)
	if rollback.Error != nil {
		fmt.Println("事务回滚至保存点失败", rollback.Error)
		return constapicode.DBTxRollbackToError
	}
	return constapicode.Success
}
