package data

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	AuthorizationEvent  = "authorization_event"
	AuthenticationEvent = "authentication_event"
	VerificationEvent   = "verification_event"
)

type AuthEvent struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	UserId     uint               `bson:"user_id"`
	EventName  string             `bson:"event_name"`
	UriRequest string             `bson:"uri_request"`
	CreatedAt  time.Time          `bson:"created_at"`
}
