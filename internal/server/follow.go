package server

import (
	"context"

	"github.com/moms-spaghetti/followers/internal/models"
	pb "github.com/moms-spaghetti/followers/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


func (s *Server) FollowUser(ctx context.Context, in *pb.FollowUserRequest) (*pb.FollowUserResponse, error) {
	mockStorageUserIDsMap := make(map[string]*models.User, len(s.MockStorage))

	for _, user := range s.MockStorage {
		mockStorageUserIDsMap[user.ID] = user
	}

	followedByUser, exists := mockStorageUserIDsMap[in.GetFollowedById()]
	if !exists {
		return nil, status.Error(codes.NotFound, "followed by user id not found")
	}

	followedUser, exists := mockStorageUserIDsMap[in.GetFollowingId()]
	if !exists {
		return nil, status.Error(codes.NotFound, "follower user id not found")
	}
	
	followedUser.FollowerIDs = append(followedUser.FollowerIDs, in.GetFollowedById())
	followedByUser.FollowingIDs = append(followedByUser.FollowingIDs, in.GetFollowingId() )

	for i, user := range s.MockStorage {
		if user.ID == followedUser.ID || user.ID == followedByUser.ID {
			s.MockStorage[i] = user
		}
	}

	followerIDs, followingIDs := models.ModelsFollowIDsToPB(followedByUser)

	return &pb.FollowUserResponse{
		FollowingIds: followingIDs,
		FollowedByIds: followerIDs,
	}, nil
}

func (s *Server) UnfollowUser(ctx context.Context, in *pb.UnfollowUserRequest) (*pb.FollowUserResponse, error) {
	mockStorageUserIDsMap := make(map[string]*models.User, len(s.MockStorage))

	for _, user := range s.MockStorage {
		mockStorageUserIDsMap[user.ID] = user
	}

	unfollowedByUser, exists := mockStorageUserIDsMap[in.GetUnfollowById()]
	if !exists {
		return nil, status.Error(codes.NotFound, "unfollowed by user id not found")
	}

	unfollowedUser, exists := mockStorageUserIDsMap[in.GetUnfollowId()]
	if !exists {
		return nil, status.Error(codes.NotFound, "unfollower user id not found")
	}

	unfollowedUserFollowers := make([]string, len(unfollowedUser.FollowerIDs) - 1)
	for i, id := range unfollowedUser.FollowerIDs {
		if id != in.GetUnfollowById() {
			unfollowedUserFollowers[i] = id
		}
	} 

	unfollowedUser.FollowerIDs = unfollowedUserFollowers

	unfollowedByUserFollowing := make([]string, len(unfollowedByUser.FollowingIDs) - 1)
	for i, id := range unfollowedByUser.FollowingIDs {
		if id != in.GetUnfollowId() {
			unfollowedByUserFollowing[i] = id
		}
	}

	unfollowedByUser.FollowingIDs = unfollowedByUserFollowing

	for i, user := range s.MockStorage {
		if user.ID == unfollowedUser.ID || user.ID == unfollowedByUser.ID {
			s.MockStorage[i] = user
		}
	}

	followerIDs, followingIDs := models.ModelsFollowIDsToPB(unfollowedByUser)

	return &pb.FollowUserResponse{
		FollowingIds: followingIDs,
		FollowedByIds: followerIDs,
	}, nil
}