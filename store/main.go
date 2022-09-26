package main

import (
	"context"
	"fmt"
	"log"
	"store/config"
	"store/inventory"
	"store/server/http"
	"store/server/telegram"
	"store/store"
	"sync"
	"time"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Println("error with loading config", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	i, err := inventory.NewPostgresInventory(ctx, cfg.PostgresConfig)
	if err != nil {
		panic(err)
	}

	s := store.New(i)
	fmt.Println(cfg.BotApiToken)
	httpServer := http.NewServer(s)
	telegramBotServer, err := telegram.NewServer(cfg.BotApiToken, s)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		if err = httpServer.Run("localhost:8080"); err != nil {
			panic(err)
		}
	}()

	go func() {
		defer wg.Done()
		telegramBotServer.Run()
	}()

	wg.Wait()
}
