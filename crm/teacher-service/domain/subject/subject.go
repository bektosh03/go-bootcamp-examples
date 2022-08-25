package subject

import "github.com/google/uuid"

// Subject represents domain object that holds required info for a subject
// All core business logic relevant to subjects should be done through this struct
type Subject struct {
	ID          uuid.UUID
	Name        string
	Description string
}
