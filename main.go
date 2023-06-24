package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

const (
	// .devcontainerの.env参照
	HOST     = "localhost"
	DATABASE = "postgres"
	USER     = "postgres"
	PASSWORD = "postgres"
)

type User struct {
	Id   int    `db:"user_id"`
	Name string `db:"user_name"`
}

func main() {
	fmt.Println("hello")
	var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOST, USER, PASSWORD, DATABASE)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var user User
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %s, Name: %s\n", strconv.Itoa(user.Id), user.Name)
	}
}
