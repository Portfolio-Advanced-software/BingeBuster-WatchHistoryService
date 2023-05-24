package globals

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	Db             *mongo.Client
	HistoryDb      *mongo.Collection
	MongoCtx       context.Context
	mongoUsername  = "user-service"
	mongoPwd       = "vLxxhmS0eJFwmteF"
	ConnURI        = "mongodb+srv://" + mongoUsername + ":" + mongoPwd + "@cluster0.fpedw5d.mongodb.net/"
	DbName         = "WatchHistoryService"
	CollectionName = "Histories"
)
