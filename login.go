package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func login(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var user User
	var dbUser User
	_ = json.NewDecoder(request.Body).Decode(&user)
	fmt.Println(user)
	collection := client.Database("users").Collection("register")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&dbUser)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	userPass := []byte(user.Password)
	dbPass := []byte(dbUser.Password)

	passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)

	if passErr != nil {
		log.Println(passErr)
		response.WriteHeader(http.StatusUnauthorized)
		response.Header().Set("Content-Type", "application/json")
		response.Write([]byte(`{"response":"Wrong Password!"}`))
	} else {
		jwtToken, err := GenerateJWT(user.Email)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Header().Set("Content-Type", "application/json")
			response.Write([]byte(`{"message":"` + err.Error() + `"}`))
			return
		}
		addCookie(response, "Bearer", jwtToken)
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	res := make(map[string]interface{})
	res["response"] = "Login Successfully"
	des, _ := json.Marshal(res)
	response.Write([]byte(des))

}
