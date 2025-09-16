package main

import (
	"BackendLinklyMedia/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Mongo API")
	r := router.Router()
	fmt.Println("Server is getting started....")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening to Port 4000.......")
}
