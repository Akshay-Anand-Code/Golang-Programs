//go get go.mongodb.org/mongo-driver/mongo
package main

import (
	"fmt"
	"log"
	"mongoAPI/router"
	"net/http"
)

func main() {
	fmt.Println("mongoDB API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000")

}
