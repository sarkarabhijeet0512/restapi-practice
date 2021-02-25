package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func getmovie(response http.ResponseWriter, request *http.Request) {
	content := request.Header.Get("content-type")
	var getmovie []Movie
	movieName := request.URL.Query().Get("name") //get name as parameter from the search
	genre := request.URL.Query().Get("genre")
	popularity := request.URL.Query().Get("popularity")
	director := request.URL.Query().Get("director")
	imdbscore := request.URL.Query().Get("imdbscore")

	collection := client.Database("imdb").Collection("movie")
	filter := bson.M{}
	if movieName != "" {
		filter["name"] =
			bson.M{
				"$regex":   movieName,
				"$options": "i",
			}
	}
	if genre != "" {
		filter["genre"] =
			bson.M{
				"$regex":   genre,
				"$options": "i",
			}
	}
	n, _ := strconv.ParseInt(popularity, 10, 64)
	if popularity != "" {
		filter["popularity"] = bson.M{
			"$gt": n,
		}
	}
	if director != "" {
		filter["director"] =
			bson.M{
				"$regex":   director,
				"$options": "i",
			}
	}
	imdb, _ := strconv.ParseFloat(imdbscore, 32)
	if imdbscore != "" {
		filter["imdbscore"] =
			bson.M{
				"$gt": imdb,
			}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var movie Movie
		cursor.Decode(&movie)
		getmovie = append(getmovie, movie)
	}

	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	if content == "application/json" {
		response.Header().Set("content-type", "application/json")
		response.WriteHeader(http.StatusOK)
		res := make(map[string]interface{})
		res["Search_data"] = getmovie
		des, _ := json.Marshal(res)
		response.Write([]byte(des))
		return
	}
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles("template.html")
	if err != nil {
		fmt.Fprintf(response, "Unable to load template")
	}
	t.Execute(response, getmovie)
}
