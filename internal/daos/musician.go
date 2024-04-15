package daos

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/lk33/jukebox/internal/daos/database"
	"github.com/lk33/jukebox/internal/dtos"
	"github.com/lk33/jukebox/pkg/config"
	"github.com/lk33/jukebox/pkg/utils"
)

type MusicianDAO struct {
	config *config.Config
	pgSvc  database.IConn
}

type IMusicianDAO interface {
	UpsertMusician(ctx context.Context, musician *dtos.MusicianRequest) error
	GetMusiciansByAlbumNameSortedByName(ctx context.Context, albumName string) ([]dtos.Musician, error)
}

func NewMusicianDAO(config *config.Config) IMusicianDAO {
	return &MusicianDAO{
		pgSvc:  database.ConnPool,
		config: config,
	}
}

func NewMockMusicianDAO(config *config.Config, pgSvc database.IConn) IMusicianDAO {
	return &MusicianDAO{
		pgSvc:  pgSvc,
		config: config,
	}
}

func (m *MusicianDAO) UpsertMusician(ctx context.Context, musician *dtos.MusicianRequest) error {
	// if database.ConnPool == nil {
	// 	return errors.New("database connection unavailable")
	// }
	queryFilePath := m.config.Database.QueryLocation + "createMusician.sql"
	query, err := utils.ReadFile(queryFilePath)
	if err != nil {
		return errors.New("Error while reading file" + queryFilePath + err.Error())
	}

	var musicianType sql.NullString
	if musician.MusicianType != "" {
		musicianType = sql.NullString{String: musician.MusicianType, Valid: true}
	} else {
		musicianType = sql.NullString{String: "", Valid: false}
	}

	musicianToInsert := dtos.Musician{
		Name:         musician.Name,
		MusicianType: musicianType,
	}
	_, err = database.ConnPool.ExecContext(ctx, query, musicianToInsert.Name, musicianToInsert.MusicianType)
	if err != nil {
		return err
	}
	return nil
}

func (m *MusicianDAO) GetMusiciansByAlbumNameSortedByName(ctx context.Context, albumName string) ([]dtos.Musician, error) {
	// if database.ConnPool == nil {
	// 	return errors.New("database connection unavailable")
	// }
	queryFilePath := m.config.Database.QueryLocation + "musiciansForMusicAlbumByName.sql"
	query, err := utils.ReadFile(queryFilePath)
	if err != nil {
		return nil, errors.New("Error while reading file" + queryFilePath + err.Error())
	}

	rows, err := database.ConnPool.QueryxContext(ctx, query, albumName)
	if err != nil {
		return nil, err
	}
	var musicians []dtos.Musician
	for rows.Next() {
		var musician dtos.Musician
		err := rows.StructScan(&musician)
		if err != nil {
			log.Print(err)
			continue
		}
		musicians = append(musicians, musician)
	}
	return musicians, nil
}
