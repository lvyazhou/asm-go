package asset_app

import (
	asset_entity "asm_platform/domain/entity/asset"
	constapicode "asm_platform/infrastructure/pkg/constants/api_code"
	utils_tool "asm_platform/infrastructure/pkg/tool/utils"
	"asm_platform/infrastructure/repo"
	"fmt"
	"strconv"
	"time"
)

type AssetApp struct {
}

// NewAssetApp 实例化对象
func NewAssetApp() *AssetApp {
	return &AssetApp{}
}

// 实现接口
var _ AssetAppInterface = &AssetApp{}

// 实例化repo
var assetRepo = repo.NewAssetRepo()

func (a AssetApp) SaveAsset() constapicode.SocError {
	asset := &asset_entity.Asset{
		ID:         utils_tool.GenerateUniqueId(),
		CreateUser: 11,
		CreateTime: time.Now(),
		UpdateUser: 22,
		UpdateTime: time.Now(),
		AssetName:  "1",
		AssetType:  "2",
		AssetLevel: 0,
		Mgr:        "3",
		Dept:       "4",
		AssetGroup: "5",
	}
	assetRepo.SaveAsset(asset)
	return constapicode.Success
}

func (a AssetApp) BatchSaveAsset() constapicode.SocError {
	var assetList []*asset_entity.Asset
	for i := 0; i < 10; i++ {
		asset := &asset_entity.Asset{
			ID:         utils_tool.GenerateUniqueId(),
			CreateUser: 11,
			CreateTime: time.Now(),
			UpdateUser: 22,
			UpdateTime: time.Now(),
			AssetName:  "1_" + strconv.Itoa(i),
			AssetType:  "2",
			AssetLevel: 0,
			Mgr:        "3",
			Dept:       "4",
			AssetGroup: "5",
		}
		assetList = append(assetList, asset)
	}
	assetRepo.BatchSaveAssetList(assetList)
	return constapicode.Success
}

func (a AssetApp) GetAssetById(id int64) (*asset_entity.Asset, constapicode.SocError) {
	asset, err := assetRepo.GetAssetById(id)
	if err != nil {
		return nil, constapicode.DocumentNotFind
	}
	return asset, constapicode.Success
}

func (a AssetApp) DeleteAssetById(id int64) constapicode.SocError {
	//TODO implement me
	panic("implement me")
}

func (a AssetApp) UpdateAsset() constapicode.SocError {
	//TODO implement me
	panic("implement me")
}

func (a AssetApp) FindAssetList() ([]*asset_entity.Asset, constapicode.SocError) {
	assetList, err := assetRepo.FindAssetList(&asset_entity.AssetQuery{})
	if err != nil {
		return nil, constapicode.DocumentNotFind
	}
	return assetList, constapicode.Success
}

func (a AssetApp) FindAssetListByPage() ([]*asset_entity.Asset, int64, constapicode.SocError) {
	assetList, count, err := assetRepo.FindAssetListByPage(&asset_entity.AssetQuery{
		AssetName: "",
		PageNo:    0,
		PageSize:  5,
	})
	if err != nil {
		return nil, 0, constapicode.DocumentNotFind
	}
	fmt.Println("asset list len ", len(assetList))
	fmt.Println("asset count ", count)
	return assetList, count, constapicode.Success
}
