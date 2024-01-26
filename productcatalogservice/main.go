package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"net"
	"productcatalogservice/handler"
	pb "productcatalogservice/proto"
	"strconv"
)

const PORT = 50015          //端口
const ADDRESS = "127.0.0.1" //ip地址
func main() {
	ipport := ADDRESS + ":" + strconv.Itoa(PORT)
	//---------------注册到consul---------------
	// 初始化consul配置
	consulConfig := api.DefaultConfig()
	//创建consul对象
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		fmt.Println("consul创建对象报错:", err)
		return
	}
	//告诉consul即将注册到服务的信息
	reg := api.AgentServiceRegistration{
		Tags:    []string{"productcatalogservice"},
		Name:    "productcatalogservice",
		Address: ADDRESS,
		Port:    PORT,
	}
	//注册grpc服务到consul上
	err = consulClient.Agent().ServiceRegister(&reg)
	if err != nil {
		fmt.Println("consul注册grpc失败:", err)
		return
	}
	//---------------下面处理grpc代码---------------
	//初始化grpc对象
	grpcServer := grpc.NewServer()

	//注册服务
	pb.RegisterProductCatalogServiceServer(grpcServer, new(handler.ProductCatalogService))
	//设置监听
	listen, err := net.Listen("tcp", ipport)
	if err != nil {
		fmt.Println("监听报错：", err)
		return
	}
	defer listen.Close()
	fmt.Println("服务启动成功...")
	err = grpcServer.Serve(listen)
	if err != nil {
		fmt.Println("grpc服务启动报错:", err)
		return
	}
}
