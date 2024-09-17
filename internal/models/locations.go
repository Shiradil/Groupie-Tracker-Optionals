package models

type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type LocationsResponse struct {
	Index []Locations `json:"index"`
}
