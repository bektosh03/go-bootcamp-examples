package main

import (
	"encoding/json"
	"strconv"
	"time"
)

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

func (e *Event) UnmarshalJSON(data []byte) error {
	m := make(map[string]string)
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	id, err := strconv.Atoi(m["id"])
	if err != nil {
		return err
	}

	start, err := time.Parse("2006-01-02 15:04:05", m["start"])
	if err != nil {
		return err
	}

	end, err := time.Parse("2006-01-02 15:04:05", m["end"])
	if err != nil {
		return err
	}

	e.ID = id
	e.Name = m["name"]
	e.Organizer = m["organizer"]
	e.Start = start
	e.End = end

	return nil
}

func (e *Event) HasIntersection(otherEvent Event) bool {
	if e.Start.Equal(otherEvent.Start) || e.End.Equal(otherEvent.End) {
		return true
	}

	if e.End.After(otherEvent.Start) && e.Start.Before(otherEvent.End) {
		return true
	}

	return false
}
