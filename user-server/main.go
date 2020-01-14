package main

import (
	"fmt"
	"github.com/micro/go-micro"
	db "github.com/yy158125/laracom/user-service/db"
	"github.com/yy158125/laracom/user-service/handler"
	pb "github.com/yy158125/laracom/user-service/proto/user"
	"github.com/yy158125/laracom/user-service/repo"
)

func main()  {
	db := db.OpenDB()
	defer db.Close()

	db.AutoMigrate(&pb.User{})

	repo := &repo.UserRepository{Db:db}

	users,err := repo.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(users)


	//Micro 创建微服务
	srv := micro.NewService(
		micro.Name("laracom.user.service"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterUserServiceHandler(srv.Server(),&handler.UserService{Repo:repo})

	if err := srv.Run(); err != nil{
		fmt.Println(err)
	}

}
