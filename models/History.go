package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type History struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	UserId   int64              `bson:"userid,omitempty"`
	MovieId  int64              `bson:"movieid,omitempty"`
	Progress string             `bson:"progress,omitempty"`
	Like     bool               `bson:"like,omitempty"`
}
