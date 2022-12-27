package domain_app

import "asm_platform/infrastructure/pkg/constants/api_code"

type DomainAppInterface interface {
	// SaveDomain 单个保存
	SaveDomain() constapicode.SocError

	// FindDomainList 查询domain list
	FindDomainList() constapicode.SocError
}
