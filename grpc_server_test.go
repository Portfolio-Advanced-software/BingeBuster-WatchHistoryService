package main

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestServerStartAndStop(t *testing.T) {
	// Create a new context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Start the server in a goroutine
	go func() {
		main() // Call the main function to start the server
	}()

	// Wait for the server to start by attempting to establish a connection
	conn, err := grpc.DialContext(ctx, "localhost:50056", grpc.WithInsecure())
	assert.NoError(t, err, "failed to connect to server")
	defer conn.Close()

	// Perform any necessary assertions or tests on the server

	// Stop the server gracefully
	cancel()

	// Wait for the server to stop
	<-ctx.Done()
}
