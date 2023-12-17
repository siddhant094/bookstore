package routes

import (
	"gofr.dev/pkg/gofr"
)

var RegisterBookStoreRoutes = func () {
	app := gofr.New()
	app.GET("book", func(ctx *gofr.Context) (interface{}, error) {
        return "Hello World!", nil
    })


	// router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	// router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	// router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	// router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	// router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")

	// app.GET("/book/", controllers.GetBooks)
	// app.GET("/book/{bookId}", controllers.GetBookById)
	// app.POST("/book/", controllers.CreateBook)
	// app.PUT("/book/{bookId}", controllers.UpdateBook)
	// app.DELETE("/book/{bookId}", controllers.DeleteBook)
}
