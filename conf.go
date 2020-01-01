package gw_s3_handler

import (
	"fmt"

	"gopkg.in/gcfg.v1"
)

// Database stores all configuration for postgresql
type Database struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

// ReadDBConnConfig reads ./config/app.ini for DB connection
func ReadDBConnConfig() Database {
	cfg := struct {
		Database
	}{}

	err := gcfg.ReadFileInto(&cfg, "./config/app.ini")
	if err != nil {
		db := Database{Name: "example", Host: "localhost"}
		return db
	}
	fmt.Println(cfg.Database.Host)
	return cfg.Database
}
