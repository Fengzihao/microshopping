package main

import (
	"currencyservice/handler"
	pb "currencyservice/proto"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

const PORT = 50012          //端口
const ADDRESS = "127.0.0.1" //ip地址

func main() {
	ipprot := ADDRESS + ":" + strconv.Itoa(PORT)
	//---------------注册到consul上-----------------
	//初始化consul配置
	consulConfig := api.DefaultConfig()
	//创建consul对象
	consulClient, consul_err := api.NewClient(consulConfig)
	if consul_err != nil {
		fmt.Println("consul创建对象报错:", consul_err)
		return
	}
	//告诉consul即将注册到服务的信息
	reg := api.AgentServiceRegistration{
		Tags:    []string{"currencyService"},
		Name:    "currencyService",
		Address: ADDRESS,
		Port:    PORT,
	}
	//注册grpc服务到consul上
	agent_err := consulClient.Agent().ServiceRegister(&reg)
	if agent_err != nil {
		fmt.Println("consul注册grpc失败:", agent_err)
		return
	}
	// ------------------grpc代码-----------------------
	//初始化grpc对象
	grpcServer := grpc.NewServer()
	//注册服务
	pb.RegisterCurrencyServiceServer(grpcServer, new(handler.CurrencyService))
	//设置监听
	listen, err := net.Listen("tcp", ipprot)
	if err != nil {
		fmt.Println("监听报错:", err)
		return
	}
	defer listen.Close()
	fmt.Println("服务启动成功。。。")
	//启动服务
	err = grpcServer.Serve(listen)
	if err != nil {
		fmt.Println("grpc服务启动报错:", err)
		return
	}
}
