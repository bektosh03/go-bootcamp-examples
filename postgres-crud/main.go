package main

import (
	"net"
	"postgres-gin-crud/config"
	"postgres-gin-crud/postgres"
	"postgres-gin-crud/router"
	"postgres-gin-crud/server"

	_ "postgres-gin-crud/docs"
)

// @title           Postgres Crud API
// @version         1.0
// @description     This is a sample server celler server.

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	db, err := postgres.Connect(cfg)
	if err != nil {
		panic(err)
	}

	s := server.New(postgres.NewRepo(db))

	r := router.InitRouter(s)

	r.Run(net.JoinHostPort(cfg.Host, cfg.Port))
}
