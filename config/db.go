package config

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	MongoURI string `json:"mongoURI"`
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
}

var DB *mongo.Client
var AppConfig Config

func LoadConfig() {
	file, err := os.Open("config/config.json")
	fmt.Println("Failed to load configuration file: ")
    if err != nil {
        log.Fatal("Failed to load configuration file: ", err)
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    if err := decoder.Decode(&AppConfig); err != nil {
        log.Fatal("Failed to decode configuration file: ", err)
    }

    log.Printf("Loaded config: %+v", AppConfig)  // Log loaded config for debugging
}

func ConnectDB() {
	LoadConfig()

	clientOptions := options.Client().ApplyURI(AppConfig.MongoURI)
	if AppConfig.User != "" && AppConfig.Password != "" {
		clientOptions.SetAuth(options.Credential{
			Username: AppConfig.User,
			Password: AppConfig.Password,
		})
	}

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal("Failed to create MongoDB client: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB: ", err)
	}

	DB = client
	log.Println("Connected to MongoDB successfully")
}

func GetCollection(collection string) *mongo.Collection {
	// Connect to MongoDB
	ConnectDB()
	if DB == nil {
        log.Fatal("MongoDB client is not initialized")
    }
	return DB.Database(AppConfig.Database).Collection(collection)
}
