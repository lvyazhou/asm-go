package domain_app

import (
	asset_entity "asm_platform/domain/entity/asset"
	constapicode "asm_platform/infrastructure/pkg/constants/api_code"
	"asm_platform/infrastructure/repo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type DomainApp struct {
}

// NewDomainApp 实例化对象
func NewDomainApp() *DomainApp {
	return &DomainApp{}
}

// 实现接口
var _ DomainAppInterface = &DomainApp{}

// 实例化repo
var domainRepo = repo.NewDomainRepo()

// 实例化资产repo
var assetRepo = repo.NewAssetRepo()

func (d DomainApp) SaveDomain() constapicode.SocError {
	// 资产信息
	var assetId = primitive.NewObjectID()
	asset := &asset_entity.Asset{
		ID:         assetId,
		CreateUser: 11111111,
		CreateTime: time.Now(),
		UpdateUser: 22222222,
		UpdateTime: time.Now(),
		AssetName:  "111111111",
		AssetType:  "222222222",
		AssetLevel: 0,
		Mgr:        "333333333",
		Dept:       "444444444",
		AssetGroup: "555555555",
	}
	assetRepo.SaveAsset(asset)

	domain1 := &asset_entity.Domain{
		ID:     primitive.NewObjectID(),
		Asset:  assetId,
		Domain: "ztz.me",
		Ip:     []string{"1.1.1.1", "2.2.2.2"},
	}
	domainRepo.SaveDomain(domain1)

	domain2 := &asset_entity.Domain{
		ID:     primitive.NewObjectID(),
		Asset:  assetId,
		Domain: "ztz.md",
		Ip:     []string{"11.11.11.11", "22.22.22.22"},
	}
	domainRepo.SaveDomain(domain2)

	return constapicode.Success
}

func (d DomainApp) FindDomainList() constapicode.SocError {
	domainRepo.FindDomainList()
	
	return constapicode.Success
}
