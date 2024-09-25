package grpc

import (
	"context"
	pb "grpc-user-serviceo/grpc-user-serviceo/pkg/grpc/user"
	"grpc-user-serviceo/internal/user/usecase"
	"log"
)

type UserHandler struct {
	usecase usecase.UserUsecase
	pb.UnimplementedUserServiceServer
}

func NewUserHandler(usecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

func (h *UserHandler) GetUser(ctx context.Context, req *pb.UserIDRequest) (*pb.UserResponse, error) {
	log.Printf("Recieved request to get user with ID: %d", req.Id)

	user, err := h.usecase.FetchUserByID(req.Id)
	if err != nil {
		log.Printf("Error fetching user with ID %d: %v", req.Id, err)
		return nil, err
	}

	log.Printf("Successful retrieval of the user: %v", user)

	return &pb.UserResponse{
		User: &pb.User{
			Id:      user.ID,
			Fname:   user.FName,
			City:    user.City,
			Phone:   user.Phone,
			Height:  user.Height,
			Married: user.Married,
		},
	}, nil
}

func (h *UserHandler) GetUsers(ctx context.Context, req *pb.UserIDsRequest) (*pb.UsersResponse, error) {
	users, err := h.usecase.FetchUsersByIDs(req.Ids)
	if err != nil {
		return nil, err
	}

	var pbUsers []*pb.User
	for _, user := range users {
		pbUsers = append(pbUsers, &pb.User{
			Id:      user.ID,
			Fname:   user.FName,
			City:    user.City,
			Phone:   user.Phone,
			Height:  user.Height,
			Married: user.Married,
		})
	}

	return &pb.UsersResponse{Users: pbUsers}, nil
}

func (h *UserHandler) SearchUsers(ctx context.Context, req *pb.SearchRequest) (*pb.UsersResponse, error) {
	criteria := map[string]interface{}{
		"city":    req.City,
		"phone":   req.Phone,
		"married": req.Married,
	}

	users, err := h.usecase.Search(criteria)
	if err != nil {
		return nil, err
	}

	var pbUsers []*pb.User
	for _, user := range users {
		pbUsers = append(pbUsers, &pb.User{
			Id:      user.ID,
			Fname:   user.FName,
			City:    user.City,
			Phone:   user.Phone,
			Height:  user.Height,
			Married: user.Married,
		})
	}

	return &pb.UsersResponse{Users: pbUsers}, nil
}
