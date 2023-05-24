package handlers

import (
	"context"
	"fmt"

	"github.com/Portfolio-Advanced-software/BingeBuster-WatchHistoryService/globals"
	"github.com/Portfolio-Advanced-software/BingeBuster-WatchHistoryService/models"
	historypb "github.com/Portfolio-Advanced-software/BingeBuster-WatchHistoryService/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *HistoryServiceServer) UpdateHistory(ctx context.Context, req *historypb.UpdateHistoryReq) (*historypb.UpdateHistoryRes, error) {
	// Get the history data from the request
	history := req.GetHistory()

	// Convert the Id string to a MongoDB ObjectId
	oid, err := primitive.ObjectIDFromHex(history.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Could not convert the supplied history id to a MongoDB ObjectId: %v", err),
		)
	}

	// Convert the data to be updated into an unordered Bson document
	update := bson.M{
		"userid":   history.GetUserid(),
		"movieid":  history.GetMovieid(),
		"progress": history.GetProgress(),
		"like":     history.GetLike(),
	}

	// Convert the oid into an unordered bson document to search by id
	filter := bson.M{"_id": oid}

	// Result is the BSON encoded result
	// To return the updated document instead of original we have to add options.
	result := globals.HistoryDb.FindOneAndUpdate(ctx, filter, bson.M{"$set": update}, options.FindOneAndUpdate().SetReturnDocument(1))

	// Decode result and write it to 'decoded'
	decoded := models.History{}
	err = result.Decode(&decoded)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Could not find history with supplied ID: %v", err),
		)
	}
	return &historypb.UpdateHistoryRes{
		History: &historypb.History{
			Id:       decoded.ID.Hex(),
			Userid:   decoded.UserId,
			Movieid:  decoded.MovieId,
			Progress: decoded.Progress,
			Like:     decoded.Like,
		},
	}, nil
}
