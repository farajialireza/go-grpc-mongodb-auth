package main

import (
	"context"
	"go-grpc-mongodb-auth/global"
	"go-grpc-mongodb-auth/grpcstuff/pb"
	"testing"
)

func Test_authServer_Signup(t *testing.T) {
	global.ConnectToDB()
	server := authServer{}
	_, err := server.Signup(context.Background(), &pb.SignupRequest{FirstName: "test_fname", LastName: "test_lname", MPhone: "999999999", Email: "test@gmail.com", Password: "123456789", RepeatPassword: "123456789"})
	if err != nil {
		t.Error(err)
	}
}
