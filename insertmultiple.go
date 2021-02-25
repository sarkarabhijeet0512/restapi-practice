package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

func insertmultimovies(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var movie []interface{}
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "It's not you it's me.",
	}
	bearerToken := request.Header.Get("cookie")
	var authorizationToken string
	authorizationTokenArray := strings.Split(bearerToken, "=")
	if len(authorizationTokenArray) > 1 {
		authorizationToken = authorizationTokenArray[1]
	}
	email, _ := VerifyToken(authorizationToken)
	if email == "" {
		returnErrorResponse(response, request, errorResponse)
	} else {
		_ = json.NewDecoder(request.Body).Decode(&movie)
		collection := client.Database("imdb").Collection("movie")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		result, err := collection.InsertMany(ctx, movie)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		} else {
			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusCreated)
			res := make(map[string]interface{})
			res["insertedID"] = result.InsertedIDs
			res["response"] = "Successfully Created"
			des, _ := json.Marshal(res)
			response.Write([]byte(des))
		}
	}
}
