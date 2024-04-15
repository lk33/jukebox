package dtos

import (
	"database/sql"
	"time"
)

type Album struct {
	AlbumID     int            `json:"album_id"`
	Name        string         `json:"name"`
	ReleaseDate time.Time      `json:"release_date"`
	Genre       string         `json:"genre"`
	Price       float64        `json:"price"`
	Description sql.NullString `json:"description"`
}

type AlbumMusician struct {
	AlbumID      int            `json:"album_id"`
	MusicianID   int            `json:"musician_id"`
	MusicianType sql.NullString `json:"musician_type"`
}

type Musician struct {
	MusicianID   int            `json:"musician_id"`
	Name         string         `json:"name"`
	MusicianType sql.NullString `json:"musician_type"`
}
