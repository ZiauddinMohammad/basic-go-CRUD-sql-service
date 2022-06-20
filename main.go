package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ziauddinmohammad/basic-go-CRUD-sql-service/models"
	"github.com/ziauddinmohammad/basic-go-CRUD-sql-service/routes"
)

func main() {
	//Connects to db and automigrates book model
	models.DbInit()

	//Creates a new router instance
	muxRouter := mux.NewRouter()
	//Add all routes
	routes.RegisterRoutes(muxRouter)

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", muxRouter))
}
