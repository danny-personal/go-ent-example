package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

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
	for i := 1; i <= 10; i++ {
		u, err := h.client.User.
			Create().
			SetAge(i).
			SetName("hoge").
			Save(ctx)
		if err != nil {
			fmt.Println("failed creating user: %w", err)
		}
		log.Println("user was created: ", u)
		time.Sleep(500 * time.Millisecond)
	}
}
