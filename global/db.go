package global

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB holds database connection
var DB mongo.Database

// NewDBContext returns a new context according to app performance
func NewDBContext(d time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), d*100/100)
}

func ConnectToDB() {
	env := os.Getenv("GRPC_MONGO_ENV")
	err := godotenv.Load(env)
	if err != nil {
		log.Println("Main Could not open .env file")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://alireza:sfNeQYadliv9oAkD@cluster0.9gs5hrc.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))
	if err != nil {
		log.Println("Could not connect to db", err)
	}

	if err != nil {
		log.Fatal("Error connecting to DB", err.Error())
	}
	DB = *client.Database("cp_platform_v01beta")
}
