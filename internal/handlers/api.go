package handlers

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/internal/models"
	"net/http"
)

const (
	artistsUrl   = "https://groupietrackers.herokuapp.com/api/artists"
	locationsUrl = "https://groupietrackers.herokuapp.com/api/locations"
	relationsUrl = "https://groupietrackers.herokuapp.com/api/relation"
)

func getJson(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch data: %v", err)
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(&target)
}

func GetArtists() ([]models.Artist, error) {
	var artists []models.Artist

	if err := getJson(artistsUrl, &artists); err != nil {
		return nil, err
	}

	return artists, nil
}

func GetLocations() ([]models.Locations, error) {
	var locationsResponse models.LocationsResponse

	if err := getJson(locationsUrl, &locationsResponse); err != nil {
		return nil, err
	}

	return locationsResponse.Index, nil
}

func GetRelations() ([]models.Relation, error) {
	var relationsResponse models.RelationsResponse

	if err := getJson(relationsUrl, &relationsResponse); err != nil {
		return nil, err
	}

	return relationsResponse.Index, nil
}

func LoadGroupDataCache() ([]models.Artist, []models.Locations, []models.Relation, error) {
	artists, err := GetArtists()
	if err != nil {
		return nil, nil, nil, err
	}

	locations, err := GetLocations()
	if err != nil {
		return nil, nil, nil, err
	}

	relations, err := GetRelations()
	if err != nil {
		return nil, nil, nil, err
	}

	return artists, locations, relations, nil
}
