package main

import (
	"fmt"
	"net/http"

	groupietracker "grptracker/functions"
)

func main() {
	fmt.Println("Server is running on http://localhost:8080/")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", groupietracker.HomeHandler)
	http.HandleFunc("/artists", groupietracker.ArtistsHandler)
	http.HandleFunc("/search/", groupietracker.SearchHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
