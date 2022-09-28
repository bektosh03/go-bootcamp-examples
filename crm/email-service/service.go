package main

import (
	"fmt"
	"log"
	"sync"
)

func NewService(emailSender EmailSender) Service {
	return Service{
		emailSender: emailSender,
	}
}

type Service struct {
	emailSender EmailSender
}

func (s Service) Run(teacherEvents <-chan TeacherRegistered) {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		s.handleTeacherRegistrations(teacherEvents)
	}()

	wg.Wait()
}

func (s Service) handleTeacherRegistrations(events <-chan TeacherRegistered) {
	for e := range events {
		body := fmt.Sprintf("Welcome to our CRM, teacher %s", e.FullName)
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
