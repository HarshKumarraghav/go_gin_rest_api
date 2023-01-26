package data

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	FirstName string             `json:"firstName"`
	LastName  string             `json:"lastName"`
	Email     string             `json:"email"`
	Age       string             `json:"age"`
}
type InUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Age       string `json:"age"`
}
type OutUser struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	FirstName string             `json:"firstName"`
	LastName  string             `json:"lastName"`
	Email     string             `json:"email"`
	Age       string             `json:"age"`
}

func (in *InUser) ToUser() User {
	return User{
		ID:        primitive.NewObjectID(),
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
		Age:       in.Age,
	}
}
func (in *User) ToOutUser() OutUser {
	return OutUser{
		ID:        in.ID,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
		Age:       in.Age,
	}
}

func (u *OutUser) ToOutUser() OutUser {
	return OutUser{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Age:       u.Age,
	}
}
