package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name  string `json:"name"`
	IdNum int    `json:"idNum`
}

func main() {
	fmt.Println("Starting server...")
	http.HandleFunc("/", ReqHandler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func ReqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := Person{
		Name:  "Jyotisman",
		IdNum: 1,
	}
	p2 := Person{}

	switch r.Method {
	case "GET":
		err := json.NewEncoder(w).Encode(p1)
		if err != nil {
			fmt.Println("Error found:", err)
		}
	case "POST":
		err := json.NewDecoder(r.Body).Decode(&p2)
		if err != nil {
			fmt.Println("Error found:", err)
		}
		fmt.Fprintf(w, "Input: %+v", p2)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
