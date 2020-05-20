package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	proxy "go-grpc-client/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
)

var conn *grpc.ClientConn

func GetClientCreds() credentials.TransportCredentials {
	cert, err := tls.LoadX509KeyPair("cert/client.pem", "cert/client.key")
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("cert/ca.pem")
	if err != nil {
		log.Fatal(err)
	}
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "localhost",
		RootCAs:      certPool,
	})
	return creds
}

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
	studentInvoker()
}

func studentInvoker() {
	class := proxy.Class_C1901
	studentClient := proxy.NewStudentServiceClient(conn)
	studentRequest := proxy.StudentRequest{Class: class}
	studentResponse, err := studentClient.GetStudentsByClass(context.Background(), &studentRequest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s: %d\n", class, studentResponse.Students)

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
