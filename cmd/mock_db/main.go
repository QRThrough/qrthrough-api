package main

import (
	"log"

	"github.com/JMjirapat/qrthrough-api/config"
	"github.com/JMjirapat/qrthrough-api/infrastructure"
	"github.com/JMjirapat/qrthrough-api/internal/core/model"
)

func init() {
	config.InitConfig()
	infrastructure.InitDB()
}

func main() {
	user := model.Alumni{
		ID:        620610020,
		Firstname: "Somchai",
		Lastname:  "Jaidee",
		Tel:       "0801234567",
	}
	if err := infrastructure.DB.Create(&user); err != nil {
		log.Printf("%v", err)
	}

	log.Printf("Created")
}
