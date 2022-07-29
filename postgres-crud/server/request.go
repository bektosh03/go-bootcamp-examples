package server

type CreateAuthorRequest struct {
	Name string `json:"name"`
}

type CreateBookRequest struct {
	Title    string `json:"title"`
	AuthorID string `json:"author_id"`
}
