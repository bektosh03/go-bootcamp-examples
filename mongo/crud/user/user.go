package user

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        string `json:"id" bson:"_id"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Age       int    `json:"age" bson:"age"`
	Email     string `json:"email" bson:"email"`
}

func (u User) MarshalBSON() ([]byte, error) {
	objID, err := primitive.ObjectIDFromHex(u.ID)
	if err != nil {
		return nil, err
	}

	m := map[string]interface{}{
		"_id":        objID,
		"first_name": u.FirstName,
		"last_name":  u.LastName,
		"age":        u.Age,
		"email":      u.Email,
	}

	return bson.Marshal(m)
}
