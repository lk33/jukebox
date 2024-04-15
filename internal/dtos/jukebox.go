package dtos

type MusicAlbums struct {
	MusicAlbums []MusicAlbum `json:"albums"`
}

type MusicAlbum struct {
	AlbumID     int     `json:"album_id"`
	Name        string  `json:"name"`
	ReleaseDate string  `json:"release_date"`
	Genre       string  `json:"genre"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type MusicAlbumsByMusician struct {
	MusicianName string       `json:"musician_name"`
	MusicAlbums  []MusicAlbum `json:"albums"`
}

type MusiciansByAlbums struct {
	AlbumName string             `json:"album_name"`
	Musicians []MusicianResponse `json:"musicians"`
}

type MusicianRequest struct {
	MusicianID   int    `json:"musician_id"`
	Name         string `json:"name"`
	MusicianType string `json:"musician_type"`
}

type MusicianResponse struct {
	Name         string `json:"name"`
	MusicianType string `json:"musician_type"`
}
