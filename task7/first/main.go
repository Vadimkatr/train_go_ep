package main

import (
	"encoding/json"
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
		RequestUri: r.RequestURI,
		Headers: Header{
			Accept:    r.Header.Get("Accept"),
			UserAgent: r.UserAgent(),
		},
	}
	j, _ := json.Marshal(resp)

	w.Write(j)
}

func main() {
	http.HandleFunc("/", handleAny)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Panic(err)
	}
}
