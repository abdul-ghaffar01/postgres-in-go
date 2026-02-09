package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CreateDocumentAndUpdate(ctx context.Context, conn *pgx.Conn) error {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("Error: %s", err.Error())
	}

	defer tx.Rollback(ctx)

	// Creating the user
	tx.Exec(ctx, "INSERT INTO demo VALUES($1, $2, $3, $4)", 3, "Abdul Ghaffar", "Islamabad", "email@gmail.com")

	// Updating the city
	tx.Exec(ctx, "UPDATE demo SET city = $1 WHERE id = $2", "Karachi", 3)

	tx.Commit(ctx)

	return nil
}

func FindUser(ctx context.Context, conn *pgx.Conn, id int) {
	row := conn.QueryRow(ctx, "Select id, name, city, email from demo where id = $1;", id)

	var name string
	var city string
	var email string
	err := row.Scan(&id, &name, &city, &email)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	fmt.Printf("Result:\n\tId: %v, name: %s, city: %s, email: %s\n", id, name, city, email)

}

func Transactions(ctx context.Context, conn *pgx.Conn) {
	fmt.Println("Transactions--------------------")

	CreateDocumentAndUpdate(ctx, conn)
	FindUser(ctx, conn, 3)
}
