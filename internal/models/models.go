package models

type User struct {
	ID string
	Name string
	FollowerIDs []string
	FollowingIDs []string
}
