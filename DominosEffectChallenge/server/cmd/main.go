package main

import (
	"fmt"
	"log"
	"net/http"
	"servicedependencygraph/controllers"
)

func main() {

	port := "8080"
	fmt.Println("Serverssss started succesfully...............")

	http.HandleFunc("/ServiceGraph", controllers.JsonMarhslling)

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatalf("Error in starting server: %v", err)
	}

}
