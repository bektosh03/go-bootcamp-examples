package main

import (
	"fmt"
	"log"
	"sync"
)

func NewService(emailSender EmailSender, consumer Consumer) Service {
	return Service{
		emailSender: emailSender,
		consumer:    consumer,
	}
}

type Service struct {
	emailSender EmailSender
	consumer    Consumer
}

func (s Service) Run() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		s.handleRegistrations(s.consumer.Events())
	}()

	wg.Wait()
}

func (s Service) handleRegistrations(events <-chan RegisteredEvent) {
	for e := range events {
		body := fmt.Sprintf("Welcome to our CRM, %s %s", e.For, e.FullName)
		email := Email{
			To:   e.Email,
			Body: []byte(body),
		}
		if err := s.emailSender.Send(email); err != nil {
			log.Printf("failed to send email to %s due to: %v\n", e.Email, err)
		}
	}
}

type EmailSender interface {
	Send(email Email) error
}
