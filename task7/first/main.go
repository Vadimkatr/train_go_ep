package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Header struct {
	Accept    string `json:"Accept"`
	UserAgent string `json:"User-Agent"`
}
type Response struct {
	Host       string `json:"host"`
	UserAgent  string `json:"user_agent"`
	RequestUri string `json:"request_uri"`
	Headers    Header `json:"headers"`
}

func handleAny(w http.ResponseWriter, r *http.Request) {
	resp := Response{
		Host:       r.Host,
		UserAgent:  r.Header.Get("User-Agent"),
		RequestUri: fmt.Sprintf("%v %v", r.Method, r.RequestURI),
		Headers: Header{
			Accept:    r.Header.Get("Accept"),
			UserAgent: r.UserAgent(),
		},
	}
	j, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(j)
}

func main() {
	http.HandleFunc("/", handleAny)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Panic(err)
	}
	log.Println("Start server at :8080")
}
