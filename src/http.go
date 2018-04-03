package main

import (
	"log"
	"net/http"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("The time is: " + r.Host ))
}




func main() {
	mux := http.NewServeMux()

	// Convert the timeHandler function to a HandlerFunc type
	th := http.HandlerFunc(timeHandler)
	// And add it to the ServeMux
	mux.Handle("/time", th)

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}
