package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CreateSchema(ctx context.Context, conn *pgx.Conn) {
	query := "CREATE TABLE demo(ID INT PRIMARY KEY, NAME VARCHAR(50), CITY VARCHAR(50), EMAIL VARCHAR(100));"
	_, err := conn.Exec(ctx, query)
	if err != nil {
		fmt.Errorf("Error occured: %s", err)
	}
}
func Crud(ctx context.Context, conn *pgx.Conn) {
	fmt.Println("Starting crud")

	// Creating a document in table

	CreateSchema(ctx, conn)
}
