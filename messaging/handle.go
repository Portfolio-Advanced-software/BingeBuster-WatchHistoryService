package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Portfolio-Advanced-software/BingeBuster-WatchHistoryService/mongodb"
)

type Message struct {
	UserId  string `json:"userid"`
	MovieId string `json:"movieid"`
	Action  string `json:"action"`
}

func HandleMessage(body []byte) {
	jsonStr := string(body)
	var msg Message
	err := json.Unmarshal([]byte(jsonStr), &msg)
	if err != nil {
		log.Println("Failed to unmarshal JSON:", err)
		return
	}

	switch msg.Action {
	case "deleteAllRecords":
		_, err := mongodb.DeleteHistoryByUserId(context.Background(), msg.UserId)
		if err != nil {
			log.Println("Failed to delete all records:", err)
		}
	default:
		fmt.Println("Unknown action:", msg.Action)
	}
}
