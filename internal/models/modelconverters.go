package models

import (
	pb "github.com/moms-spaghetti/followers/protobuf"
)

func ModelsUserToPB(in *User) *pb.User {
	followerIDs, followingIDs := ModelsFollowIDsToPB(in)
	return &pb.User{
		 Id: in.ID,
		 Name: in.Name,
		 FollowerIds: followerIDs,
		 FollowingIds: followingIDs,
	}
}

func ModelsUsersToPB(in []*User) []*pb.User {
	pbUsers := make([]*pb.User, len(in))

	for i, user := range in {
		pbUsers[i]= ModelsUserToPB(user)
	}

	return pbUsers
}

func ModelsFollowIDsToPB(in *User) ([]string, []string) {
	followerIDs := make([]string, len(in.FollowerIDs))
	followingIDs := make([]string, len(in.FollowingIDs))

	for i, id := range in.FollowerIDs {
		followerIDs[i] = id
	}

	for i, id := range in.FollowingIDs {
		followingIDs[i] = id
	}

	return followerIDs, followingIDs
}
