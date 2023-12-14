package mongodb

import (
	"context"
	"fmt"

	"golang_startup/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const userCollection = "users"

// FindUser - Returns a user for a given ID or an error if anything fails
func (db *Database) FindUser(ctx context.Context, id string) (model.User, error) {
	var user model.User

	collection := db.client.Database(db.database).Collection(userCollection)

	cursor, err := collection.Find(ctx, bson.M{"_id": id})
	if err != nil {
		return model.User{}, err
	}

	if cursor.RemainingBatchLength() == 0 {
		return model.User{}, fmt.Errorf("user not found")
	}

	for cursor.Next(ctx) {
		err = cursor.Decode(&user)
		if err != nil {
			return model.User{}, err
		}
	}

	return user, nil
}

// UpdateUser - Updates a user and returns an error if anything fails
func (db *Database) UpdateUser(ctx context.Context, user model.User) error {
	collection := db.client.Database(db.database).Collection(userCollection)

	update := bson.M{"$set": user}

	_, err := collection.UpdateByID(ctx, user.ID, update)

	return err
}

// CreateUser - Creates a user and returns its ID or an error, if anything fails
func (db *Database) CreateUser(ctx context.Context, user model.User) (string, error) {
	collection := db.client.Database(db.database).Collection(userCollection)

	id := primitive.NewObjectID().String()

	user.ID = id

	_, err := collection.InsertOne(ctx, user)

	return id, err
}

// DeleteUser - Deletes a User for a given ID and could return an error if anything fails
func (db *Database) DeleteUser(ctx context.Context, id string) error {
	collection := db.client.Database(db.database).Collection(userCollection)

	_, err := collection.DeleteOne(ctx, bson.M{
		"_id": id,
	})

	return err
}

// FindAllUsers - Returns all users for a given limit and offset
func (db *Database) FindAllUsers(ctx context.Context, offset int64, limit int64) ([]model.User, error) {
	var users []model.User

	collection := db.client.Database(db.database).Collection(userCollection)

	queryOptions := options.Find()
	queryOptions.SetSkip(offset)
	queryOptions.SetLimit(limit)

	cursor, err := collection.Find(ctx, bson.M{}, queryOptions)
	if err != nil {
		return []model.User{}, err
	}

	for cursor.Next(ctx) {
		var user model.User

		err = cursor.Decode(&user)
		if err != nil {
			return []model.User{}, err
		}

		users = append(users, user)
	}

	return users, nil
}
