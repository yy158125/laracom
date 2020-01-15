package handler

import (
	"context"
	"errors"
	"github.com/yy158125/laracom/user-service/pkg/auth"
	pb "github.com/yy158125/laracom/user-service/proto/user"
	"github.com/yy158125/laracom/user-service/repo"
	"github.com/yy158125/laracom/user-service/service"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type UserService struct {
	Repo repo.Repository
	Token service.Authable
}

// Create(context.Context, *User, *Response) error
//	Get(context.Context, *User, *Response) error
//	GetAll(context.Context, *Request, *Response) error

func (srv *UserService) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := srv.Repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}
func (srv *UserService) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	users, err := srv.Repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}

func (srv *UserService) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	// 对密码进行哈希加密
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hashedPass)
	if err := srv.Repo.Create(req); err != nil {
		return err
	}
	res.User = req
	return nil

}

func (srv *UserService) Auth(ctx context.Context,req *pb.User, res *pb.Token) error  {
	log.Println("Logging in with:", req.Email, req.Password)
	// 获取用户信息
	user, err := srv.Repo.GetByEmail(req.Email)
	if err != nil {
		return err
	}
	log.Println(user)

	if err := auth.Compare(user.Password,req.Password); err != nil{
		return err
	}
	// token
	token, err := srv.Token.Sign(user)
	if err != nil {
		return err
	}
	res.Token = token
	return nil
}

func (srv *UserService) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	claims, err := srv.Token.Parse(req.Token)
	if err != nil {
		return err
	}

	if claims.User.Id == "" {
		return errors.New("无效用户")
	}
	res.Valid = true

	return nil
}

