package main

import (
	"context"
	pbHello "grpc-101/grpc/hello"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	log.Println("Connected to server")

	s := pbHello.NewUserServiceClient(conn)
	// resp, err := s.CreateUser(context.Background(), &pbHello.CreateUserRequest{
	// 	User: &pbHello.User{
	// 		Name:  "Ilham",
	// 		Email: "ilham@gmail.com",
	// 	},
	// })

	// if err != nil {
	// 	log.Fatalf("could not create user: %v", err)
	// }

	// log.Println("User created successfully", resp.User.Name, resp.User.Email)

	resp, err := s.SayHello(context.Background(), &pbHello.HelloRequest{Name: "ighfar hasbi"})
	if err != nil {
		panic(err)
	}

	log.Printf("Greeting: %s", resp.Message)

	gb, err := s.SayGoodbye(context.Background(), &pbHello.HelloRequest{Name: "ipay yapi"})
	if err != nil {
		panic(err)
	}

	log.Printf("Greeting: %s", gb.Message)
}
