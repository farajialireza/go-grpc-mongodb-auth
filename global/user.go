package global

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/o1egl/paseto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User information struct
type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	FirstName string             `bson:"fname"`
	LastName  string             `bson:"lname"`
	MPhone    string             `bson:"mphone"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
}

// Create PASETO token after successful login
func (u User) CreateToken() string {
	env := os.Getenv("GRPC_MONGO_ENV")
	err := godotenv.Load(env)
	if err != nil {
		log.Println("Main Could not open .env file")
	}

	anghezi, err := paseto.NewV2().Encrypt([]byte(os.Getenv("PASETO_SECRET")), u.ID, nil)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return anghezi
}
