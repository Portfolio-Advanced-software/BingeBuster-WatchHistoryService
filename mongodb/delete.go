package mongodb

import (
	"context"
	"fmt"

	"github.com/Portfolio-Advanced-software/BingeBuster-WatchHistoryService/globals"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteHistoryByID(ctx context.Context, id string) (bool, error) {
	// Convert the ID string to an ObjectID
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, fmt.Errorf("invalid ID format: %v", err)
	}

	// Create a filter for the ID field
	filter := bson.M{"_id": oid}

	// Delete the document matching the filter
	result, err := globals.Db.Database(globals.DbName).Collection(globals.CollectionName).DeleteOne(ctx, filter)
	if err != nil {
		return false, fmt.Errorf("failed to delete history: %v", err)
	}

	// Check if the document exists
	if result.DeletedCount == 0 {
		return false, fmt.Errorf("history with ID %s not found", id)
	}

	return true, nil
}

func DeleteHistoryByUserId(ctx context.Context, userId string) (bool, error) {
	// Create a filter for the userID field
	filter := bson.M{"userid": userId}

	// Delete the documents matching the filter
	result, err := globals.Db.Database(globals.DbName).Collection(globals.CollectionName).DeleteMany(ctx, filter)
	if err != nil {
		return false, fmt.Errorf("failed to delete history: %v", err)
	}

	// Check if any documents were deleted
	if result.DeletedCount == 0 {
		return false, fmt.Errorf("no history records found for userID: %s", userId)
	}

	return true, nil
}
