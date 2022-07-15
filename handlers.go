package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Wold")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Wold from home")
}

func UserPostRequest(w http.ResponseWriter, r *http.Request) {
	decode := json.NewDecoder(r.Body)
	var user User
	err := decode.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	response, err := user.ToJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// fmt.Fprintf(w, "Payload %v\n", response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	decode := json.NewDecoder(r.Body)
	var metadata Metadata
	err := decode.Decode(&metadata)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	fmt.Fprintf(w, "Payload %v\n", metadata)

}
