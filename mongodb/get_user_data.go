package mongodb

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Portfolio-Advanced-software/BingeBuster-WatchHistoryService/globals"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllUserData(ctx context.Context, id string, callback func(string) error) error {
	// Create a filter for the userID field
	filter := bson.M{"userid": id}

	// Delete the documents matching the filter
	cur, err := globals.HistoryDb.Find(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to find records: %v", err)
	}
	defer cur.Close(ctx) // Close the cursor once done

	var records []bson.M

	// Iterate over the cursor and append the documents to the records slice
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			return fmt.Errorf("failed to decode document: %v", err)
		}
		records = append(records, result)
	}

	if err := cur.Err(); err != nil {
		return fmt.Errorf("cursor error: %v", err)
	}

	// Convert the records slice to a JSON string
	jsonData, err := json.Marshal(records)
	if err != nil {
		return fmt.Errorf("failed to marshal records to JSON: %w", err)
	}

	// Convert the JSON byte array to a string
	dataString := string(jsonData)

	if err := callback(dataString); err != nil {
		return fmt.Errorf("callback error: %w", err)
	}

	return nil
}
