package config

import "time"

type Config struct {
	App              Application      `yaml:"app"`
	Server           Server           `yaml:"server"`
	Database         DatabaseConfig   `yaml:"database"`
	DataRestrictions DataRestrictions `yaml:"dataRestrictions"`
}

type Application struct {
	Environment string `yaml:"environment"`
	Name        string `yaml:"name"`
}

type Server struct {
	Port         int           `yaml:"port" default:"8080"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
	ProfilerPort int           `yaml:"profilePort" default:"6060"`
}
type DatabaseConfig struct {
	HostName      string             `yaml:"hostName"`
	Port          int                `yaml:"port"`
	UserName      string             `yaml:"userName"`
	Password      string             `yaml:"password"`
	DatabaseName  string             `yaml:"databaseName"`
	Schema        string             `yaml:"schema"`
	Settings      ConnectionSettings `yaml:"settings"`
	QueryLocation string             `yaml:"queryLocation"`
}

type ConnectionSettings struct {
	MaxOpenConnections    int `yaml:"maxOpenConnections"`
	MaxIdleConnections    int `yaml:"maxIdleConnections"`
	MaxConnectionLifetime int `yaml:"maxConnLifetime"`
}

type DataRestrictions struct {
	MinAlbumCharacters    int   `yaml:"minAlbumCharacters" default:"5"`
	MinMusicianCharacters int   `yaml:"minMusicianCharacters" default:"3"`
	PriceRange            Price `yaml:"priceRange"`
}

type Price struct {
	MinimumPrice float64 `yaml:"minimumPrice" default:"100"`
	MaximumPrice float64 `yaml:"maximumPrice" default:"1000"`
}
