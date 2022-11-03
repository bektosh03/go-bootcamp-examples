package main

import (
	"context"
	"fmt"
	"mongo/crud/repository"
	"mongo/crud/user"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	mongoRepo, err := repository.NewMongo(ctx)
	if err != nil {
		panic(err)
	}

	u := user.User{
		ID:        "63635f9094fcaf465684d742",
		FirstName: "Shaxzod3",
		LastName:  "Ibrohimov3",
		Age:       24,
		Email:     "shaxzod3@gmail.com",
	}

	if err = mongoRepo.DeleteUser(ctx, u.ID); err != nil {
		panic(err)
	}

	users, err := mongoRepo.ListUsers(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(users)
}
