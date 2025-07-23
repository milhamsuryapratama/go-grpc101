package main

import (
	"context"
	"database/sql"
	"fmt"
	"grpc-101/grpc/hello"
	"log"
	"net"

	_ "github.com/lib/pq"

	"google.golang.org/grpc"
)

type greetingService struct {
	hello.UnimplementedUserServiceServer
}

func (s *greetingService) SayHello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {
	response := &hello.HelloResponse{
		Message: "Hello " + req.Name,
	}
	return response, nil
}

func (s *greetingService) SayGoodbye(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {
	response := &hello.HelloResponse{
		Message: "Goodbye " + req.Name,
	}
	return response, nil
}

func (s *greetingService) CreateUser(ctx context.Context, req *hello.CreateUserRequest) (*hello.CreateUserResponse, error) {
	response := &hello.CreateUserResponse{
		Message: "User created successfully",
		User:    req.User,
	}
	return response, nil
}

var db *sql.DB

func connectToDatabase() *sql.DB {
	connStr := "postgres://ighfarhasbiash:@localhost:5432/meeting_rooms_sk?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	return db
}

func (s *greetingService) GetUserById(ctx context.Context, req *hello.GetUserByIdRequest) (*hello.GetUserByIdResponse, error) {
	var user hello.UserDb
	row := db.QueryRow("SELECT id, name, email FROM user_try WHERE id = $1", req.Id)
	err := row.Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with id %d not found", req.Id)
		}
		return nil, fmt.Errorf("failed to get user: %v", err)
	}
	return &hello.GetUserByIdResponse{UsersDb: &user}, nil
}

func main() {
	db = connectToDatabase() // connect ke database
	defer db.Close()
	// start gRPC server disini
	port_server := ":8081"
	listener, err := net.Listen("tcp", port_server)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", port_server, err)
	}

	s := grpc.NewServer()

	hello.RegisterUserServiceServer(s, &greetingService{})
	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
