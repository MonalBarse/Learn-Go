package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MonalBarse/Learn-Go/tree/main/24-API-MongoDB/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName         = "netflix"
	collectionName = "watchlist"
)

// Collection is a reference to the MongoDB collection
var collection *mongo.Collection

func init() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("MONGO_URI is not set")
	}

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	checkError(err)

	// Ping the MongoDB server to verify connection
	err = client.Ping(context.TODO(), nil)
	checkError(err)

	fmt.Println("Connected to MongoDB!")
	collection = client.Database(dbName).Collection(collectionName)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// CRUD Functions

func insertOneMovie(movie models.Netflix) {
	insertResult, err := collection.InsertOne(context.TODO(), movie)
	checkError(err)
	fmt.Println("Inserted a Record ", insertResult.InsertedID)
}

func updateOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	checkError(err)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	checkError(err)

	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
}

func deleteOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	checkError(err)

	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.TODO(), filter)
	checkError(err)

	fmt.Printf("Deleted %v Documents!\n", result.DeletedCount)
}

func deleteAllRecords() {
	result, err := collection.DeleteMany(context.Background(), bson.D{{}})
	checkError(err)

	fmt.Printf("Deleted %v Documents!\n", result.DeletedCount)
}

func findAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	checkError(err)

	var movies []primitive.M
	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		checkError(err)
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())

	return movies
}

// HTTP Handlers

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := findAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie models.Netflix
	err := json.NewDecoder(r.Body).Decode(&movie)
	checkError(err)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode("Movie has been watched!")
}

func DeleteOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode("Movie has been cleared from watchlist")
}

func DeleteAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	deleteAllRecords()
	json.NewEncoder(w).Encode("All movies have been cleared from watchlist")
}
