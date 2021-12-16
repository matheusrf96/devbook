package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Running at the port 5000")
	r := router.Generate()

	log.Fatal(http.ListenAndServe(":5000", r))
}
