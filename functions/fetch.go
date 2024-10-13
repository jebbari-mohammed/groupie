package groupietracker

import (
	"encoding/json"
	"io"
	"net/http"
)

type Artist struct {
	ID            int      `json:"id"`
	IMAGE         string   `json:"image"`
	NAME          string   `json:"name"`
	MEMBERS       []string `json:"members"`
	CREATIONDATES int      `json:"creationDate"`
	FIRSTALBUM    string   `json:"firstAlbum"`
}
type Location struct {
		ID       int      `json:"id"`
		Location []string `json:"locations"`
}
type ConcertDate struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`

}
type Relation struct {
		ID             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
}

var DataLocations struct {
	Index     []Location        `json:"index"`
}
var DataDates struct {
	Index     []ConcertDate        `json:"index"`
}
var DataRelations struct {
	Index     []Relation       `json:"index"`
}
var Artists []Artist

func fitchArtist() error {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	
	err = json.Unmarshal(res, &Artists)
	if err != nil {
		return err
	}
	return nil
}
func fitchLocation() error {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	println()
	err = json.Unmarshal(res, &DataLocations)
	if err != nil {
		return err
	}
	return nil
}
func fitchDates() error {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(res, &DataDates)
	if err != nil {
		return err
	}
	return nil
}
func fitchRelations() error {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(res, &DataRelations)
	if err != nil {
		return err
	}
	return nil
}