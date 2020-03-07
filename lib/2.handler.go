package lib

import (
	pb "github.com/ruandao/micro-shippy-user-service-ser/proto/user"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

type Handler struct {
	Repository   Repository
	TokenService *TokenService
}

func (srv *Handler) Create(ctx context.Context, user *pb.User, resp *pb.Response) error {
	// Generates a hashed version of our password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPass)
	if err := srv.Repository.Create(ctx, user); err != nil {
		return err
	}
	resp.User = user
	return nil
}

func (srv *Handler) Get(ctx context.Context, user *pb.User, resp *pb.Response) error {
	storeUser, err := srv.Repository.Get(ctx, user.Id)
	if err != nil {
		return err
	}
	resp.User = UnmarshalUser(storeUser)
	return nil
}

func (srv *Handler) GetAll(ctx context.Context, _ *pb.Request, resp *pb.Response) error {
	storeUsers, err := srv.Repository.GetAll(ctx)
	if err != nil {
		return err
	}
	users := make([]*pb.User, 0, len(storeUsers))
	for _, user := range storeUsers {
		users = append(users, UnmarshalUser(user))
	}
	resp.Users = users
	return nil
}

func (srv *Handler) Auth(ctx context.Context, req *pb.User, resp *pb.Token) error {
	user, err := srv.Repository.GetByEmail(ctx, req.Email)
	if err != nil {
		return err
	}

	// Compares our given password against the hashed password
	// stored in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}

	token, err := srv.TokenService.Encode(UnmarshalUser(user))
	if err != nil {
		return err
	}
	resp.Token = token
	return nil
}

func (srv *Handler) ValidateToken(context.Context, *pb.Token, *pb.Token) error {
	// get user from token
	return nil
}
