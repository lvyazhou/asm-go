package repo

import (
	user_entity "asm_platform/domain/entity/user"
	"asm_platform/domain/repository"
	dao "asm_platform/infrastructure/pkg/database/mysql"
	"asm_platform/infrastructure/pkg/slog"
	"context"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepositoryDBT(ctx context.Context) *UserRepo {
	return &UserRepo{db: dao.DBT(ctx)}
}

func NewUserRepositoryDB() *UserRepo {
	return &UserRepo{db: dao.DB()}
}

var _ repository.UserRepository = &UserRepo{}

func (u *UserRepo) SaveUser(user *user_entity.User) (*user_entity.User, error) {
	err := u.db.Create(&user).Error
	if err != nil {
		slog.Errorf("user save error %v ", err.Error())
		return nil, err
	}
	return user, nil
}

func (u *UserRepo) GetUser(uid int64) (user *user_entity.User, err error) {
	if err = u.db.First(&user, uid).Error; err != nil {
		return nil, err
	}
	return user, err
}

// FindUserList 查询用户列表分页
func (u *UserRepo) FindUserList(userQuery *user_entity.UserQuery) ([]*user_entity.User, int64, error) {
	var query string
	var args []interface{}
	query += " u_status in (1,2)"
	// 用户名称
	if userQuery.Uname != "" && len(userQuery.Uname) > 0 {
		query += " and (u_name like ? or u_account like ?)"
		args = append(args, userQuery.Uname+"%")
		args = append(args, userQuery.Uname+"%")
	}

	// 用户邮箱
	if userQuery.Email != "" && len(userQuery.Email) > 0 {
		query += " and email like ?"
		args = append(args, userQuery.Email+"%")
	}

	query += " order by create_time desc"

	// 分页
	page := userQuery.Page
	size := userQuery.Size
	if size == 0 {
		size = 10
	}
	if page == 0 {
		page = 1
	}
	var totalCount int64
	var result []*user_entity.User
	err := dao.DB().
		Table("u_user").
		Select("id,u_account,email,u_name,u_mobile,u_pwd,u_status,create_time,create_user_id,update_time,update_user_id").
		Where(query, args...).
		Count(&totalCount).
		Limit(size).
		Offset((page - 1) * size).
		Scan(&result)

	if err.Error != nil {
		slog.Errorf("find sys user list %v error: %v\n", err)
		return nil, 0, err.Error
	}
	return result, totalCount, nil
}

func (u *UserRepo) GetUserByAccount(userQuery *user_entity.User) (user *user_entity.User, err error) {
	if err = u.db.Where("u_account = ?", userQuery.Account).Or("email = ?", userQuery.Email).Find(&user).Error; err != nil {
		return nil, err
	}
	return user, err
}
