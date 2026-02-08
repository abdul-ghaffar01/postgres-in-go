package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://learner:3069921@localhost:5432/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())
	fmt.Println("Connected to Postgres")

	Crud(ctx, conn)
}
