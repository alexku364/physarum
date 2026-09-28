package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"physarum/internal/api"
	"physarum/internal/database"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background())

	router := api.Router(db)

	fmt.Println("Server started on :8080")

	err = http.ListenAndServe("0.0.0.0:8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
