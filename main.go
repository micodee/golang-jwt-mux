package main

import (
	"fmt"
	authcontroller "jwtMux/controllers/authController"
	"jwtMux/models"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	models.ConnectDB()
	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	port := "8000"
	fmt.Println("server running on port", port)
	http.ListenAndServe("localhost:"+port, r)
}
