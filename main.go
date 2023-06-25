package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/danny-personal/go-ent-example/ent"
	_ "github.com/lib/pq"
)

const (
	// .devcontainerの.env参照
	HOST     = "localhost"
	DATABASE = "postgres"
	USER     = "postgres"
	PASSWORD = "postgres"
	PORT     = "5432"
)

type Handler struct {
	client *ent.Client
}

func NewHandler(client *ent.Client) *Handler {
	return &Handler{client: client}
}

func main() {
	var connectionString string = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", HOST, PORT, USER, DATABASE, PASSWORD)
	client, err := ent.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	handler := NewHandler(client)
	http.HandleFunc("/", handler.userHandler)
	http.ListenAndServe(":8080", nil)
}

func (h *Handler) userHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	u, err := h.client.User.Query().All(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("user returned: ", u)

	if err := json.NewEncoder(w).Encode(u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
