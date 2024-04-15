package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lk33/jukebox/internal/dtos"
	"github.com/lk33/jukebox/internal/services/jukebox"
	"github.com/lk33/jukebox/pkg/config"
)

type JukeboxService struct {
	config          *config.Config
	albumService    jukebox.JukeboxAlbumService
	musicianService jukebox.JukeboxMusicianService
}

func setJukeboxRoutes(router *mux.Router, config *config.Config) {
	s := NewJukeboxService(config)
	router.HandleFunc("/jukebox/v1/album", s.UpsertAlbum).Methods("POST")
	router.HandleFunc("/jukebox/v1/musician", s.UpsertMusician).Methods("POST")
	router.HandleFunc("/jukebox/v1/albums", s.GetAlbums).Methods("GET")
	router.HandleFunc("/jukebox/v1/albums", s.GetAlbumsByMusicianName).Queries("musician_name", "{musician_name}").Methods("GET")
	router.HandleFunc("/jukebox/v1/musicians", s.GetMusiciansByAlbum).Queries("album_name", "{album_name}").Methods("GET")
}

func NewJukeboxService(config *config.Config) *JukeboxService {
	return &JukeboxService{
		config:          config,
		albumService:    jukebox.NewAlbumService(config),
		musicianService: jukebox.NewMusicianService(config),
	}
}

func (s *JukeboxService) UpsertAlbum(w http.ResponseWriter, r *http.Request) {
	var album dtos.MusicAlbum

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&album)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = s.albumService.UpsertAlbum(r.Context(), album)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *JukeboxService) UpsertMusician(w http.ResponseWriter, r *http.Request) {
	var musician dtos.MusicianRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&musician)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = s.musicianService.UpsertMusician(r.Context(), musician)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *JukeboxService) GetAlbumsByMusicianName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	musicianName := vars["musician_name"]
	albums, err := s.albumService.GetAlbumsByMusicianName(r.Context(), musicianName)
	if err != nil {
		writeJSONMessage(w, r, err.Error(), ERR_MSG, http.StatusInternalServerError)
		return
	}
	writeJSONStruct(w, r, albums, http.StatusOK)
}

func (s *JukeboxService) GetMusiciansByAlbum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	albumName := vars["album_name"]
	musicians, err := s.musicianService.GetMusiciansByAlbum(r.Context(), albumName)
	if err != nil {
		writeJSONMessage(w, r, err.Error(), ERR_MSG, http.StatusInternalServerError)
		return
	}
	writeJSONStruct(w, r, musicians, http.StatusOK)
}

func (s *JukeboxService) GetAlbums(w http.ResponseWriter, r *http.Request) {
	albums, err := s.albumService.GetAlbums(r.Context())
	if err != nil {
		writeJSONMessage(w, r, err.Error(), ERR_MSG, http.StatusInternalServerError)
		return
	}
	writeJSONStruct(w, r, albums, http.StatusOK)
}
