package main

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SECRETKEY is KEY
var SECRETKEY = []byte("aws12")

// Movie is struct
type Movie struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Popularity float64            `json:"popularity,omitempty" bson:"popularity,omitempty"`
	Name       string             `json:"name" bson:"name,omitempty"`
	Director   string             `json:"director,omitempty" bson:"director,omitempty"`
	Imdbscore  float64            `json:"imdbscore,omitempty" bson:"imdbscore,omitempty"`
	Genre      []string           `json:"genre" bson:"genre,omitempty"`
}

// User is a sruct
type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Username string             `json:"username,omitempty" bson:"username,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
}

// SuccessResponse is response struct
type SuccessResponse struct {
	Code     int
	Message  string
	Response interface{}
}

// ErrorResponse struct
type ErrorResponse struct {
	Code    int
	Message string
}

// Claims struct
type Claims struct {
	Email string
	// Password string
	jwt.StandardClaims
}
