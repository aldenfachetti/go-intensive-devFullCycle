package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/aldenfachetti/go-intensive-devFullCycle/internal/infra/database"
	"github.com/aldenfachetti/go-intensive-devFullCycle/internal/usecase"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/aldenfachetti/go-intensive-devFullCycle/pkg/kafka"
)

func main() {
	db, err := sql.Open("sqlite3", "./orders.db")
	if err != nil {
		panic(err)
	}
	defer db.Close() // run everything and then run close

	repository := database.NewOrderRepository(db)
	usecase := usecase.CalculateFinalPrice{OrderRepository: repository}

	msgChanKafka := make(chan *ckafka.Message)
	topics := []string{"orders"}
	servers := "host.docker.internal:9094"
	go kafka.Consume(topics, servers, msgChanKafka)
	kafkaworker(msgChanKafka, usecase)
}

func kafkaworker(msgChan chan *ckafka.Message, uc usecase.CalculateFinalPrice) {
	for msg := range msgChan {
		var OrderInputDTO usecase.OrderInputDTO
		err := json.Unmarshal(msg.Value, &OrderInputDTO)
		if err != nil {
			panic(err)
		}

		outputDTO, err := uc.Execute(OrderInputDTO)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Kafka has processed order %s\n", outputDTO.ID)
	}
}