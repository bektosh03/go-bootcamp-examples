package main

import "log"

func main() {
	service := NewService(LoggingEmailSender{})
	service.Run(make(chan TeacherRegistered))
}

type LoggingEmailSender struct{}

func (s LoggingEmailSender) Send(email Email) error {
	log.Printf("Send an email to %s, with body: %s\n", email.To, email.Body)
	return nil
}
