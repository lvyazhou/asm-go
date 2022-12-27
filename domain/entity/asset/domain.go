package asset_entity

import "go.mongodb.org/mongo-driver/bson/primitive"

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
