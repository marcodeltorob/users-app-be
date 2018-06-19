package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/router"
)

func main() {
	httpRouter := router.Initrouter()

	fmt.Printf("Sever running on: 8888 port\n")
	log.Fatal(http.ListenAndServe(":8888", httpRouter))
}
