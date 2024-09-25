package main

import (
	pb "grpc-user-serviceo/grpc-user-serviceo/pkg/grpc/user"
	userHandler "grpc-user-serviceo/internal/user/delivery/grpc"
	userRepo "grpc-user-serviceo/internal/user/repository/memory"
	userUsecase "grpc-user-serviceo/internal/user/usecase"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	repo := userRepo.NewUserRepository()
	usecase := userUsecase.NewUserUsecase(repo)
	handler := userHandler.NewUserHandler(usecase)

	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, handler)

	// Register reflection service on gRPC server for development use only
	reflection.Register(server)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("gRPC server is running on port 50051")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
