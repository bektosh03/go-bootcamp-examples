package telegram

import (
	"context"
	"fmt"
	tele "gopkg.in/telebot.v3"
	"log"
	"store/product"
	"store/store"
	"strings"
	"time"
)

func NewServer(apiToken string, s *store.Store) (Server, error) {
	b, err := tele.NewBot(tele.Settings{
		Token: apiToken,
		Poller: &tele.LongPoller{
			Timeout: time.Second * 10,
		},
	})
	if err != nil {
		return Server{}, err
	}

	return Server{
		bot:   b,
		store: s,
	}, nil
}

type Server struct {
	bot   *tele.Bot
	store *store.Store
}

func (s Server) Run() {
	s.bot.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	s.bot.Handle("/products", func(c tele.Context) error {
		products, err := s.store.ListProducts(context.Background())
		if err != nil {
			log.Println("failed to list products:", err)
			return c.Send("something went wrong")
		}

		return c.Send(buildProductsResponse(products))
	})

	s.bot.Handle("/markup", func(c tele.Context) error {
		return c.Send(tele.ReplyMarkup{
			InlineKeyboard: [][]tele.InlineButton{
				{
					{
						Unique:          "",
						Text:            "some text",
						Data:            "",
						InlineQuery:     "",
						InlineQueryChat: "",
						Login:           nil,
					},
				},
			},
		})
	})

	s.bot.Start()
}

func buildProductsResponse(list product.List) string {
	var response strings.Builder
	for i, p := range list {
		response.WriteString(
			fmt.Sprintf(
				"No %d: name = %s, quantity = %d, price = %.2f\n",
				i+1, p.Name, p.Quantity, p.Price,
			),
		)
	}

	return response.String()
}
