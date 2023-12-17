package main

import (
	"fmt"
	"go-mysql/controllers"
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


func main() {
	// initialise gofr object
	app := gofr.New()
	fmt.Println("Starting server on port 9000...")
	app.GET("/greet", func(ctx *gofr.Context) (interface{}, error) {
        return "Hello Duniya!", nil
    })

	app.GET("/book", controllers.GetBooks)
	app.GET("/book/{bookId}", controllers.GetBookById)
	app.POST("/book", func (ctx *gofr.Context) (interface{}, error) {
		var emp models.Book
		if err := ctx.Bind(&emp); err != nil {
			ctx.Logger.Errorf("error in binding: %v", err)
			return nil, errors.InvalidParam{Param: []string{"body"}}
		}
	
		resp, err := controllers.CreateBook(ctx, emp)
		if err != nil {
			return nil, err
		}
	
		return resp, nil
	})

	app.PUT("/book/{bookId}", controllers.UpdateBook)
	app.DELETE("/book/{bookId}", controllers.DeleteBook)

	
	// app.GET("/greet", func(ctx *gofr.Context) (interface{}, error) {
	// 	// Get the value using the redis instance
	// 	value, err := ctx.Redis.Get(ctx.Context, "greeting").Result()

    //     return value, err
    // })

	// app.POST("/customer/{name}", func(ctx *gofr.Context) (interface{}, error) {
	// 	name := ctx.PathParam("name")

	// 	// Inserting a customer row in database using SQL
	// 	_, err := ctx.DB().ExecContext(ctx, "INSERT INTO customers (name) VALUES (?)", name)

	// 	return nil, err
	// })

	// app.GET("/customer", func(ctx *gofr.Context) (interface{}, error) {
	// 	var customers []Customer

	// 	// Getting the customer from the database using SQL
	// 	rows, err := ctx.DB().QueryContext(ctx, "SELECT * FROM customers")
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	for rows.Next() {
	// 		var customer Customer
	// 		if err := rows.Scan(&customer.ID, &customer.Name); err != nil {
	// 			return nil, err
	// 		}

	// 		customers = append(customers, customer)
	// 	}

	// 	// return the customer
	// 	return customers, nil
	// })

	// Starts the server, it will listen on the default port 8000.
	// it can be over-ridden through configs
	app.Start()
}