package daos

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/lk33/jukebox/internal/daos/database"
	"github.com/lk33/jukebox/internal/dtos"
	"github.com/lk33/jukebox/pkg/config"
	"github.com/lk33/jukebox/pkg/utils"
)

type AlbumDAO struct {
	config *config.Config
	pgSvc  database.IConn
}

type IAlbumDAO interface {
	UpsertAlbum(ctx context.Context, album *dtos.MusicAlbum) error
	GetAlbumsSortedByReleaseDate(ctx context.Context) ([]dtos.Album, error)
	GetAlbumsByMusicianNameSortedByPrice(ctx context.Context, musicianName string) ([]dtos.Album, error)
}

func NewAlbumDAO(config *config.Config) IAlbumDAO {
	return &AlbumDAO{
		pgSvc:  database.ConnPool,
		config: config,
	}
}

func NewMockAlbumDAO(config *config.Config, pgSvc database.IConn) IAlbumDAO {
	return &AlbumDAO{
		pgSvc:  pgSvc,
		config: config,
	}
}

func (a *AlbumDAO) UpsertAlbum(ctx context.Context, album *dtos.MusicAlbum) error {

	releaseDate, err := time.Parse("2006-01-02", album.ReleaseDate)
	if err != nil {
		log.Print("Error parsing release date:", err)
		return err
	}
	var description sql.NullString
	if album.Description != "" {
		description = sql.NullString{String: album.Description, Valid: true}
	} else {
		description = sql.NullString{String: "", Valid: false}
	}

	musicAlbum := dtos.Album{
		Name:        album.Name,
		ReleaseDate: releaseDate,
		Genre:       album.Genre,
		Price:       album.Price,
		Description: description,
	}
	queryFilePath := a.config.Database.QueryLocation + "upsertMusicAlbums.sql"
	query, err := utils.ReadFile(queryFilePath)
	if err != nil {
		return errors.New("Error while reading file " + queryFilePath + err.Error())
	}
	_, err = database.ConnPool.ExecContext(ctx, query, musicAlbum.Name, musicAlbum.ReleaseDate, musicAlbum.Genre, musicAlbum.Price, musicAlbum.Description)
	if err != nil {
		return err
	}
	return nil
}

func (a *AlbumDAO) GetAlbumsSortedByReleaseDate(ctx context.Context) ([]dtos.Album, error) {
	queryFilePath := a.config.Database.QueryLocation + "musicAlbumsByReleaseDate.sql"
	query, err := utils.ReadFile(queryFilePath)
	if err != nil {
		return nil, errors.New("Error while reading file" + queryFilePath + err.Error())
	}

	rows, err := database.ConnPool.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var albums []dtos.Album
	for rows.Next() {
		var album dtos.Album
		err := rows.StructScan(&album)
		if err != nil {
			log.Print(err)
			continue
		}
		albums = append(albums, album)
	}
	return albums, nil
}

func (a *AlbumDAO) GetAlbumsByMusicianNameSortedByPrice(ctx context.Context, musicianName string) ([]dtos.Album, error) {
	// if database.ConnPool == nil {
	// 	return errors.New("database connection unavailable")
	// }
	queryFilePath := a.config.Database.QueryLocation + "musicAlbumsForMusiciansByPrice.sql"
	query, err := utils.ReadFile(queryFilePath)
	if err != nil {
		return nil, errors.New("Error while reading file" + queryFilePath + err.Error())
	}

	rows, err := database.ConnPool.QueryxContext(ctx, query, musicianName)
	if err != nil {
		return nil, err
	}
	var albums []dtos.Album
	for rows.Next() {
		var album dtos.Album
		err := rows.StructScan(&album)
		if err != nil {
			log.Print(err)
			continue
		}
		albums = append(albums, album)
	}
	return albums, nil
}
