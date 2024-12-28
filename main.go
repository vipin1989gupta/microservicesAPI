package main

import (
	"fmt"

	"github.com/microservicesAPI/app"
)

func main() {
	fmt.Println("Hello, World!")

	app.Start()

	// http.HandleFunc("/greet", greet)
	// http.HandleFunc("/customers", getAllCustomers)

	// log.Fatal(http.ListenAndServe("localhost:8080", nil))

}
