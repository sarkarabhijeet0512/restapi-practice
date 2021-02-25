package main

import (
	"log"

	"github.com/dgrijalva/jwt-go"
)

// GenerateJWT is used for generating a token
func GenerateJWT(email string) (string, error) {
	// expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Email: email,
		// Password: password,
		StandardClaims: jwt.StandardClaims{

			// ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString(SECRETKEY)
	if err != nil {
		log.Println("Error in JWT token generation")
		return "", err
	}
	return tokenString, nil
}

// VerifyToken is function for Token verification
func VerifyToken(tokenString string) (email string, err error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return SECRETKEY, nil
	})
	if token != nil {
		return claims.Email, nil
	}
	return "", err
}
