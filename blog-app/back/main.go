package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"net/http"
)

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	r := chi.NewRouter()

	r.Use(cors.AllowAll().Handler)
	r.Use(middleware.Logger)

	r.Post("/sign-in", func(w http.ResponseWriter, r *http.Request) {
		var request SignInRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, render.M{
				"error": "bad request",
			})
			return
		}
		fmt.Println("request:", request)

		if request.Username == "bek" && request.Password == "123" {
			fmt.Println("Success")
			render.Status(r, http.StatusOK)
			render.JSON(w, r, render.M{
				"success": true,
			})
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, render.M{
			"success": false,
		})
	})

	r.Get("/blogs", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, blogs)
	})

	http.ListenAndServe("localhost:8080", r)
}

var (
	blogs = []Blog{
		{
			Title:  "Title1",
			Author: "Author1",
			Body:   "Some text1",
		},
		{
			Title:  "Title2",
			Author: "Author2",
			Body:   "Some text2",
		},
		{
			Title:  "Title3",
			Author: "Author3",
			Body:   "Some text3",
		},
	}
)
