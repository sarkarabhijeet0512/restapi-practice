package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func updatemovie(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var movie Movie
	_ = json.NewDecoder(request.Body).Decode(&movie)
	collection := client.Database("imdb").Collection("movie")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := Movie{ID: id}
	update := bson.M{"$set": movie}
	// upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		// Upsert:         &upsert,
	}
	resp := bson.M{}
	result := collection.FindOneAndUpdate(ctx, filter, update, &opt).Decode(&resp)

	if result != nil {
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusNotFound)
		res := make(map[string]interface{})
		// res["updated_data"] = resp
		res["response"] = "Data Not Found"
		des, _ := json.Marshal(res)
		response.Write([]byte(des))
	} else {
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		res := make(map[string]interface{})
		res["updated_data"] = resp
		res["response"] = "Updated Successfully"
		des, _ := json.Marshal(res)
		response.Write([]byte(des))
	}
}
