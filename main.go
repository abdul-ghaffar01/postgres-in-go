package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func main() {
	conn, err := pgx.Connect(context.Background(), "postgres://learner:3069921@localhost:5432/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())
	fmt.Println("Connected to Postgres")
}
