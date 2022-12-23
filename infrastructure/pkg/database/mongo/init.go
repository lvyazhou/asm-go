package mgo

import (
	"asm_platform/infrastructure/config"
	"asm_platform/infrastructure/pkg/slog"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

// Init 初始化Mongo链接
func Init() {
	if err := newMgoDriver(); err != nil {
		slog.Panicf("mongo error on database initialization: %s\n", err)
		return
	}
}

// 建立mgo驱动
func newMgoDriver() error {
	conf := config.GetConfig()
	uri := conf.GetString("mgo.local.uri")

	// Create a new client and connect to the server
	c, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		slog.Errorf("mongo client error: %v", err.Error())
		return err
	}
	// Ping the primary
	if err := c.Ping(context.TODO(), readpref.Primary()); err != nil {
		slog.Error("Could not connect to MongoDB, please check configuration.")
		return err
	}
	client = c
	slog.Infof("mongo Connected...")
	return nil
}

type Mgo struct {
	Database   string `json:"db"`
	Collection string `json:"collection"`
}

// NewMgoCollection 返回mog collection实例
func (mgo *Mgo) NewMgoCollection() *mongo.Collection {
	return client.Database(mgo.Database).Collection(mgo.Collection)
}
