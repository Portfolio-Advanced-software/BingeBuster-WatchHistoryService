package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type History struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	UserId   string             `bson:"userid,omitempty"`
	MovieId  string             `bson:"movieid,omitempty"`
	Progress string             `bson:"progress,omitempty"`
	Like     bool               `bson:"like,omitempty"`
}
