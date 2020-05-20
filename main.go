package main

import (
	"context"
	"encoding/json"
	"fmt"
	proxy "go-grpc-client/api"
	"google.golang.org/grpc"
	"log"
)

var conn *grpc.ClientConn

func main() {
	//客户端拨号RPC服务器，返回连接对象
	//conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	var err error
	conn, err = grpc.Dial(":8081", grpc.WithTransportCredentials(GetClientCreds()))

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	helloInvoker()
	productInvoker()
}

func helloInvoker() {
	//在连接对象创建服务器上的服务对象的客户端（服务代理）
	helloClient := proxy.NewHelloServiceClient(conn)
	helloRequest := proxy.HelloRequest{Name: "Salmon"}
	//通过服务代理对象调用远程服务对象方法
	helloResponse, err := helloClient.SayHello(context.Background(), &helloRequest)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("返回response:", helloResponse.Message)
}
func productInvoker() {
	//返回Product的列表
	productClient := proxy.NewProductServiceClient(conn)
	querySize := proxy.QuerySize{Size: 10}
	productResponseList, err := productClient.GetProductList(context.Background(), &querySize)
	if err != nil {
		log.Println(err)
	}

	//把List转json
	bytes, err := json.Marshal(productResponseList.ProductList)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("返回ProductList:", string(bytes))
}
