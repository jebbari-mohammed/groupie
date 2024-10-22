package groupietracker

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)
var searchedartist []Artist

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	searchedartist = []Artist{}
	errorrr := fitchArtist()
	if errorrr != nil {
		http.Error(w, errorrr.Error(), http.StatusInternalServerError)
		return
	}
	searched := r.FormValue("Search")
	founded := map[string]bool{}
for _, artist := range Artists {
	if strings.Contains(strings.ToLower(artist.NAME), strings.ToLower(searched)) && !founded[artist.NAME] {
		searchedartist = append(searchedartist, artist)
		founded[artist.NAME] = true
	}
	if strings.Contains(strings.ToLower(artist.FIRSTALBUM), strings.ToLower(searched)) && !founded[artist.NAME] {
		searchedartist = append(searchedartist, artist)
		founded[artist.NAME] = true
	}
	if strings.Contains(strings.ToLower(strconv.Itoa(artist.CREATIONDATES)), strings.ToLower(searched)) && !founded[artist.NAME] {
		searchedartist = append(searchedartist, artist)
		founded[artist.NAME] = true
	}
	for _, member := range artist.MEMBERS {
		if strings.Contains(strings.ToLower(member), strings.ToLower(searched)) && !founded[artist.NAME] {
			searchedartist = append(searchedartist, artist)
			founded[artist.NAME] = true
		}
	}
	for _, location := range DataLocations.Index {
		for _, loc := range location.Location {
			if strings.Contains(strings.ToLower(loc), strings.ToLower(searched)) && !founded[artist.NAME] {
				searchedartist = append(searchedartist, artist)
				founded[artist.NAME] = true
			}
		}
	}
}

temp, err := template.ParseFiles("templates/Searched.html")
if err != nil {
	fmt.Println(err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
}
temp.Execute(w, searchedartist)
}
