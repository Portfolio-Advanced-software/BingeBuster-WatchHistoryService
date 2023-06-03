package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Portfolio-Advanced-software/BingeBuster-WatchHistoryService/globals"
	"github.com/Portfolio-Advanced-software/BingeBuster-WatchHistoryService/mongodb"
)

type Message struct {
	UserId  string `json:"user_id"`
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
	case "getAllRecords":
		err := mongodb.GetAllUserData(context.Background(), msg.UserId, SendMessage)
		if err != nil {
			log.Println("Failed to get all records:", err)
		}
	case "deleteAllRecords":
		_, err := mongodb.DeleteHistoryByUserId(context.Background(), msg.UserId)
		if err != nil {
			log.Println("Failed to delete all records:", err)
		}
	default:
		fmt.Println("Unknown action:", msg.Action)
	}
}

func SendMessage(data string) error {
	conn, err := ConnectToRabbitMQ(globals.RabbitMQUrl)
	if err != nil {
		log.Fatalf("Can't connect to RabbitMQ: %s", err)
	}
	ProduceMessage(conn, data, "user_data")

	return nil
}
