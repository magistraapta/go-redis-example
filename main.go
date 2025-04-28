package main

import (
	"log"
	"net/http"
	"os"
	"redis-caching/database"
	"redis-caching/handler"
	"redis-caching/utils"

	"github.com/joho/godotenv"
)

func main() {

	loadEnv()

	// initialize database connection
	db, err := database.ConnectDatabase()

	if err != nil {
		log.Fatal("failed to connect database")
	}

	// initialize redis connecton
	rdb := utils.NewClientRedis()

	// setup router
	mux := http.NewServeMux()

	// setup handler
	userHandler := handler.NewUserHandler(db, rdb)

	mux.Handle("/user/", http.HandlerFunc(userHandler.GetUser))

	// serve on port :8000
	if err := http.ListenAndServe(os.Getenv("PORT"), mux); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// load .env file
func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("failed to get .env file")
	}
}
