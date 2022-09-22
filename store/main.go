package main

import (
	"context"
	"store/inventory"
	"store/server/http"
	"store/server/telegram"
	"store/store"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	i, err := inventory.NewPostgresInventory(ctx)
	if err != nil {
		panic(err)
	}

	s := store.New(i)

	httpServer := http.NewServer(s)
	telegramBotServer, err := telegram.NewServer("5653971769:AAGRBrNQLubAAeI_P7MZmXhW-j_BiRDS7s4", s)
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
