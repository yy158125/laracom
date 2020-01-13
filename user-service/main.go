package main

import (
	"fmt"
	"github.com/micro/go-micro"
	db "github.com/yy158125/laracom/user-service/db"
	"github.com/yy158125/laracom/user-service/handler"
	pb "github.com/yy158125/laracom/user-service/proto/user"
	repository "github.com/yy158125/laracom/user-service/repo"
)

func main() {
	db := db.OpenDB()
	defer db.Close()

	db.AutoMigrate(&pb.User{})

	repo := &repository.UserRepository{Db:db}

	// Micro 创建微服务流程
	srv := micro.NewService(
		micro.Name("laracom.user.service"),
		micro.Version("latest"),
	)
	srv.Init()

	// 注册处理器
	pb.RegisterUserServiceHandler(srv.Server(), &handler.UserService{Repo: repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

}
