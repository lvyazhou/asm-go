package attack_app

import (
	"asm_platform/application/dto"
	"database/sql"
	"fmt"
	"time"
)

type OracleConn struct {
}

func (o *OracleConn) Attack(info *dto.HostInfo, user string, pass string, timeout int64) (flag bool, err error) {
	flag = false
	Host, Port, Username, Password := info.Host, info.Port, user, pass
	dataSourceName := fmt.Sprintf("oracle://%s:%s@%s:%s/orcl", Username, Password, Host, Port)
	db, err := sql.Open("oracle", dataSourceName)
	if err == nil {
		db.SetConnMaxLifetime(time.Duration(timeout) * time.Second)
		db.SetConnMaxIdleTime(time.Duration(timeout) * time.Second)
		db.SetMaxIdleConns(0)
		defer db.Close()
		err = db.Ping()
		if err == nil {
			flag = true
		}
	}
	return flag, err
}
