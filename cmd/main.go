package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	db "github.com/lk33/jukebox/internal/daos/database"
	"github.com/lk33/jukebox/internal/handlers"
	"github.com/lk33/jukebox/pkg/config"
	yaml "gopkg.in/yaml.v2"
)

func main() {

	cancelCtx, cancelF := context.WithCancel(context.Background())
	var (
		cfg config.Config
	)
	defer cancelF()
	loadLocalConfig(&cfg)

	// profiler
	go func() {
		profilerPort := cfg.Server.ProfilerPort
		address := net.JoinHostPort("", strconv.Itoa(profilerPort))
		err := http.ListenAndServe(address, nil)
		if err != nil && err != http.ErrServerClosed {
			log.Print("could not start remote profile server", "err", err, "address", address)
		}
	}()

	stop := make(chan struct{})
	go db.PostgresInit(&cfg.Database, stop)

	router := handlers.GetRouter(&cfg)
	httpServer := &http.Server{
		Addr:         net.JoinHostPort("", strconv.Itoa(cfg.Server.Port)),
		Handler:      router,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("could not start http server: ", err)
		}
	}()

	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)
	signal := <-sign
	log.Print("Cleaning the App as app exit is invoked.", signal.String())
	go gracefulShutdown(cancelCtx, httpServer, stop)
	log.Print("sleeping 3 seconds before exit")
	time.Sleep(3 * time.Second)
	os.Exit(1)
}

func loadLocalConfig(config *config.Config) {
	data, err := os.ReadFile("../config_definition.yaml")
	if err != nil {
		log.Fatal("unable to load configuration from config_definition.yaml", err)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal("unable to load configuration from config_definition.yaml", err)
	}
}

func gracefulShutdown(ctx context.Context, httpServer *http.Server, stop chan struct{}) {
	log.Print("stopping the http server")
	err := httpServer.Shutdown(ctx)
	if err != nil {
		log.Print("could not stop http server", "err", err.Error())
	}
	log.Print("closing database conneciton")
	if db.ConnPool != nil {
		err := db.ConnPool.Close()
		if err != nil {
			log.Print("could not close database connection", "err", err.Error())
		}
	} else {
		stop <- struct{}{}
	}
}
