package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)


func about (w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is your About response")
	fmt.Fprint(w, "This is your path", r.URL.Path)
}


func home (w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is your Home response")
	fmt.Fprint(w, "This is your path", r.URL.Path)
}

type Response struct {
	Message string `json:"message"`
	Status string	`json:"status"`
}

func perfectResponse (rw http.ResponseWriter, req * http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	res := Response {
		Message: "This is your perfect response message",
		Status: "ok",
	}

	// Create a JSON encoder that writes directly to the HTTP response stream and then using that encode the res
	json.NewEncoder(rw).Encode(res)

}

func PrintGoServer () {
	fmt.Println("****************************Printing go server basics*******************")

	// in go we write handler functions to handle any request on a given path/endpoint. The handleFunc takes the path and a function which uses frollowing param
	// http.ResponseWrite : Used to send response back to client
	// *http.Request : it contains request data like url, heasers, method, body
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Helo from first go server, I am here to serve you")
		fmt.Print("*******************dontest : ",r.Body, r.Method, r.Header)
	})

	http.HandleFunc("/about", about)
	http.HandleFunc("/home", home)
	http.HandleFunc("/perfect", perfectResponse)

	// A simple command to start go server
	http.ListenAndServe(":8080", nil)
}