package repository

import (
	"asm_platform/domain/entity/domain"
)

type DomainRepository interface {
	// SaveDomain 保存域名
	SaveDomain(domain *domain_entity.Domain) error

	// FindDomainList 查找域名
	FindDomainList(query *domain_entity.DomainQuery) ([]*domain_entity.DomainLookup, error)
}
