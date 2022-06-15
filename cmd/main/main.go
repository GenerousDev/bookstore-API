package main

import (
	"fmt"
	"go-bookstore/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// use the name of the imported package to get the methods inward eg "routes"
func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)

	fmt.Println("Started the server on port 9010\n")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
