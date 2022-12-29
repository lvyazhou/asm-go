package repo

import (
	"asm_platform/domain/entity/domain"
	"asm_platform/domain/repository"
	"asm_platform/infrastructure/pkg/database/mongo"
	"asm_platform/infrastructure/pkg/slog"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
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

func (d DomainRepo) SaveDomain(domain *domain_entity.Domain) error {
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

func (d DomainRepo) FindDomainList(query *domain_entity.DomainQuery) ([]*domain_entity.DomainLookup, error) {
	// 数据库链接
	coll := domainMgoDb.NewMgoCollection()
	// 查询条件
	//id, _ := primitive.ObjectIDFromHex("63aa5a6db09f564ed4881223")
	//matchStage2 := bson.D{{"$match", bson.D{{"asset", id}, {"domain", "ztz.me"}}}}
	//fmt.Println(matchStage2)

	bsonD := query.BuildQueryFilter()
	fmt.Println(bsonD)

	matchStage := bson.D{{"$match", bsonD}}
	fmt.Println(matchStage)
	// 分页
	pageNo := query.PageNo
	pageSize := query.PageSize
	pageNo = (pageNo - 1) * pageSize
	skipStage := bson.D{{"$skip", pageNo}}
	limitStage := bson.D{{"$limit", pageSize}}

	lookupStage := bson.D{{"$lookup", bson.D{{"from", "b_asset"}, {"localField", "asset"}, {"foreignField", "_id"}, {"as", "asset"}}}}
	unwindStage := bson.D{{"$unwind", bson.D{{"path", "$asset"}, {"preserveNullAndEmptyArrays", false}}}}

	showLoadedCursor, err := coll.Aggregate(context.TODO(), mongo.Pipeline{matchStage, skipStage, limitStage, lookupStage, unwindStage})
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	var showsLoaded []bson.M
	if err = showLoadedCursor.All(context.TODO(), &showsLoaded); err != nil {
		panic(err)
	}

	var ds []*domain_entity.DomainLookup
	//fmt.Println(showsLoaded)
	for _, v := range showsLoaded {
		//fmt.Println(k)
		//fmt.Println(v)
		// convert m to s
		var s *domain_entity.DomainLookup
		bsonBytes, _ := bson.Marshal(v)
		err = bson.Unmarshal(bsonBytes, &s)
		if err != nil {
			fmt.Println(err.Error())
		}
		//fmt.Println(s)
		//fmt.Println(s.Id.Hex())
		ds = append(ds, s)
	}
	//fmt.Println(ds)
	return ds, nil
}
