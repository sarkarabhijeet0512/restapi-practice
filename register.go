package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

func register(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var user User
	_ = json.NewDecoder(request.Body).Decode(&user)
	user.Password = getHash([]byte(user.Password))
	collection := client.Database("users").Collection("register")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
	} else {
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusCreated)
		res := make(map[string]interface{})
		res["insertedID"] = result.InsertedID
		res["response"] = "Registered Successfully"
		des, _ := json.Marshal(res)
		response.Write([]byte(des))
	}
	// json.NewEncoder(response).Encode(result) isse result ka data print hota hai
}
