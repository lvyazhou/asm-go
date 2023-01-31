package attack_app

import (
	"asm_platform/application/dto"
	"asm_platform/infrastructure/pkg/slog"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoConn struct {
}

func (m *MongoConn) Attack(info *dto.HostInfo, user string, pass string, timeout int64) (flag bool, err error) {
	flag = true
	var mongoUri string
	if user == "" && pass == "" {
		mongoUri = fmt.Sprintf("mongodb://%v:%v", info.Host, info.Port)
	} else {
		mongoUri = fmt.Sprintf("mongodb://%v:%v@%v:%v/?maxPoolSize=20&w=majority", user, pass, info.Host, info.Port)
	}
	// Create a new client and connect to the server
	c, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		flag = false
		slog.Errorf("mongo client error: %v", err.Error())
	}
	// Ping the primary
	if err := c.Ping(context.TODO(), readpref.Primary()); err != nil {
		flag = false
		slog.Error("Could not connect to MongoDB, please check configuration.")
	}
	return flag, nil
}
