package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	singletonDBClient *mongo.Client
)

func GetDBClient() (*mongo.Client, error) {
	if singletonDBClient == nil {
		println("DBClient has not been init, please init it.")
		return nil, errors.New("DBClient has not been init")
	}
	return singletonDBClient, nil
}

func InitDBClient(ctx context.Context, opt options.ClientOptions) error {
	if singletonDBClient == nil {
		var err error = nil
		singletonDBClient, err = mongo.Connect(ctx, &opt)
		if err != nil {
			fmt.Printf("connect failed, err = %s", err.Error())
			return err
		} else if err = testDBConnect(); err != nil {
			println("test DB connect failed, err = %s", err.Error())
			return err
		}
	}
	println("init DB success")
	return nil
}

func testDBConnect() error {
	if singletonDBClient != nil {
		if err := singletonDBClient.Ping(context.TODO(), readpref.Primary()); err != nil {
			return err
		}
	}
	return nil
}

func DisconnectDBClient() error {
	if singletonDBClient == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	if err := singletonDBClient.Disconnect(ctx); err != nil {
		return err
	}
	return nil
}
