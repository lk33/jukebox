package jukebox

import (
	"context"
	"errors"
	"log"

	"github.com/lk33/jukebox/internal/daos"
	"github.com/lk33/jukebox/internal/dtos"
	"github.com/lk33/jukebox/pkg/config"
)

type JukeboxMusicianService interface {
	UpsertMusician(context.Context, dtos.MusicianRequest) error
	GetMusiciansByAlbum(context.Context, string) (*dtos.MusiciansByAlbums, error)
}

type JukeboxMusician struct {
	musician daos.IMusicianDAO
	config   *config.Config
}

func NewMusicianService(config *config.Config) JukeboxMusicianService {
	musician := daos.NewMusicianDAO(config)
	return &JukeboxMusician{
		musician: musician,
		config:   config,
	}
}

func (m *JukeboxMusician) UpsertMusician(ctx context.Context, musician dtos.MusicianRequest) error {
	if len(musician.Name) < m.config.DataRestrictions.MinMusicianCharacters {
		log.Print("Invalid musician Name:", musician.Name)
		return errors.New("Invalid musician Name:" + musician.Name)
	}
	err := m.musician.UpsertMusician(ctx, &musician)
	return err
}

func (m *JukeboxMusician) GetMusiciansByAlbum(ctx context.Context, albumName string) (*dtos.MusiciansByAlbums, error) {
	if len(albumName) < m.config.DataRestrictions.MinAlbumCharacters {
		log.Print("Invalid album name:", albumName)
		return nil, errors.New("Invalid album name:" + albumName)
	}
	musicians, err := m.musician.GetMusiciansByAlbumNameSortedByName(ctx, albumName)
	var musicianResponse []dtos.MusicianResponse
	for _, musician := range musicians {
		musicianResponse = append(musicianResponse, dtos.MusicianResponse{
			Name:         musician.Name,
			MusicianType: musician.MusicianType.String,
		})
	}
	if err != nil {
		return nil, err
	}
	response := &dtos.MusiciansByAlbums{
		AlbumName: albumName,
		Musicians: musicianResponse,
	}
	return response, nil
}
