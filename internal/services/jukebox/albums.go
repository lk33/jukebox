package jukebox

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/lk33/jukebox/internal/daos"
	"github.com/lk33/jukebox/internal/dtos"
	"github.com/lk33/jukebox/pkg/config"
)

type JukeboxAlbumService interface {
	UpsertAlbum(context.Context, dtos.MusicAlbum) error
	GetAlbums(context.Context) (*dtos.MusicAlbums, error)
	GetAlbumsByMusicianName(context.Context, string) (*dtos.MusicAlbums, error)
}

type JukeboxAlbum struct {
	album  daos.IAlbumDAO
	config *config.Config
}

func NewAlbumService(config *config.Config) JukeboxAlbumService {
	album := daos.NewAlbumDAO(config)
	return &JukeboxAlbum{
		album:  album,
		config: config,
	}
}

func NewMockAlbumService(config *config.Config, album daos.IAlbumDAO) JukeboxAlbumService {
	return &JukeboxAlbum{
		album:  album,
		config: config,
	}
}

func (a *JukeboxAlbum) UpsertAlbum(ctx context.Context, album dtos.MusicAlbum) error {
	if len(album.Name) < a.config.DataRestrictions.MinAlbumCharacters {
		log.Print("Invalid album name:", album.Name)
		return errors.New("Invalid album Name:" + album.Name)
	}
	if album.Price < 100 && album.Price > 1000 {
		log.Print("Invalid album price:", album.Price)
		return errors.New("Invalid album price:" + strconv.FormatFloat(album.Price, 'f', -1, 64))
	}
	err := a.album.UpsertAlbum(ctx, &album)
	return err
}

func (a *JukeboxAlbum) GetAlbums(ctx context.Context) (*dtos.MusicAlbums, error) {
	result, err := a.album.GetAlbumsSortedByReleaseDate(ctx)
	if err != nil {
		return nil, err
	}
	var albums []dtos.MusicAlbum
	for _, musicAlbum := range result {
		album := dtos.MusicAlbum{
			AlbumID:     musicAlbum.AlbumID,
			Name:        musicAlbum.Name,
			ReleaseDate: musicAlbum.ReleaseDate.Format("2006-01-02"),
			Genre:       musicAlbum.Genre,
			Price:       musicAlbum.Price,
			Description: musicAlbum.Description.String,
		}
		albums = append(albums, album)
	}
	response := &dtos.MusicAlbums{
		MusicAlbums: albums,
	}
	return response, nil
}

func (a *JukeboxAlbum) GetAlbumsByMusicianName(ctx context.Context, musicianName string) (*dtos.MusicAlbums, error) {
	if len(musicianName) < a.config.DataRestrictions.MinMusicianCharacters {
		log.Print("Invalid album name:", musicianName)
		return nil, errors.New("Invalid album name:" + musicianName)
	}
	result, err := a.album.GetAlbumsByMusicianNameSortedByPrice(ctx, musicianName)
	if err != nil {
		return nil, err
	}
	var albums []dtos.MusicAlbum
	for _, musicAlbum := range result {
		album := dtos.MusicAlbum{
			AlbumID:     musicAlbum.AlbumID,
			Name:        musicAlbum.Name,
			ReleaseDate: musicAlbum.ReleaseDate.Format("2006-01-02"),
			Genre:       musicAlbum.Genre,
			Price:       musicAlbum.Price,
			Description: musicAlbum.Description.String,
		}
		albums = append(albums, album)
	}
	response := &dtos.MusicAlbums{
		MusicAlbums: albums,
	}
	return response, nil
}
