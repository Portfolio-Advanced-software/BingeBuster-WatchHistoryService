package handlers

import (
	"context"
	"fmt"

	"github.com/Portfolio-Advanced-software/BingeBuster-WatchHistoryService/globals"
	"github.com/Portfolio-Advanced-software/BingeBuster-WatchHistoryService/models"
	historypb "github.com/Portfolio-Advanced-software/BingeBuster-WatchHistoryService/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *HistoryServiceServer) ReadHistory(ctx context.Context, req *historypb.ReadHistoryReq) (*historypb.ReadHistoryRes, error) {
	// convert string id (from proto) to mongoDB ObjectId
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	result := globals.HistoryDb.FindOne(ctx, bson.M{"_id": oid})
	// Create an empty history to write our decode result to
	data := models.History{}
	// decode and write to data
	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find history with Object Id %s: %v", req.GetId(), err))
	}
	// Cast to ReadHistoryRes type
	response := &historypb.ReadHistoryRes{
		History: &historypb.History{
			Id:       oid.Hex(),
			Userid:   data.UserId,
			Movieid:  data.MovieId,
			Progress: data.Progress,
			Like:     data.Like,
		},
	}
	return response, nil
}
