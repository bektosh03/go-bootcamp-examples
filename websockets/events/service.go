package main

import "fmt"

type Service struct {
	repo      Repository
	conflicts chan<- ConflictMessage
}

func NewService(repo Repository, conflicts chan<- ConflictMessage) Service {
	return Service{
		repo:      repo,
		conflicts: conflicts,
	}
}

func (s Service) AddEvent(event Event) error {
	if err := s.repo.AddEvent(event); err != nil {
		return err
	}

	go func() {
		otherEvents, err := s.repo.ListEvents(event.Organizer)
		if err != nil {
			fmt.Println("failed to list events:", err)
			return
		}

		conflicts := s.getConflicts(event, otherEvents)
		if len(conflicts) < 1 {
			return
		}
		fmt.Println("sending conflicts to ch")
		s.conflicts <- ConflictMessage{
			To:        event.Organizer,
			Conflicts: conflicts,
		}
	}()

	return nil
}

func (s Service) getConflicts(event Event, otherEvents []Event) []Event {
	conflicts := make([]Event, 0)
	for _, otherEvent := range otherEvents {
		if event.HasIntersection(otherEvent) {
			conflicts = append(conflicts, otherEvent)
		}
	}

	return conflicts
}
