package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Message struct {
	A int
	B string
}

func (m *Message) String() string {
	return fmt.Sprintf("A : %v, B: %v", m.A, m.B)
}

func printHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received %s", r.Method)
	decoder := json.NewDecoder(r.Body)
	var m Message
	decoder.Decode(&m)
	fmt.Printf("Received %v", m)
}

func main() {
	http.HandleFunc("/", printHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
