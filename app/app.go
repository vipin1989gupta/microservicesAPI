package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	/*
	* The http.HandleFunc() function is a wrapper around the DefaultServeMux.
	 */
	router := mux.NewRouter()
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customer", createCustomer).Methods(http.MethodPost)

	router.HandleFunc("/api/time", createApiTime).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

/*
* Using http.NewServeMux() provides several benefits over using the default http package directly:
* 1. Routing: ServeMux allows you to define multiple routes and their corresponding handlers. This makes it easier to manage different endpoints in your application.
* Using mux, you can create a custom router with specific routes and handlers, which can be more flexible and easier to manage than using the default router.
* 2. Custom Handlers: You can define custom handlers for specific routes, which can help in organizing your code better.
* By using a custom ServeMux, you can isolate your routes and handlers from the global state, which can help prevent conflicts and make your code more modular.
* 3. Middleware: You can implement middleware by wrapping handlers, which can be useful for logging, authentication, etc.
* With a custom ServeMux, you can easily add middleware to your routes by wrapping the handlers, which is more difficult to achieve with the default http package.
* 4. Performance: ServeMux can be more efficient for routing requests compared to using a single handler function with a switch statement.
* A custom ServeMux can be more easily tested in isolation compared to the global http package's default ServeMux.
 */

/*
 * Mux is more of a library that provides a router that simplefies the route definitions and a way to handle requests. It is not a framework like Gin or Echo.
 * Gin is a framework with a lot of features like middleware, request handling, etc. Mux is a simple router that can be used with the standard http package.
 */
