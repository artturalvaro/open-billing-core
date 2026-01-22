package main

import (
	"encoding/json"
	"log"
	"net/http"

	app "open-billing-core/internal/application/subscription"
	"open-billing-core/internal/infrastructure/repository/memory"
)

func setupServer() http.Handler {
	repo := memory.NewSubscriptionRepository()
	useCase := app.NewCreateSubscription(repo)

	mux := http.NewServeMux()

	mux.HandleFunc("/subscriptions", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var input app.CreateSubscriptionInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		sub, err := useCase.Execute(input)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(sub)
	})

	return mux
}

func main() {
	log.Println("HTTP server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", setupServer()))
}
