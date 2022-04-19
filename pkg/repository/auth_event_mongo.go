package repository

import (
	"context"
	"github.com/sergey-timtsunyk/todo/pkg/data"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type AuthEventMongo struct {
	db *mongo.Database
}

func NewAuthEventMongo(db *mongo.Database) *AuthEventMongo {
	return &AuthEventMongo{db: db}
}

func (r *AuthEventMongo) Create(userId uint, uriRequest string, eventName string) error {
	collection := r.db.Collection(authEventCollection)
	_, err := collection.InsertOne(context.TODO(), data.AuthEvent{
		ID:         primitive.NewObjectID(),
		UserId:     userId,
		EventName:  eventName,
		UriRequest: uriRequest,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	})

	return err
}
