package handler

import (
	"context"
	"user/domain/service"
	pb "user/proto"
)

type User struct {
	UserDataService service.IUserDataService
}

func (u *User) Register(ctx context.Context, request *pb.UserRegisterRequest, response *pb.UserRegisterResponse) error {
	//TODO implement me
	panic("implement me")
}

func (u *User) Login(ctx context.Context, request *pb.UserLoginRequest, response *pb.UserLoginResponse) error {
	//TODO implement me
	panic("implement me")
}

func (u *User) GetUserInfo(ctx context.Context, request *pb.UserInfoRequest, response *pb.UserInfoResponse) error {
	//TODO implement me
	panic("implement me")
}
