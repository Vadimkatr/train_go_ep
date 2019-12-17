package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handleGet()).Methods(http.MethodGet)
	router.HandleFunc("/", handlePost()).Methods(http.MethodPost)

	if err := http.ListenAndServe(":8081", router); err != nil {
		log.Fatal(err)
	}
}

func handleGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// server html page for get request (or reload page)
		http.ServeFile(w, r, "page.html")
	}
}

func handlePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		name := "token=" + r.PostForm.Get("name")
		address := r.PostForm.Get("address")
		// This part of code dont work for we,
		// I cant set cookies there,
		// THIS func (http.SetCookie) just don't work there. Why? I dont know.
		// I spend for this bug several hours and I dont understand why is happening.
		// But if you set another cookies like bellow, all will be all right and code works well)
		// http.SetCookie(w, &http.Cookie{
		// 	Name:    "6",
		// 	Value:   "666",
		// })
		http.SetCookie(w, &http.Cookie{
			Name:  fmt.Sprintf("token=%s", name),
			Value: address,
		})

		http.ServeFile(w, r, "page.html")
	}
}
