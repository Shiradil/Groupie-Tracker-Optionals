package models

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Data struct {
	Artists []Artist
}

type AllArtists struct {
	ID    int    `json:"id"`
	Image string `json:"image"`
	Name  string `json:"name"`
}
