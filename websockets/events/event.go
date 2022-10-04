package main

import "time"

type ConflictMessage struct {
	To        string  `json:"to"`
	Conflicts []Event `json:"conflicts"`
}

type Event struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Organizer string    `json:"organizer"`
	Start     time.Time `json:"start"`
	End       time.Time `json:"end"`
}

func (e Event) HasIntersection(otherEvent Event) bool {
	if e.Start.Equal(otherEvent.Start) || e.End.Equal(otherEvent.End) {
		return true
	}

	if e.End.After(otherEvent.Start) && e.Start.Before(otherEvent.End) {
		return true
	}

	return false
}
