package server

import (
	"context"
	"reflect"
	"testing"

	"github.com/moms-spaghetti/followers/internal/models"
	pb "github.com/moms-spaghetti/followers/protobuf"
)

func TestServer_FollowUser(t *testing.T) {
	type fields struct {
		UnimplementedUserServiceServer pb.UnimplementedUserServiceServer
		MockStorage                    []*models.User
	}
	type args struct {
		in  *pb.FollowUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.FollowUserResponse
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				UnimplementedUserServiceServer: pb.UnimplementedUserServiceServer{},
				MockStorage: []*models.User{
					{
						ID: "id_1",
						Name: "user_1",
						FollowerIDs: []string{},
						FollowingIDs: []string{},
					},
					{
						ID: "id_2",
						Name: "user_2",
						FollowerIDs: []string{},
						FollowingIDs: []string{
							"id_3",
						},
					},
				},
			},
			args: args{
				in: &pb.FollowUserRequest{
					FollowingId: "id_1",
					FollowedById: "id_2",
				},
			},
			want: &pb.FollowUserResponse{
				FollowingIds: []string{
					"id_3",
					"id_1",
				},
				FollowedByIds: []string{},
			},
			wantErr: false,
		},
		{
			name: "followed user doesnt exist",
			fields: fields{
				UnimplementedUserServiceServer: pb.UnimplementedUserServiceServer{},
				MockStorage: []*models.User{
					{
						ID: "id_2",
						Name: "user_2",
						FollowerIDs: []string{},
						FollowingIDs: []string{},
					},
				},
			},
			args: args{
				in: &pb.FollowUserRequest{
					FollowingId: "id_1",
					FollowedById: "id_2",
				},
			},
			want: nil,
			wantErr: true,
		},
		{
			name: "follower user doesnt exist",
			fields: fields{
				UnimplementedUserServiceServer: pb.UnimplementedUserServiceServer{},
				MockStorage: []*models.User{
					{
						ID: "id_1",
						Name: "user_1",
						FollowerIDs: []string{},
						FollowingIDs: []string{},
					},
				},
			},
			args: args{
				in: &pb.FollowUserRequest{
					FollowingId: "id_1",
					FollowedById: "id_2",
				},
			},
			want: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedUserServiceServer: tt.fields.UnimplementedUserServiceServer,
				MockStorage:                    tt.fields.MockStorage,
			}
			got, err := s.FollowUser(context.TODO(), tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.FollowUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.FollowUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
