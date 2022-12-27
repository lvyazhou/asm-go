package repo

import (
	"asm_platform/domain/entity/asset"
	"asm_platform/domain/repository"
	"asm_platform/infrastructure/pkg/database/mongo"
	"asm_platform/infrastructure/pkg/slog"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DomainRepo struct {
}

func NewDomainRepo() *DomainRepo {
	return &DomainRepo{}
}

// 实现接口
var _ repository.DomainRepository = &DomainRepo{}

// database name
var domainMgoDb = mgo.Mgo{
	Database:   "asm",
	Collection: "b_domain",
}

func (d DomainRepo) SaveDomain(domain *asset_entity.Domain) error {
	// 数据库链接
	coll := domainMgoDb.NewMgoCollection()

	result, err := coll.InsertOne(context.TODO(), domain)
	if err != nil {
		slog.Errorf("domain collection insert error: %v", err.Error())
		return err
	}
	id := result.InsertedID
	slog.Infof("domain collection insert id: %v", id)
	return nil
}

func (d DomainRepo) FindDomainList() error {
	// 数据库链接
	coll := domainMgoDb.NewMgoCollection()
	id, _ := primitive.ObjectIDFromHex("63aa5a6db09f564ed4881223")

	matchStage := bson.D{{"$match", bson.D{{"asset", id}}}}

	lookupStage := bson.D{{"$lookup", bson.D{{"from", "b_asset"}, {"localField", "asset"}, {"foreignField", "_id"}, {"as", "asset"}}}}
	unwindStage := bson.D{{"$unwind", bson.D{{"path", "$asset"}, {"preserveNullAndEmptyArrays", false}}}}

	showLoadedCursor, err := coll.Aggregate(context.TODO(), mongo.Pipeline{matchStage, lookupStage, unwindStage})
	if err != nil {
		panic(err)
	}
	var showsLoaded []bson.M
	if err = showLoadedCursor.All(context.TODO(), &showsLoaded); err != nil {
		panic(err)
	}
	fmt.Println(showsLoaded)
	return nil
}
