package main

import (
	"fmt"
	"goChat/internal/common/auth"
	"goChat/internal/common/config"
	"goChat/internal/common/db"
	"goChat/internal/user"
	"log"
	"net/http"
)

func main() {
	config.LoadEnv()

	database := db.InitMongo()

	// Logic connect mongoDB
	userRepo := user.NewRepository(database)

	// Tạo HTTP cho api
	userHandler := user.NewHandler(userRepo)

	authHandler := auth.NewHandler(userRepo)

	http.HandleFunc("/api/login", authHandler.Login)
	http.HandleFunc("/api/register", userHandler.Register)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to My Go Chat Server Version 0.0.1")
	})

	fmt.Println("Server is running at http://localhost:9999")

	// Mở cổng :9999
	// Nếu Server error, Program sẽ stop và log error.
	log.Fatal(http.ListenAndServe(":9999", nil))
}
