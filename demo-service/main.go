package main

import (
	"context"
	pb "laracom/demo-service/proto/demo"
	"github.com/micro/go-micro"
	"log"
)

const (
	grpcPort = ":9999"
	httpPort = ":8000"
	appName = "Demo Service"
	address = "localhost:9999"
)

type DemoService struct {}

func (d *DemoService) SayHello(ctx context.Context,req *pb.DemoRequest, rsp *pb.DemoResponse) error {
	rsp.Text = "hello" + req.Name
	return nil
}

func main()  {
	// 通过-mode参数判断启动哪个模式的代码
	//mode := flag.String("mode","grpc","mode:grpc/http/client")
	//flag.Parse()
	//fmt.Println("run mode:", *mode)

	service := micro.NewService(
		micro.Name("laracom.demo.service"),
	)

	service.Init()

	pb.RegisterDemoServiceHandler(service.Server(),new(DemoService))

	if err := service.Run(); err != nil{
		log.Fatal(err)
	}






}
