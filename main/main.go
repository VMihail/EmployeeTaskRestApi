package main

import (
	"EmployeeAPI/store"
	"flag"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	config := store.NewConfig()
	config.DataBaseURL = "host=localhost dbname=restapi_dev sslmode=disable"
	s := store.New(config)
	if err := s.Open(); err != nil {
		log.Fatal(err)
	}
}
