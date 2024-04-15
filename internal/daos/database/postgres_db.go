package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	_ "github.com/lib/pq"
	"github.com/lk33/jukebox/pkg/config"
)

var ConnPool IConn = nil
var pgOnce sync.Once

type IConn interface {
	Close() error
	ExecContext(ctx context.Context, sql string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, sql string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, sql string, args ...interface{}) *sql.Row
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
}

func PostgresInit(conf *config.DatabaseConfig, stop chan struct{}) {
	if conf.HostName == "" || conf.HostName == "None" {
		log.Print("postgres connection failed. hostName = ", conf.HostName)
		return
	}
	ticker := time.NewTicker(1 * time.Second)
	defer func() {
		ticker.Stop()
	}()
	if ConnPool == nil {
		pgOnce.Do(func() {
			for {
				select {
				case <-ticker.C:
					psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
						conf.UserName, conf.Password, conf.HostName, conf.Port, conf.DatabaseName)
					conn, err := sqlx.Connect("postgres", psqlInfo)
					if err != nil {
						log.Print("Failed to connnect", err)
						continue
					}
					connSettings := conf.Settings
					multiplier := connSettings.MaxConnectionLifetime
					conn.SetMaxOpenConns(connSettings.MaxOpenConnections)
					conn.SetMaxIdleConns(connSettings.MaxIdleConnections)
					conn.SetConnMaxLifetime(time.Duration(multiplier) * time.Minute)
					conn = conn.Unsafe()
					conn.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
					ConnPool = conn
					log.Print("successful postgres connect")
					return
				case <-stop:
					log.Print("stop postgres connection tries")
					return
				}
			}
		})

	}
}
