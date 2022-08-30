package response

type Teacher struct {
	ID          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	SubjectID   string `json:"subject_id"`
}

type Subject struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
