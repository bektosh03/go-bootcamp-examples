package main

type Repository interface {
	AddEvent(event Event) error
	ListEvents(organizer string) ([]Event, error)
}

type InMemoryRepository struct {
	events map[string]Event
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		events: make(map[string]Event),
	}
}

func (r *InMemoryRepository) AddEvent(event Event) error {
	r.events[event.Organizer] = event

	return nil
}

func (r *InMemoryRepository) ListEvents(organizer string) ([]Event, error) {
	events := make([]Event, 0, len(r.events))
	for k, v := range r.events {
		if k == organizer {
			continue
		}
		events = append(events, v)
	}

	return events, nil
}
