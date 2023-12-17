package controllers

import (
	// "encoding/json"
	"fmt"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
	"go-sql/models"
)

type Book struct{
	ID    int    `json:"id"`
	Name string `json:"name"` 
	Author string `json:"author"`
	Publication string `json:"publication"`	
}

// var NewBook models.Book 

func GetBooks(ctx *gofr.Context) (interface{}, error) {
	var books []Book

	// Getting the customer from the database using SQL
	rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM books")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Name); err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	// return the customer
	return books, nil
	// return "Hello World! GetBooks", nil
}
func GetBookById(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("BookId")
	fmt.Println(id)
	return "Hello World! GetBookById.", nil

}

func CreateBook(ctx *gofr.Context, , emp models.Book) (models.Book, error) {
	var resp models.Book

	queryInsert := "INSERT INTO employees (id, name, email, phone, city) VALUES (?, ?, ?, ?, ?)"

	// Execute the INSERT query
	result, err := ctx.DB().ExecContext(ctx, queryInsert, emp.ID, emp.Name, emp.Email, emp.Phone, emp.City)

	if err != nil {
		return models.Book{}, errors.DB{Err: err}
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return models.Book{}, errors.DB{Err: err}
	}

	// Now, use a separate SELECT query to fetch the inserted data
	querySelect := "SELECT id, name, email, phone, city FROM employees WHERE id = ?"

	// Use QueryRowContext to execute the SELECT query and get a single row result
	err = ctx.DB().QueryRowContext(ctx, querySelect, lastInsertID).
		Scan(&resp.ID, &resp.Name, &resp.Email, &resp.Phone, &resp.City)

	// Handle the error if any
	if err != nil {
		return models.Book{}, errors.DB{Err: err}
	}

	return resp, nil
}

func UpdateBook(ctx *gofr.Context) (interface{}, error) {
	return "Hello World! UpdateBook", nil
}
func DeleteBook(ctx *gofr.Context) (interface{}, error) {
	return "Hello World! DeleteBook", nil
}
