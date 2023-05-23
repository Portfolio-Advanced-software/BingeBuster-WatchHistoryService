package models

type History struct {
	UserId   int64  `bson:"userid,omitempty"`
	MovieId  int64  `bson:"movieid,omitempty"`
	Progress string `bson:"progress,omitempty"`
	Like     bool   `bson:"like,omitempty"`
}
