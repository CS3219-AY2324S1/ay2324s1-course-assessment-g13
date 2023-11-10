package config

import (
	"context"
	"log"
	"os"
	"net/http"
	"net/url"
	"question-service/models"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection
var CategoriesCollection *mongo.Collection
var ctx = context.TODO()

const minDocuments int64 = 100
const leetcodeQuestionsURL = `https://asia-southeast1-peer-preps-assignment6.cloudfunctions.net/GetProblems/?`

func ConnectDb() {
	MONGO_URI := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(MONGO_URI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	Collection = client.Database("questions-service").Collection("questions")
	CategoriesCollection = client.Database("questions-service").Collection("categories")
}

func PopulateDb() {

	count, err := Collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if count >= minDocuments {
		return
	}

	params := url.Values{}
  params.Add("offset", "0")
  params.Add("page-size", "100")
	url := leetcodeQuestionsURL + params.Encode()

	resp, err := http.Get(url)
	if err != nil {
			log.Fatal(err)
	}
	defer resp.Body.Close()

	
	var apiResponse models.APIResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		log.Fatal(err)
	}

	problems := apiResponse.Problems

	for _, problem := range problems {
		question := models.Question{
			Title:       problem.Title,
			Description: problem.Description,
			Categories:  problem.Categories,
			Complexity:  problem.Complexity,
		}

		if err := Collection.FindOne(context.TODO(), bson.M{"title": question.Title}).Err(); err == nil {
			continue
		}
		_, err := Collection.InsertOne(context.TODO(), question)
		if err != nil {
			log.Fatal(err)
		}

		for _, category := range question.Categories {
			if err := CategoriesCollection.FindOne(context.TODO(), bson.M{"category": category}).Err(); err == nil {
				continue
			}
			_, err := CategoriesCollection.InsertOne(context.TODO(), bson.M{"category": category})
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
