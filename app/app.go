package app

import "net/http"

func Start() {
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	http.ListenAndServe("localhost:8000", nil)
}
