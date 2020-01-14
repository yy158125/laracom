package main

import (
	"context"
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	pb "github.com/yy158125/laracom/user-service/proto/user"
	"log"
	"os"
)

func main()  {

	// 初始化客户端服务，定义命令行参数标识
	service := micro.NewService(
		micro.Flags(
			cli.StringFlag{
				Name:"id",
				Usage:"your id",
			},
			cli.StringFlag{
				Name: "name",
				Usage: "your name",
			},
			cli.StringFlag{
				Name: "email",
				Usage: "Your email",
			},
			cli.StringFlag{
				Name: "password",
				Usage: "Your password",
			},
		),
	)

	// 远程服务客户端调用句柄
	client := pb.NewUserServiceClient("laracom.user.service",service.Client())

	service.Init(
		micro.Action(func(c *cli.Context) {
			name := c.String("name")
			email := c.String("email")
			password := c.String("password")
			log.Println(name,email,password)

			// 调用用户服务
			//r, err := client.Create(context.TODO(),&pb.User{
			//	Name:       name,
			//	Email:      email,
			//	Password:   password,
			//})
			//if err != nil {
			//	log.Fatalf("创建用户失败: %v",err)
			//}
			//
			//log.Printf("创建用户成功: %s", r.User.Id)

			all, err := client.GetAll(context.Background(),&pb.Request{})

			fmt.Println(err)

			fmt.Printf("%v \n %T \n",all.Users,all.Users)

			if err != nil {
				log.Fatalf("获取所有用户失败: %v", err)
			}
			//for _,user := range all.Users{
			//	log.Println(user)
			//}
			os.Exit(0)
		}),
	)

	if err := service.Run(); err != nil{
		log.Fatalf("用户客户端启动失败: %v", err)
	}

}
