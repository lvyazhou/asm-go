package repository

import "asm_platform/domain/entity/asset"

type DomainRepository interface {
	// SaveDomain 保存域名
	SaveDomain(domain *asset_entity.Domain) error

	// FindDomainList 查找域名
	FindDomainList() error
}
