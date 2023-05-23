package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	models "github.com/Portfolio-Advanced-software/BingeBuster-WatchHistoryService/models"
	mongodb "github.com/Portfolio-Advanced-software/BingeBuster-WatchHistoryService/mongodb"
	historypb "github.com/Portfolio-Advanced-software/BingeBuster-WatchHistoryService/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HistoryServiceServer struct {
	historypb.UnimplementedHistoryServiceServer
}

func (s *HistoryServiceServer) CreateHistory(ctx context.Context, req *historypb.CreateHistoryReq) (*historypb.CreateHistoryRes, error) {
	// Essentially doing req.History to access the struct with a nil check
	history := req.GetHistory()
	if history == nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid history")
	}
	// Now we have to convert this into a History type to convert into BSON
	data := models.History{
		// ID:    Empty, so it gets omitted and MongoDB generates a unique Object ID upon insertion.
		UserId:   history.GetUserId(),
		MovieId:  history.GetMovieId(),
		Progress: history.GetProgress(),
		Like:     history.GetLike(),
	}

	// Insert the data into the database, result contains the newly generated Object ID for the new document
	result, err := historydb.InsertOne(mongoCtx, data)
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

func (s *HistoryServiceServer) ReadHistory(ctx context.Context, req *historypb.ReadHistoryReq) (*historypb.ReadHistoryRes, error) {
	// convert string id (from proto) to mongoDB ObjectId
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	result := historydb.FindOne(ctx, bson.M{"_id": oid})
	// Create an empty history to write our decode result to
	data := models.History{}
	// decode and write to data
	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find history with Object Id %s: %v", req.GetId(), err))
	}
	// Cast to ReadHistoryRes type
	response := &historypb.ReadHistoryRes{
		History: &historypb.History{
			UserId:   data.UserId,
			MovieId:  data.MovieId,
			Progress: data.Progress,
			Like:     data.Like,
		},
	}
	return response, nil
}

func (s *HistoryServiceServer) ListHistories(req *historypb.ListHistoriesReq, stream historypb.HistoryService_ListHistoriesServer) error {
	// Initiate a history type to write decoded data to
	data := &models.History{}
	// collection.Find returns a cursor for our (empty) query
	cursor, err := historydb.Find(context.Background(), bson.M{})
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	// An expression with defer will be called at the end of the function
	defer cursor.Close(context.Background())
	// cursor.Next() returns a boolean, if false there are no more items and loop will break
	for cursor.Next(context.Background()) {
		// Decode the data at the current pointer and write it to data
		err := cursor.Decode(data)
		// check error
		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}
		// If no error is found send blog over stream
		stream.Send(&historypb.ListHistorysRes{
			History: &historypb.History{
				UserId:   data.UserId,
				MovieId:  data.MovieId,
				Progress: data.Progress,
				Like:     data.Like,
			},
		})
	}
	// Check if the cursor has any errors
	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unkown cursor error: %v", err))
	}
	return nil
}

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
		"userid":   history.GetUserId(),
		"movieId":  history.GetMovieId(),
		"progress": history.GetProgress(),
		"like":     history.GetLike(),
	}

	// Convert the oid into an unordered bson document to search by id
	filter := bson.M{"_id": oid}

	// Result is the BSON encoded result
	// To return the updated document instead of original we have to add options.
	result := historydb.FindOneAndUpdate(ctx, filter, bson.M{"$set": update}, options.FindOneAndUpdate().SetReturnDocument(1))

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
			UserId:   decoded.UserId,
			MovieId:  decoded.MovieId,
			Progress: decoded.Progress,
			Like:     decoded.Like,
		},
	}, nil
}

func (s *HistoryServiceServer) DeleteHistory(ctx context.Context, req *historypb.DeleteHistoryReq) (*historypb.DeleteHistoryRes, error) {
	// Get the ID (string) from the request message and convert it to an Object ID
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	// Check for errors
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	// DeleteOne returns DeleteResult which is a struct containing the amount of deleted docs (in this case only 1 always)
	// So we return a boolean instead
	_, err = historydb.DeleteOne(ctx, bson.M{"_id": oid})
	// Check for errors
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find/delete history with id %s: %v", req.GetId(), err))
	}
	// Return response with success: true if no error is thrown (and thus document is removed)
	return &historypb.DeleteHistoryRes{
		Success: true,
	}, nil
}

const (
	port = ":50056"
)

var db *mongo.Client
var historydb *mongo.Collection
var mongoCtx context.Context

var mongoUsername = "user-service"
var mongoPwd = "vLxxhmS0eJFwmteF"
var connUri = "mongodb+srv://" + mongoUsername + ":" + mongoPwd + "@cluster0.fpedw5d.mongodb.net/"

var dbName = "HistoryService"
var collectionName = "Historys"

func main() {
	// Configure 'log' package to give file name and line number on eg. log.Fatal
	// Pipe flags to one another (log.LstdFLags = log.Ldate | log.Ltime)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Starting server on port :50056...")

	// Set listener to start server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Unable to listen on port %p: %v", lis.Addr(), err)
	}

	// Set options, here we can configure things like TLS support
	opts := []grpc.ServerOption{}
	// Create new gRPC server with (blank) options
	s := grpc.NewServer(opts...)
	// Create HistoryService type
	srv := &HistoryServiceServer{}

	// Register the service with the server
	historypb.RegisterHistoryServiceServer(s, srv)

	// Initialize MongoDb client
	fmt.Println("Connecting to MongoDB...")
	db = mongodb.ConnectToMongoDB(connUri)

	// Bind our collection to our global variable for use in other methods
	historydb = db.Database(dbName).Collection(collectionName)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	fmt.Println("Server succesfully started on port :50056")

	// Right way to stop the server using a SHUTDOWN HOOK
	// Create a channel to receive OS signals
	c := make(chan os.Signal)

	// Relay os.Interrupt to our channel (os.Interrupt = CTRL+C)
	// Ignore other incoming signals
	signal.Notify(c, os.Interrupt)

	// Block main routine until a signal is received
	// As long as user doesn't press CTRL+C a message is not passed and our main routine keeps running
	<-c

	// After receiving CTRL+C Properly stop the server
	fmt.Println("\nStopping the server...")
	s.Stop()
	lis.Close()
	fmt.Println("Closing MongoDB connection")
	db.Disconnect(mongoCtx)
	fmt.Println("Done.")

}
