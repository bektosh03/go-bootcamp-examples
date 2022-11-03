package repository

import (
	"context"
	"mongo/crud/user"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	users *mongo.Collection
}

func NewMongo(ctx context.Context) (*Mongo, error) {
	opts := &options.ClientOptions{}
	opts.ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(opts)
	if err != nil {
		return nil, err
	}

	if err = client.Connect(ctx); err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return &Mongo{
		users: client.Database("test").Collection("users"),
	}, nil
}

func (m *Mongo) CreateUser(ctx context.Context, u user.User) (string, error) {
	res, err := m.users.InsertOne(ctx, u)
	if err != nil {
		return "", err
	}

	id, _ := res.InsertedID.(primitive.ObjectID)
	return id.Hex(), nil
}

func (m *Mongo) ListUsers(ctx context.Context) ([]user.User, error) {
	users := make([]user.User, 0)
	cur, err := m.users.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	if err = cur.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (m *Mongo) UpdateUser(ctx context.Context, u user.User) error {
	id, err := primitive.ObjectIDFromHex(u.ID)
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id": id,
	}
	update := bson.M{
		"$set": bson.M{
			"first_name": u.FirstName,
			"last_name":  u.LastName,
			"age":        u.Age,
			"email":      u.Email,
		},
	}

	_, err = m.users.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (m *Mongo) DeleteUser(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	if _, err = m.users.DeleteOne(ctx, bson.M{"_id": objID}); err != nil {
		return err
	}

	return nil
}

// TODO get single user => GET /user/:id
