package main

import (
	"context"
	"errors"
	"log"
	"regexp"
	"strings"
	"time"

	"go-grpc-mongodb-auth/global"
	"go-grpc-mongodb-auth/grpcstuff/pb"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func (server authServer) Signup(_ context.Context, in *pb.SignupRequest) (*pb.SignupResponse, error) {
	// Checking entered password length
	if len(in.GetPassword()) < 8 {
		return &pb.SignupResponse{Result: false}, errors.New("password must be at least 8 characters")
	}

	// Checking if email is entered
	if !strings.EqualFold(in.GetEmail(), "") {
		// Checking email format using regex
		match, _ := regexp.MatchString("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$", in.GetEmail())
		if !match {
			return &pb.SignupResponse{Result: false}, errors.New("entered email is not a valid email address")
		}
	}

	// Check if entered passwords are equal
	if !strings.EqualFold(in.GetPassword(), in.GetRepeatPassword()) {
		return &pb.SignupResponse{Result: false}, errors.New("entered passwords are not equal")
	}

	// Hashing password using bcrypt
	pw, err := bcrypt.GenerateFromPassword([]byte(in.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return &pb.SignupResponse{Result: false}, err
	}

	// Store user input data into global.User struct
	user := global.User{
		ID:        primitive.NewObjectID(),
		FirstName: in.GetFirstName(),
		LastName:  in.GetLastName(),
		MPhone:    in.GetMPhone(),
		Email:     in.GetEmail(),
		Password:  string(pw),
	}

	// Connecting to DB server and store information in DB
	ctx, cancel := global.NewDBContext(5 * time.Second)
	defer cancel()
	_, err = global.DB.Collection("mmbrs").InsertOne(ctx, user)
	if err != nil {
		log.Println("Error inserting new user", err.Error())
		return &pb.SignupResponse{Result: false}, err
	}

	return &pb.SignupResponse{Result: true, MPhone: user.MPhone}, err
}
