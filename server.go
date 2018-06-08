package main

import (

	Service "Rentmatics_Appartment/Services"
	
	
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const (
	contenttypeJSON = "application/json; charset=utf-8"
)

func Serve() bool {

	router := mux.NewRouter()


	router.HandleFunc("/GetAllAppartments", Service.GetAllAppartments)
	router.HandleFunc("/InsertContactProperty", Service.InsertContactProperty)

	

	serverhttp := func() {
		c := cors.New(cors.Options{
			AllowedOrigins:   []string{"*", "http://develop.rentmatics.com", "http://api.msg91.com"},
			AllowCredentials: true,
		})
		handler := c.Handler(router)

		fmt.Println("Server should be available at http",":8084")
		fmt.Println(http.ListenAndServe(":8084", handler))
	}

	
	

	// Schedule API server
	go serverhttp()

	return true
}
