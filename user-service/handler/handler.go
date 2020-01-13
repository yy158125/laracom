package handler

import (
	"context"
	pb "github.com/yy158125/laracom/user-service/proto/user"
	"github.com/yy158125/laracom/user-service/repo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo repo.Repository
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
	return nil

}

// docker run -d -p 9003:3306 --name rn3 -e MYSQL_MASTER_HOST=rn1 -e MYSQL_MASTER_PORT=3306 -e MYSQL_ROOT_PASSWORD=abc123456 -e MYSQL_REPLICATION_USER=backup -e MYSQL_REPLICATION_PASSWORD=backup123 -v rnv3:/var/lib/mysql --net=swarm_mysql rep