package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	config "github.com/piyakker/GoMall/configs"
	"github.com/piyakker/GoMall/health"
)

func main() {

	cfg := config.LoadConfig()

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		checkers := []health.Checker{
			health.NewPostgresChecker(cfg),
			health.NewRedisChecker(cfg),
			health.NewRabbitMQChecker(cfg),
		}

		if err := health.RunAll(ctx, checkers...); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintf(w, "Unhealthy:\n%v", err)
			return
		}
		fmt.Fprintln(w, "OK")
	})

	fmt.Println("API Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
