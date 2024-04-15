package handlers

import (
	"github.com/gorilla/mux"
	"github.com/lk33/jukebox/pkg/config"
)

func GetRouter(config *config.Config) *mux.Router {
	muxRouter := mux.NewRouter()
	setJukeboxRoutes(muxRouter, config)
	setPingRoutes(muxRouter, config)
	return muxRouter
}
