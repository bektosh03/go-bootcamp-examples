package user

type User struct {
	ID        string `json:"id" bson:"_id"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Age       int    `json:"age" bson:"age"`
	Email     string `json:"email" bson:"email"`
}
