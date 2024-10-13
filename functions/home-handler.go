package groupietracker

import (
	"fmt"
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(404)
		http.ServeFile(w, r, "eroors/404.html")
		return
	}
	if r.Method != "GET" {
		w.WriteHeader(405)
		http.ServeFile(w, r, "eroors/405.html")
		return
	}
	errorrr := fitchArtist()
	if errorrr != nil {
		w.WriteHeader(500)
		http.ServeFile(w, r, "eroors/500.html")
		return
	}
	tmp, err := template.ParseFiles("templates/home.html")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		http.ServeFile(w, r, "eroors/500.html")
		return
	}
	err = tmp.Execute(w, Artists)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		http.ServeFile(w, r, "eroors/500.html")
		return
	}
}
