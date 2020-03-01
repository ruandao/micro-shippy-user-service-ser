package main

import (
	"context"
	"github.com/jinzhu/gorm"
	pb "github.com/ruandao/micro-shippy-user-service/ser/proto/user"
)

type User struct {
	Id       string `json: "Id"`
	Name     string `json: "Name"`
	Company  string `json: "Company"`
	Email    string `json: "Email"`
	Password string `json: "Password"`
}
type Response struct {
	User   *User
	Users  []*User
	Errors []*Error
}
type Request struct {
}
type Token struct {
}
type Error struct {
	Code        int32  `json: "code"`
	Description string `json: "description"`
}

func MarshalUser(user *pb.User) *User {
	return &User{
		Id:       user.Id,
		Name:     user.Name,
		Company:  user.Company,
		Email:    user.Email,
		Password: user.Password,
	}
}
func UnmarshalUser(user *User) *pb.User {
	return &pb.User{
		Id:      user.Id,
		Name:    user.Name,
		Company: user.Company,
		Email:   user.Email,
	}
}

func MarshalResponse(response *pb.Response) *Response {
	users := make([]*User, 0, len(response.Users))
	for _, user := range response.Users {
		users = append(users, MarshalUser(user))
	}

	errors := make([]*Error, 0, len(response.Errors))
	for _, xerror := range response.Errors {
		errors = append(errors, MarshalError(xerror))
	}
	return &Response{
		User:   MarshalUser(response.User),
		Users:  users,
		Errors: errors,
	}
}
func MarshalError(error2 *pb.Error) *Error {
	return &Error{
		Code:        error2.Code,
		Description: error2.Description,
	}
}

type Repository interface {
	GetAll(ctx context.Context) ([]*User, error)
	Get(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, user *pb.User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func (repo *UserRepository) Create(ctx context.Context, user *pb.User) error {
	if err := repo.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) GetByEmail(ctx context.Context, email string) (*User, error) {
	user := &User{Email:email}
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) GetAll(ctx context.Context) ([]*User, error) {
	var users []*User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) Get(ctx context.Context, id string) (*User, error) {
	user := &User{}
	user.Id = id
	if err := repo.db.First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
