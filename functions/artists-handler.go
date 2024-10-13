package groupietracker

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type ArtistBeta struct {
	Artist   Artist
	Location []string
	Date     []string
	Relation map[string][]string
}

func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(405)
		http.ServeFile(w, r, "eroors/405.html")
		return
	}
	er, err, errr, errrr := fitchArtist(), fitchRelations(), fitchDates(), fitchLocation()
	if er != nil || err != nil || errr != nil || errrr != nil {
		w.WriteHeader(500)
		http.ServeFile(w, r, "eroors/500.html")
		fmt.Println(r)
		return
	}
	id := r.FormValue("id")
	idd, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(404)
		http.ServeFile(w, r, "eroors/404.html")
		return
	}
	var i bool
	var selectedArtist Artist
	for _, Artist := range Artists {
		if idd == Artist.ID {
			i = true
			selectedArtist = Artist
			break
		}
	}
	if !i {
		w.WriteHeader(404)
		http.ServeFile(w, r, "eroors/404.html")
		return
	}
	var selectedLocation Location
	for _, Location := range DataLocations.Index {
		if idd == Location.ID {
			selectedLocation = Location
			break
		}
	}
	var selectedDates ConcertDate
	for _, date := range DataDates.Index {
		if idd == date.ID {
			selectedDates = date
			break
		}
	}
	var selecteRelation Relation
	for _, relation := range DataRelations.Index {
		if idd == relation.ID {
			selecteRelation = relation
			break
		}
	}

	data := ArtistBeta{
		Artist:   selectedArtist,
		Location: selectedLocation.Location,
		Date:     selectedDates.Dates,
		Relation: selecteRelation.DatesLocations,
	}
	tmp, err := template.ParseFiles("templates/artists.html")
	if err != nil {
		// Log the error for debugging purposes
		log.Printf("Error parsing template: %v", err)
		// Set a custom error status code before serving the error file
		w.WriteHeader(http.StatusInternalServerError) // Ensure this is the only place where WriteHeader is called for this request
		http.ServeFile(w, r, "eroors/500.html")
		return
	}
	err = tmp.Execute(w, data)
	if err != nil {
		// Log the error for debugging purposes
		log.Printf("Error executing template: %v", err)
		// Do not set the status code again if it's already set
		http.ServeFile(w, r, "eroors/500.html")
		return
	}
}
