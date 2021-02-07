package main

import (
	"fmt"
	"log"
)

func main() {

	/*
		TODO make ref data lke stock symbols available via HTTP and / or gRPC server.
	*/

	result := GetSymbols("US")
	log.Print("printing fetched result")
	fmt.Println(result)

}
