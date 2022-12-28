package domain_entity

import (
	"asm_platform/application/vo"
	"asm_platform/domain/entity/asset"
	contansassetcode "asm_platform/infrastructure/pkg/constants/asset_code"
	utils_tool "asm_platform/infrastructure/pkg/tool/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
)

// Domain 域名资产
type Domain struct {
	// ID
	ID primitive.ObjectID `bson:"_id,omitempty"`

	// 资产ID
	Asset primitive.ObjectID `bson:"asset,omitempty"`

	// 域名名称
	Domain string `bson:"domain,omitempty"`

	// IP list
	Ip []string `bson:"ip,omitempty"`
}

// IsEmpty 判断是否为空
func (domain *Domain) IsEmpty() bool {
	if domain == nil {
		return false
	}
	return reflect.DeepEqual(domain, Domain{})
}

// DomainLookup 关联查询实体
type DomainLookup struct {
	Domain string             `bson:"domain"`
	Ip     []string           `bson:"ip"`
	Id     primitive.ObjectID `bson:"_id"`
	Asset  asset_entity.Asset `bson:"asset"`
}

// IsEmpty 判断是否为空
func (domain *DomainLookup) IsEmpty() bool {
	if domain == nil {
		return false
	}
	return reflect.DeepEqual(domain, Domain{})
}

// DomainToVo 实体转化为vo
func (domain *DomainLookup) DomainToVo() *vo.DomainVo {
	return &vo.DomainVo{
		ID:     domain.Id.Hex(),
		Domain: domain.Domain,
		IpList: domain.Ip,
		Asset: vo.AssetVo{
			ID:         domain.Asset.ID.Hex(),
			AssetNo:    domain.Asset.AssetNo,
			AssetName:  domain.Asset.AssetName,
			AssetType:  domain.Asset.AssetType,
			AssetLevel: contansassetcode.AssetLevel(domain.Asset.AssetLevel).String(),
			AssetGroup: domain.Asset.AssetGroup,
			Mgr:        domain.Asset.Mgr,
			Dept:       domain.Asset.Dept,
			CreateTime: utils_tool.FormatTimeToString(domain.Asset.CreateTime),
		},
	}
}
