package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func deletemovie(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var movie Movie
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Authentication Token Mismatched",
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
		collection := client.Database("imdb").Collection("movie")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := collection.FindOneAndDelete(ctx, Movie{ID: id}).Decode(&movie)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{ "message": "` + err.Error() + `" }`))

		} else {
			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusOK)
			res := make(map[string]interface{})
			res["deletedID"] = id
			res["response"] = "Deleted Successfully"
			des, _ := json.Marshal(res)
			response.Write([]byte(des))
		}
	}

}
