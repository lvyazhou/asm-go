package repo

import (
	role_entity "asm_platform/domain/entity/role"
	"asm_platform/domain/repository"
	dao "asm_platform/infrastructure/pkg/database/mysql"
	"asm_platform/infrastructure/pkg/slog"
	"context"
	"gorm.io/gorm"
)

type RoleRepo struct {
	db *gorm.DB
}

func NewRoleRepositoryDBT(ctx context.Context) *RoleRepo {
	return &RoleRepo{db: dao.DBT(ctx)}
}

func NewRoleRepositoryDB() *RoleRepo {
	return &RoleRepo{db: dao.DB()}
}

var _ repository.RoleRepository = &RoleRepo{}

func (r RoleRepo) FindRoleListByUserId(userId int64) (result []*role_entity.Role, error error) {
	var query string
	var args []interface{}
	query = "a.r_status = 1 and b.user_id = ? "
	args = append(args, userId)

	err := dao.DB().
		Table("u_role as a").
		Select("DISTINCT a.id,a.r_name").
		Joins("LEFT JOIN u_user_role b ON a.id = b.role_id").
		Where(query, args...).
		Scan(&result)

	if err.Error != nil {
		slog.Errorf("find role list by user id %v error: %v\n", userId, err)
		return nil, err.Error
	}
	return result, nil
}
