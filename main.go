package main

import (
	"github.com/micro/go-micro"
	"github.com/ruandao/micro-shippy-user-service-ser/lib"
	pb "github.com/ruandao/micro-shippy-user-service-ser/proto/user"
	"log"
)

const (
	defaultDB = "database:5432"
)

func main() {
	db, err := lib.CreateConnection()
	if err != nil {
		log.Fatalf("connect database err: %v", err)
	}
	defer db.Close()

	db.AutoMigrate(&lib.User{})

	repository := &lib.UserRepository{DB: db}
	tokenService := &lib.TokenService{}
	h := &lib.Handler{Repository:repository, TokenService:tokenService}
	
	srv := micro.NewService(
		micro.Name(lib.CONST_SER_NAME_USER_SERVICE),
	)
	srv.Init()

	pb.RegisterUserServiceHandler(srv.Server(), h)

	if err := srv.Run(); err != nil {
		log.Fatalf("user service err: %v", err)
	}
}
