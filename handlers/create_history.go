package handlers

import (
	"context"
	"fmt"

	"github.com/Portfolio-Advanced-software/BingeBuster-WatchHistoryService/globals"
	"github.com/Portfolio-Advanced-software/BingeBuster-WatchHistoryService/models"
	historypb "github.com/Portfolio-Advanced-software/BingeBuster-WatchHistoryService/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *HistoryServiceServer) CreateHistory(ctx context.Context, req *historypb.CreateHistoryReq) (*historypb.CreateHistoryRes, error) {
	// Essentially doing req.History to access the struct with a nil check
	history := req.GetHistory()
	if history == nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid history")
	}
	// Now we have to convert this into a History type to convert into BSON
	data := models.History{
		// ID:    Empty, so it gets omitted and MongoDB generates a unique Object ID upon insertion.
		UserId:   history.GetUserid(),
		MovieId:  history.GetMovieid(),
		Progress: history.GetProgress(),
		Like:     history.GetLike(),
	}

	// Insert the data into the database, result contains the newly generated Object ID for the new document
	result, err := globals.HistoryDb.InsertOne(globals.MongoCtx, data)
	// check for potential errors
	if err != nil {
		// return internal gRPC error to be handled later
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	// add the id to history, first cast the "generic type" (go doesn't have real generics yet) to an Object ID.
	oid := result.InsertedID.(primitive.ObjectID)
	// Convert the object id to it's string counterpart
	history.Id = oid.Hex()
	// return the blog in a CreateHistoryRes type
	return &historypb.CreateHistoryRes{History: history}, nil
}
