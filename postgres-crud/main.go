package main

import (
	"postgres-gin-crud/config"
	"postgres-gin-crud/postgres"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	_, err = postgres.Connect(cfg)
	if err != nil {
		panic(err)
	}
}
