package repo

import (
	asset_entity "asm_platform/domain/entity/asset"
	"asm_platform/domain/repository"
	mgo "asm_platform/infrastructure/pkg/database/mongo"
	"asm_platform/infrastructure/pkg/slog"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type AssetRepo struct {
}

func NewAssetRepo() *AssetRepo {
	return &AssetRepo{}
}

// 实现接口
var _ repository.AssetRepository = &AssetRepo{}

// database name
var mgoDb = mgo.Mgo{
	Database:   "asm",
	Collection: "b_asset",
}

func (a AssetRepo) SaveAsset(asset *asset_entity.Asset) error {

	// collection
	var assetCollection = mgoDb.NewMgoCollection()
	result, err := assetCollection.InsertOne(context.TODO(), asset)
	if err != nil {
		slog.Errorf("asset collection insert error: %v", err.Error())
		return err
	}
	id := result.InsertedID
	slog.Infof("asset collection insert id: %v", id)
	return nil
}

func (a AssetRepo) BatchSaveAssetList(assetList []*asset_entity.Asset) error {
	// collection
	var assetCollection = mgoDb.NewMgoCollection()
	var newAssetList []interface{}
	if len(assetList) > 0 {
		for a := range assetList {
			newAssetList = append(newAssetList, assetList[a])
		}
	}
	result, err := assetCollection.InsertMany(context.TODO(), newAssetList)
	if err != nil {
		slog.Errorf("asset collection batch insert error: %v", err.Error())
		return err
	}
	slog.Infof("asset collection batch insert id list: %v", result.InsertedIDs)
	return nil
}

func (a AssetRepo) UpdateAsset(asset *asset_entity.Asset) error {
	var coll = mgoDb.NewMgoCollection()
	filter := bson.D{{"id", asset.ID}}
	// 对象转换为bson d
	update := bson.D{{"$set", bson.D{{"asset_name", "lyz"}, {"asset_type", 4}}}}

	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		slog.Errorf("asset collection update error: %v", err.Error())
		return err
	}
	slog.Infof("asset collection update id list: %v", result.UpsertedID)
	return nil
}

func (a AssetRepo) DeleteAssetById(id int64) error {
	var coll = mgoDb.NewMgoCollection()
	filter := bson.D{{"id", id}}
	result, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		slog.Errorf("asset collection delete error: %v", err.Error())
		return err
	}
	slog.Infof("asset collection delete id list: %v", result.DeletedCount)
	return nil
}

func (a AssetRepo) GetAssetById(id int64) (asset *asset_entity.Asset, err error) {
	// collection
	var assetCollection = mgoDb.NewMgoCollection()
	filter := bson.D{{"id", id}}
	err = assetCollection.FindOne(context.TODO(), filter).Decode(&asset)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, err
	}
	return asset, err
}

func (a AssetRepo) FindAssetList(assetQuery *asset_entity.AssetQuery) (assetList []*asset_entity.Asset, err error) {
	var coll = mgoDb.NewMgoCollection()
	filter := bson.D{{"asset_type", "2"}}
	var opts = options.Find()
	opts.SetSort(bson.D{{"id", -1}})
	cursor, err := coll.Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var asset *asset_entity.Asset
		if err := cursor.Decode(&asset); err != nil {
			return nil, err
		}
		assetList = append(assetList, asset)
	}
	return assetList, err
}

func (a AssetRepo) FindAssetListByPage(assetQuery *asset_entity.AssetQuery) ([]*asset_entity.Asset, int64, error) {
	var coll = mgoDb.NewMgoCollection()
	filter := bson.D{}
	// 计算count
	count, err := coll.CountDocuments(context.TODO(), filter)
	if err != nil {
		slog.Errorf("asset collection find count error : %v", err)
		return nil, 0, err
	}
	// 分页查询
	pageNo := (assetQuery.PageNo - 1) * assetQuery.PageSize // 当前页数
	var opts = &options.FindOptions{
		Skip:  &pageNo,
		Limit: &assetQuery.PageSize,
		Sort:  bson.D{{"id", -1}}, // 1 升序 -1 降序
	}

	cursor, err := coll.Find(context.TODO(), filter, opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	var assetList []*asset_entity.Asset
	for cursor.Next(context.TODO()) {
		var asset *asset_entity.Asset
		if err := cursor.Decode(&asset); err != nil {
			return nil, 0, err
		}
		assetList = append(assetList, asset)
	}
	return assetList, count, err
}
