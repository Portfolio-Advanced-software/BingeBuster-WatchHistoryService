package handlers

import (
	"context"
	"fmt"

	"github.com/Portfolio-Advanced-software/BingeBuster-WatchHistoryService/mongodb"
	historypb "github.com/Portfolio-Advanced-software/BingeBuster-WatchHistoryService/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *HistoryServiceServer) DeleteHistory(ctx context.Context, req *historypb.DeleteHistoryReq) (*historypb.DeleteHistoryRes, error) {
	// Get the ID (string) from the request message and convert it to an Object ID
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	// Check for errors
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}

	// Delete the history record by ID
	success, err := mongodb.DeleteHistoryByID(ctx, oid.Hex())
	// Check for errors
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find/delete history with id %s: %v", req.GetId(), err))
	}

	// Return response with success status
	return &historypb.DeleteHistoryRes{
		Success: success,
	}, nil
}
