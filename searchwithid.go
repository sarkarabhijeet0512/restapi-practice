package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getmoviewithid(response http.ResponseWriter, request *http.Request) {
	// response.Header().Set("content-type", "application/json")

	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var movie Movie
	collection := client.Database("imdb").Collection("movie")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, Movie{ID: id}).Decode(&movie)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
	} else {
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		res := make(map[string]interface{})
		res["Search_data"] = movie
		des, _ := json.Marshal(res)
		response.Write([]byte(des))
	}
}
