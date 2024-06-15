package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml: "name"`
	City    string `json:"city" xml: "city"`
	ZipCode string `json:"zip_code" xml: "zip_code"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Darshan", City: "bangalore", ZipCode: "560073"},
		{Name: "Nikitha", City: "Mysore", ZipCode: "560063"},
		{Name: "Nikdhan", City: "bangalore", ZipCode: "560073"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}
