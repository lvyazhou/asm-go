package asset_app

import (
	"asm_platform/application/dto"
	"asm_platform/application/vo"
	asset_entity "asm_platform/domain/entity/asset"
	constapicode "asm_platform/infrastructure/pkg/constants/api_code"
)

type AssetAppInterface interface {
	// SaveAsset 单个保存
	SaveAsset() constapicode.SocError

	// BatchSaveAsset 批量保存
	BatchSaveAsset() constapicode.SocError

	// GetAssetById 查询根据ID
	GetAssetById(id int64) (*asset_entity.Asset, constapicode.SocError)

	// DeleteAssetById 删除根据ID
	DeleteAssetById(id int64) constapicode.SocError

	// UpdateAsset 更新根据ID
	UpdateAsset() constapicode.SocError

	// FindAssetList 查询列表
	FindAssetList() ([]*asset_entity.Asset, constapicode.SocError)

	// FindAssetListByPage 查询分页
	FindAssetListByPage(query *dto.AssetQueryDto) ([]*vo.AssetVo, int64, constapicode.SocError)
}
