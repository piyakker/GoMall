package main

import (
	"fmt"
	"net/http"

	config "github.com/piyakker/GoMall/configs"
)

func main() {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})

	cfg := config.LoadConfig()
	fmt.Printf("DB Name: %s\n", cfg.DBName)

	fmt.Println("API Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
