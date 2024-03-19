package main

import (
	"fmt"
	"net/http"

	"Assignment2/controllers"
	"Assignment2/database"

	"github.com/gorilla/mux"
)

func main() {
	database.StartDB()
	r := mux.NewRouter()

	// Rute
	r.HandleFunc("/orders", controllers.CreateOrder).Methods("POST")
	r.HandleFunc("/orders", controllers.GetOrders).Methods("GET")
	r.HandleFunc("/orders/{orderId}", controllers.UpdateOrder).Methods("PUT")
	r.HandleFunc("/orders/{orderId}", controllers.DeleteOrder).Methods("DELETE")

	// Start server
	fmt.Println("Server started on port 9015")
	http.ListenAndServe(":9015", r)
}
