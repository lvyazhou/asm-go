package dto

import (
	asset_entity "asm_platform/domain/entity/asset"
	"time"
)

// AssetDto 资产DTO
type AssetDto struct {
	//资产标识
	AssetNo string `json:"asset_no"`
	//资产名称
	AssetName string `json:"asset_name"`
	// 资产类型
	AssetType string `json:"asset_type"`
	// 资产级别(高 中 低)
	AssetLevel int8 `json:"asset_level"`
	// 资产组
	AssetGroup string `json:"asset_group"`
	// 资产负责人
	Mgr string `json:"mgr"`
	// 资产部门
	Dept string `json:"dept"`
}

// AssetQueryDto 资产查询DTO
type AssetQueryDto struct {
	// 当前页码
	pageNo int64 `json:"page_no"`

	// 每页多少条
	PageSize int64 `json:"page_size"`

	// 资产名称
	AssetName string `json:"asset_name"`
}

// AssetDtoConvertEntity 资产dto转为实体
func (dto *AssetDto) AssetDtoConvertEntity(id, userId int64) *asset_entity.Asset {
	return &asset_entity.Asset{
		ID:         id,
		AssetNo:    dto.AssetNo,
		AssetName:  dto.AssetName,
		AssetType:  dto.AssetType,
		AssetLevel: dto.AssetLevel,
		Mgr:        dto.Mgr,
		Dept:       dto.Dept,
		AssetGroup: dto.AssetGroup,
		CreateUser: userId,
		CreateTime: time.Now(),
	}
}

// AssetQueryDtoConvertEntity 资产查询dto转化为实体
func (dto *AssetQueryDto) AssetQueryDtoConvertEntity() *asset_entity.AssetQuery {
	return &asset_entity.AssetQuery{
		AssetName: dto.AssetName,
		PageNo:    dto.pageNo,
		PageSize:  dto.PageSize,
	}
}
