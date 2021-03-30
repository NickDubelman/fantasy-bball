package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func init() {
	// Initialize config from env vars
	getenv := func(key, fallback string) string {
		value := os.Getenv(key)
		if len(value) == 0 {
			return fallback
		}
		return value
	}

	port, err := strconv.Atoi(getenv("BBALL_DB_PORT", "3306"))
	if err != nil {
		log.Fatal("Could not parse BBALL_DB_PORT as int")
	}

	c = Configuration{
		Database: databaseConfiguration{
			User:     getenv("BBALL_DB_USER", "root"),
			Password: getenv("BBALL_DB_PASSWORD", ""),
			Host:     getenv("BBALL_DB_HOST", "localhost"),
			Port:     port,
			DBName:   getenv("BBALL_DB_DBNAME", "fantasy"),
		},
	}

}

var c Configuration

func Get() Configuration {
	return c
}

type Configuration struct {
	Database databaseConfiguration
}

type databaseConfiguration struct {
	User, Password string
	Host           string
	Port           int
	DBName         string
}

func (c databaseConfiguration) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		c.User, c.Password, c.Host, c.Port, c.DBName,
	)
}
