package vo

import "reflect"

// DomainVo 域名vo
type DomainVo struct {
	// id
	ID string `json:"id"`

	// domain
	Domain string `json:"domain"`

	// ip list
	IpList []string `json:"ip_list"`

	// 关联资产
	Asset AssetVo `json:"asset"`
}

func (vo *DomainVo) IsEmpty() bool {
	if vo == nil {
		return false
	}
	return reflect.DeepEqual(vo, DomainVo{})
}
