package main

import (
	"fmt"
	"goChat/internal/common/config"
	"log"
	"net/http"
)

func main() {
	config.LoadEnv()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to My Go Chat Server Version 0.0.1")
	})

	fmt.Println("Server is running at http://localhost:9999")

	// Mở cổng :9999
	// Nếu Server error, Program sẽ stop và log error.
	log.Fatal(http.ListenAndServe(":9999", nil))
}
