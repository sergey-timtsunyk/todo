package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	authEventCollection = "auth_event"
)

type ConfigMongo struct {
	Host     string
	Port     string
	DBName   string
	User     string
	Password string
}

func NewConfigMongo(cfg ConfigMongo) (*mongo.Database, error) {
	ctxPing, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/", cfg.User, cfg.Password, cfg.Host, cfg.Port)
	client, err := mongo.Connect(ctxPing, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctxPing, nil)
	if err != nil {
		return nil, err
	}

	return client.Database(cfg.DBName), nil
}
