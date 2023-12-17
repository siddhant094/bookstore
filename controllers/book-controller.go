package controllers

import (
	"fmt"
	"go-mysql/models"

	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type Book struct{
	ID string `json:"id"`
	Name string `json:"name"` 
	Author string `json:"author"`
	Publication string `json:"publication"`	
}


func GetBooks(ctx *gofr.Context) (interface{}, error) {
	var books []Book

	rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM books")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Name, &book.Author, &book.Publication); err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}
func GetBookById(ctx *gofr.Context) (interface{}, error) {
	var books []Book
	id := ctx.PathParam("bookId")
	fmt.Println(id)
	queryInsert := "SELECT * FROM books WHERE id=?"

	rows, err := ctx.DB().QueryContext(ctx, queryInsert, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Name, &book.Author, &book.Publication); err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func CreateBook(ctx *gofr.Context, emp models.Book) (models.Book, error) {
	var resp models.Book

	queryInsert := "INSERT INTO books (name, author, publication) VALUES (?, ?, ?)"

	result, err := ctx.DB().ExecContext(ctx, queryInsert, emp.Name, emp.Author, emp.Publication)

	if err != nil {
		return models.Book{}, errors.DB{Err: err}
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return models.Book{}, errors.DB{Err: err}
	}

	querySelect := "SELECT * FROM books WHERE id = ?"

	err = ctx.DB().QueryRowContext(ctx, querySelect, lastInsertID).
		Scan(&resp.ID, &resp.Name, &resp.Author, &resp.Publication)

	if err != nil {
		return models.Book{}, errors.DB{Err: err}
	}

	return resp, nil
}

func UpdateBook(ctx *gofr.Context, emp models.Book) (models.Book, error) {
	var resp models.Book
	id := ctx.PathParam("bookId")
	fmt.Println(id)
	
	queryInsert := "UPDATE books SET name=?, author=?, publication=? WHERE id=?"

	_, err := ctx.DB().ExecContext(ctx, queryInsert, emp.Name, emp.Author, emp.Publication, id)

	if err != nil {
		return models.Book{}, errors.DB{Err: err}
	}

	if err != nil {
		return models.Book{}, errors.DB{Err: err}
	}

	querySelect := "SELECT * FROM books WHERE id = ?"

	err = ctx.DB().QueryRowContext(ctx, querySelect, id).
		Scan(&resp.ID, &resp.Name, &resp.Author, &resp.Publication)

	if err != nil {
		return models.Book{}, errors.DB{Err: err}
	}

	return resp, nil
}


func DeleteBook(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("bookId")
	fmt.Println(id)

	queryInsert := "DELETE FROM books WHERE id=?"

	result, err := ctx.DB().ExecContext(ctx, queryInsert, id)
	fmt.Println(result)

	if err != nil {
		fmt.Println("ERROR ")
		fmt.Println(err)
		return models.Book{}, errors.DB{Err: err}
	}
	fmt.Println("Success ")
	return "Book Deleted", nil
}
