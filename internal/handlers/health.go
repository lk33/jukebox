package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/lk33/jukebox/internal/dtos"
	"github.com/lk33/jukebox/pkg/config"
)

type HealthService struct {
	config *config.Config
}

func setPingRoutes(router *mux.Router, config *config.Config) {
	h := NewHeathService(config)
	router.HandleFunc("/jukebox/live", h.liveHandler).Methods("GET")
	router.HandleFunc("/jukebox/ready", h.readyHandler).Methods("GET")
}

var startTime int64

func init() {
	startTime = time.Now().Unix() * 1e3
}

func NewHeathService(config *config.Config) *HealthService {
	return &HealthService{
		config: config,
	}
}

func (h *HealthService) liveHandler(w http.ResponseWriter, r *http.Request) {
	response := dtos.HealthProbeResponse{
		State: true,
		Message: &dtos.LiveProbeResponse{
			AppName:   h.config.App.Name,
			Env:       h.config.App.Environment,
			AppStatus: "OK",
			UpTime:    uptime(),
		},
	}
	writeJSONStruct(w, r, response, http.StatusOK)
}

func (h *HealthService) readyHandler(w http.ResponseWriter, r *http.Request) {
	response := dtos.HealthProbeResponse{
		State: true,
	}
	if response.State {
		writeJSONStruct(w, r, response, http.StatusOK)
	} else {
		writeJSONStruct(w, r, response, http.StatusInternalServerError)
	}
}

func uptime() string {
	currentTime := (time.Now().Unix()) * 1e3
	duration := currentTime - startTime
	return strconv.FormatInt(duration, 10) + "ms"
}
