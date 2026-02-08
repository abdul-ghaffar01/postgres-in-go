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
		fmt.Printf("Error occured: %s\n", err)
	}
}
func Crud(ctx context.Context, conn *pgx.Conn) {
	fmt.Println("Starting crud")
	// CreateSchema(ctx, conn)

	// Creating a document in table
	fmt.Println("Creating a document")
	var id int
	err := conn.QueryRow(ctx, "INSERT INTO demo VALUES($1, $2, $3, $4) returning id;", 1, "Abdul Ghaffar", "Islamabad", "ags@gmail.com").Scan(&id)

	fmt.Printf("id: %v\n", id)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	fmt.Println("Reading the same document")
	row := conn.QueryRow(ctx, "Select id, name, city, email from demo where id = $1;", 1)

	var name string
	var city string
	var email string
	err = row.Scan(&id, &name, &city, &email)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	fmt.Printf("Result:\n\tId: %v, name: %s, city: %s, email: %s\n", id, name, city, email)

	fmt.Println("Updating the same document")
	var updatedName string
	err = conn.QueryRow(ctx, "update demo set name = $1 where id = $2 returning name;", "Azlan Ali", 1).Scan(&updatedName)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	fmt.Println("Updated name: ", updatedName)

	fmt.Println("Deleting the same document")
	_, err = conn.Exec(ctx, "Delete from demo where id = $1", 1)

	fmt.Println("Crud completed successfully")
}
