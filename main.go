package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func main() {
	fmt.Println("Starting the application...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	fmt.Println("Connected to MongoDB!")
	fmt.Println("Connected to server 127.0.0.1:8080")
	router := mux.NewRouter()
	router.HandleFunc("/", handler)
	router.HandleFunc("/api/register", register).Methods("POST")
	router.HandleFunc("/api/login", login).Methods("POST")
	router.HandleFunc("/api/movie", getmovie).Methods("GET")
	router.HandleFunc("/api/movie/{id}", getmoviewithid).Methods("GET")
	router.HandleFunc("/api/movie", createmovie).Methods("POST")
	router.HandleFunc("/api/multimovie", insertmultimovies).Methods("POST")
	router.HandleFunc("/api/movie/{id}", updatemovie).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", deletemovie).Methods("DELETE")
	http.ListenAndServe(":8080", router)

}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
