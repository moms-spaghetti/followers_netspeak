package server

import (
	"context"

	"github.com/moms-spaghetti/followers/internal/models"
	pb "github.com/moms-spaghetti/followers/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	MockStorage []*models.User
}

func (s *Server) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	for _, user := range s.MockStorage {
		if user.ID == in.GetId() {
			return &pb.GetUserResponse{
				User: models.ModelsUserToPB(user),
			}, nil
		}
	}

	return nil, status.Error(codes.NotFound, "user id not found")
}

func (s *Server) GetAllUsers(ctx context.Context, in *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	totalUsers := len(s.MockStorage)

	if totalUsers == 0 {
		return nil, status.Error(codes.Internal, "user db empty")
	}
	
	return &pb.GetAllUsersResponse{
		Users: models.ModelsUsersToPB(s.MockStorage),
	}, nil
}
