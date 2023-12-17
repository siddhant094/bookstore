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

	app.PUT("/book/{bookId}", func (ctx *gofr.Context) (interface{}, error) {
		var emp models.Book
		if err := ctx.Bind(&emp); err != nil {
			ctx.Logger.Errorf("error in binding: %v", err)
			return nil, errors.InvalidParam{Param: []string{"body"}}
		}
	
		resp, err := controllers.UpdateBook(ctx, emp)
		if err != nil {
			return nil, err
		}
	
		return resp, nil
	})
	app.DELETE("/book/{bookId}", controllers.DeleteBook)

	app.Start()
}