package repository

import "asm_platform/domain/entity/asset"

type AssetRepository interface {

	// SaveAsset 资产保存
	SaveAsset(asset *asset_entity.Asset) error

	// BatchSaveAssetList 批量资产保存
	BatchSaveAssetList(assetList []*asset_entity.Asset) error

	// UpdateAsset 资产更新
	UpdateAsset(asset *asset_entity.Asset) error

	// DeleteAssetById 资产删除
	DeleteAssetById(id int64) error

	// GetAssetById 根据ID查找资产
	GetAssetById(id int64) (*asset_entity.Asset, error)

	// FindAssetList 查询资产
	FindAssetList(assetQuery *asset_entity.AssetQuery) ([]*asset_entity.Asset, error)

	// FindAssetListByPage 查询资产分页
	FindAssetListByPage(assetQuery *asset_entity.AssetQuery) ([]*asset_entity.Asset, int64, error)
}
