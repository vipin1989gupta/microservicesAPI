package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode int    `json:"zip_code" xml:"zip_code"`
}

type TimeZone struct {
	Time string `json:"current_time"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Vipin Gupta", City: "Bangalore", Zipcode: 560076},
		{Name: "Pragya Gupta", City: "Bangalore", Zipcode: 560076},
		{Name: "Kashvi Gupta", City: "Bangalore", Zipcode: 560076},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		fmt.Println("Sendint XML response")
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		fmt.Println("Sending JSON response")
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["customer_id"]
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, "Customer ID: ", customerID)
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request received")
}

func createApiTime(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API Time")
	response := make(map[string]string)

	timeZone := r.URL.Query().Get("tz")

	if timeZone == "" {
		fmt.Println("Time Zone not provided so using UTC")
		timeZone = "UTC"
	}

	fmt.Println("Time Zone: ", timeZone)
	// Check if user requested for multiple timezones
	if len(timeZone) > 1 {
		// Extract the comma delimited timezones
		timeZones := strings.Split(timeZone, ",")
		for _, tz := range timeZones {
			loc, err := time.LoadLocation(tz)
			if err != nil {
				fmt.Errorf("Error loading location: %v", err)
				http.Error(w, "Invalid time zone", http.StatusBadRequest)
				return
			}
			currentTime := time.Now().In(loc).Format("2006-01-02 15:04:05.999999999 -0700 MST")
			response[tz] = currentTime
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
