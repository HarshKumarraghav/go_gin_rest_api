package data

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// A Repository is a type that can create, get, update, and delete Users.
// @property Create - This method creates a new user.
// @property GetUser - This method is used to get a single user from the database.
// @property GetUsers - This method will return all the users in the database.
// @property UpdateUser - This is a function that takes a user and returns a user.
// @property {error} DeleteUser - This is a function that takes in an ID and deletes the user with that
// ID.
type Repository interface {
	// Create a User
	Create(in InUser) (User, error)
	// // Get a User
	GetUser(id string) (User, error)
	// // Get All Users
	GetUsers() ([]User, error)
	// // Update a User
	UpdateUser(id string, upd map[string]interface{}) (InUser, error)
	// // Delete a User
	DeleteUser(id string) bool
}
type Repo struct {
	db      *mongo.Collection
	context context.Context
}

func (s *Repo) Create(in InUser) (User, error) {
	faq := in.ToUser()
	_, err := s.db.InsertOne(s.context, &faq)
	if err != nil {
		return faq, err
	}
	return faq, nil
}

func (s *Repo) GetUsers() ([]User, error) {
	var users []User
	var cursor *mongo.Cursor
	var user User
	for cursor.Next(s.context) {
		if err := cursor.Decode(&user); err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}
func (s *Repo) UpdateUser(id string, upd map[string]interface{}) (InUser, error) {
	uid, _ := primitive.ObjectIDFromHex(id)
	var faq InUser
	if err := s.db.FindOneAndUpdate(s.context, bson.M{"_id": uid}, map[string]interface{}{
		"$set": upd,
	}).Err(); err != nil {
		return faq, err
	}
	if err := s.db.FindOne(s.context, bson.M{"_id": uid}).Decode(&faq); err != nil {
		return faq, err
	}
	return faq, nil
}
func (s *Repo) DeleteUser(id string) bool {
	uid, _ := primitive.ObjectIDFromHex(id)
	delete, err := s.db.DeleteOne(s.context, bson.M{"_id": uid})
	if err != nil {
		return false
	}
	return delete.DeletedCount == 1
}
func (s *Repo) GetUser(id string) (User, error) {
	oid, _ := primitive.ObjectIDFromHex(id)
	var faq User
	err := s.db.FindOne(s.context, bson.M{"_id": oid}).Decode(&faq)
	if err != nil {
		return faq, errors.New("faq not found with this id")
	}
	return faq, nil
}

// NewRepo returns a Repo which can be used for various operations later.
func NewRepo(db *mongo.Database) Repository {
	ctx := context.TODO()
	return &Repo{db: db.Collection("botintent"), context: ctx}
}
