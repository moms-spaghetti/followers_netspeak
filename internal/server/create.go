package server

import (
	"context"

	"github.com/google/uuid"
	"github.com/moms-spaghetti/followers/internal/models"
	pb "github.com/moms-spaghetti/followers/protobuf"
)

func (s *Server) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	newUser := &models.User{
		ID: uuid.NewString(),
		Name: in.GetName(),
	}
	s.MockStorage = append(s.MockStorage, newUser)

	return &pb.CreateUserResponse{
		User: models.ModelsUserToPB(newUser),
	}, nil
}