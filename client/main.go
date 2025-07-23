package main

import (
	"context"
	pbHello "grpc-101/grpc/hello"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// untuk membuat grpc client
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	log.Println("Connected to server")

	s := pbHello.NewUserServiceClient(conn)

	// call greeting methods
	greetingService(s)

	// call CreateUser method
	createUser(s)

	// call GetUserById method
	getUserById(s)
}

func greetingService(s pbHello.UserServiceClient) {
	resp, err := s.SayHello(context.Background(), &pbHello.HelloRequest{Name: "maulana azis"})
	if err != nil {
		panic(err)
	}
	log.Printf("Greeting: %s", resp.Message)

	gb, err := s.SayGoodbye(context.Background(), &pbHello.HelloRequest{Name: "lay"})
	if err != nil {
		panic(err)
	}
	log.Printf("Greeting: %s", gb.Message)
}

func createUser(s pbHello.UserServiceClient) {
	resp, err := s.CreateUser(context.Background(), &pbHello.CreateUserRequest{
		User: &pbHello.User{
			Name:    "maulana azis",
			Email:   "aziss@gmail.com",
			Age:     25,
			Hobbies: []string{"watching movies", "coding", "sport"},
		},
	})
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}
	log.Println("User created successfully", resp.User)
}

func getUserById(s pbHello.UserServiceClient) {
	resp, err := s.GetUserById(context.Background(), &pbHello.GetUserByIdRequest{
		Id: 6,
	})
	if err != nil {
		log.Fatalf("could not get user by id: %v", err)
	}
	log.Println("User found successfully", resp.UsersDb)
}
