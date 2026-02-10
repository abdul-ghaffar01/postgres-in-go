package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Simple style of connection
	ctx := context.Background()
	conn, _ := pgx.Connect(ctx, "postgres://learner:3069921@localhost:5432/testdb")
	defer conn.Close(ctx)

	conn.Exec(ctx, "Some sql")

}
func ConnectPooling() {
	// it creates pool of multiple connections and uses which one is free,
	//  we can't create connection for each request because each connection is a process in postgres backend.

	// Creating a pool
	ctx := context.Background()
	dburl := "connection string"
	pool, err := pgxpool.New(ctx, dburl)
	if err != nil {
		fmt.Errorf("%s", err.Error())
	}
	defer pool.Close()

	sql := "Select ...."
	pool.Exec(ctx, sql)

}
