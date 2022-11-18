package server

import (
	"context"
	"reflect"
	"testing"

	"github.com/moms-spaghetti/followers/internal/models"
	pb "github.com/moms-spaghetti/followers/protobuf"
)

func TestServer_GetUser(t *testing.T) {
	type fields struct {
		UnimplementedUserServiceServer pb.UnimplementedUserServiceServer
		MockStorage                    []*models.User
	}
	type args struct {
		in  *pb.GetUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.GetUserResponse
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
				},
			},
			args: args{
				in: &pb.GetUserRequest{
					Id: "id_1",
				},
			},
			want: &pb.GetUserResponse{
				User: &pb.User{
					Id: "id_1",
					Name: "user_1",
					FollowerIds: []string{},
					FollowingIds: []string{},
				},
			},
			wantErr: false,
		},
		{
			name: "no user",
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
				in: &pb.GetUserRequest{
					Id: "id_2",
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
			got, err := s.GetUser(context.TODO(), tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
