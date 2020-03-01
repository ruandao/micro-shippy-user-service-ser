package main

import (
	"github.com/micro/go-micro"
	pb "github.com/ruandao/micro-shippy-user-service/ser/proto/user"
	"log"
)

const (
	defaultDB = "database:5432"
)

func main() {
	db, err := CreateConnection()
	if err != nil {
		log.Fatalf("connect database err: %v", err)
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	repository := &UserRepository{db}
	tokenService := &TokenService{}
	h := &handler{repository, tokenService}
	
	srv := micro.NewService(
		micro.Name("go.micro.srv.user"),
	)
	srv.Init()

	pb.RegisterUserServiceHandler(srv.Server(), h)

	if err := srv.Run(); err != nil {
		log.Fatalf("user service err: %v", err)
	}
}
