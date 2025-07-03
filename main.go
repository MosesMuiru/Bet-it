package main

import (
	"encoding/json"
	"fmt"
	"kwik/rabbit"
	"kwik/storage/models"
	"kwik/storage/store"
	"log"
)

type Slip struct {
	Id     uint   `json:"id"`
	Status string `json:"status"`
}

func main() {

	slip := models.Slip{
		Stake:  "100",
		Status: "Won",
		Selections: []models.Selection{
			{IncomeId: 101, Status: "won"},
			{IncomeId: 102, Status: "lost"},
		},
	}

	store := store.NewDb()
	store.CreateSlip(&slip)
	// ehee, testing the store functions
	q, ch := rabbit.NewQ()

	if ch == nil {
		fmt.Println("THis is nil")
		return
	}
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {

		fmt.Println("Failed to read messages", err)
	}

	if err != nil {
		fmt.Println("Error unmarshaling:", err)
		return
	}

	for d := range msgs {

		var newSlip Slip
		json.Unmarshal(d.Body, &newSlip)

		err := store.UpdateSlipStatusById(newSlip.Id, newSlip.Status)
		if err != nil {
			fmt.Println("got an errot", err)
		}

		fmt.Println("got an status", newSlip.Status)
		log.Printf("Received a message: %s", string(d.Body))
	}
}
