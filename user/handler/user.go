package handler

import (
	"context"
	"user/domain/model"
	"user/domain/service"
	pb "user/proto"
)

type User struct {
	UserDataService service.IUserDataService
}

func (u *User) Register(ctx context.Context, request *pb.UserRegisterRequest, response *pb.UserRegisterResponse) error {
	userRegister := &model.User{
		UserName:     request.UserName,
		FirstName:    request.FirstName,
		HashPassword: request.Pwd,
	}
	_, err := u.UserDataService.AddUser(userRegister)
	if err != nil {
		return err
	}
	response.Message = "添加成功"
	return nil
}

func (u *User) Login(ctx context.Context, request *pb.UserLoginRequest, response *pb.UserLoginResponse) error {
	isOk, err := u.UserDataService.CheckPwd(request.UserName, request.Pwd)
	if err != nil {
		return err
	}
	response.IsSuccess = isOk
	return nil
}

func (u *User) GetUserInfo(ctx context.Context, request *pb.UserInfoRequest, response *pb.UserInfoResponse) error {
	userInfo, err := u.UserDataService.FindUserByName(request.UserName)
	if err != nil {
		return err
	}
	response.UserId = userInfo.ID
	response.FirstName = userInfo.FirstName
	response.UserName = userInfo.UserName
	return nil
}
