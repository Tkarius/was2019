package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := InitRouter()
	fmt.Println("Http server starting up.")
	log.Fatal(http.ListenAndServe(":443", router))
}
