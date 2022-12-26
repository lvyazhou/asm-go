package asset_app

import (
	"asm_platform/application/dto"
	"asm_platform/application/vo"
	asset_entity "asm_platform/domain/entity/asset"
	constapicode "asm_platform/infrastructure/pkg/constants/api_code"
	"asm_platform/infrastructure/repo"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		ID:         primitive.NewObjectID(),
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
			ID:         primitive.NewObjectID(),
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

func (a AssetApp) FindAssetListByPage(query *dto.AssetQueryDto) ([]*vo.AssetVo, int64, constapicode.SocError) {
	assetEntity := query.AssetQueryDtoConvertEntity()
	results, count, err := assetRepo.FindAssetListByPage(assetEntity)
	if err != nil {
		return nil, 0, constapicode.DocumentNotFind
	}
	var assetList []*vo.AssetVo
	if len(results) > 0 {
		// 实例化user repo
		d := repo.NewUserRepositoryDB()
		for a := range results {
			asset := results[a]
			// convert to vo
			assetVo := asset.AssetToVo()
			// 查询用户创建人员
			cuser, cerr := d.GetUser(asset.CreateUser)
			if !cuser.IsEmpty() && cerr == nil {
				assetVo.CreateUser = cuser.Name
			}
			assetList = append(assetList, assetVo)
		}
	}
	return assetList, count, constapicode.Success
}
