package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/Portfolio-Advanced-software/BingeBuster-WatchHistoryService/globals"
	"github.com/Portfolio-Advanced-software/BingeBuster-WatchHistoryService/handlers"
	"github.com/Portfolio-Advanced-software/BingeBuster-WatchHistoryService/messaging"
	mongodb "github.com/Portfolio-Advanced-software/BingeBuster-WatchHistoryService/mongodb"
	historypb "github.com/Portfolio-Advanced-software/BingeBuster-WatchHistoryService/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50056"
)

func main() {
	// Configure 'log' package to give file name and line number on eg. log.Fatal
	// Pipe flags to one another (log.LstdFLags = log.Ldate | log.Ltime)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Starting server on port " + port + "...")

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
	srv := &handlers.HistoryServiceServer{}

	// Register the service with the server
	historypb.RegisterHistoryServiceServer(s, srv)

	// Initialize MongoDb client
	fmt.Println("Connecting to MongoDB...")
	globals.Db = mongodb.ConnectToMongoDB(globals.ConnURI)

	// Bind our collection to our global variable for use in other methods
	globals.HistoryDb = globals.Db.Database(globals.DbName).Collection(globals.CollectionName)

	// Start listening for messages RabbitMQ
	go messaging.ConsumeMessage("watch_history_queue", messaging.HandleMessage)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	fmt.Println("Server succesfully started on port " + port)

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
	globals.Db.Disconnect(globals.MongoCtx)
	fmt.Println("Done.")

}
