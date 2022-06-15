package routes

import (
	"go-bookstore/pkg/controllers"

	"github.com/gorilla/mux"
	// "github.com/omolajacodec/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/create", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/getbyid/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/update/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/delete/{bookId}", controllers.DeleteBook).Methods("DELETE")
}

// Putting "/" at the back of path might throw error is not placed the same way in testing platform
