package userapp

import (
	"context"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/entity"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/userdomain"
)

type UserApplication struct {
	userUsecase userdomain.IUserUsecase
	v1.UnimplementedUserServiceServer
}

func NewUserApplication(user userdomain.IUserUsecase) *UserApplication {
	return &UserApplication{
		userUsecase: user,
	}
}

func (s *UserApplication) CreateUser(ctx context.Context, in *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	userId, err := s.userUsecase.CreateUser(ctx, in.Mobile, in.Email, in.AccountId)
	if err != nil {
		return nil, err
	}
	return &v1.CreateUserResponse{
		Meta: &v1.Metadata{
			BizCode: 200,
			Message: "success",
		},
		UserId: userId,
	}, nil
}

func (s *UserApplication) UpdateUserInfo(ctx context.Context, in *v1.UpdateUserInfoRequest) (*v1.UpdateUserInfoResponse, error) {
	err := s.userUsecase.UpdateUserInfo(ctx, &entity.User{
		ID:              in.UserId,
		Name:            in.Name,
		Avatar:          in.Avatar,
		BackgroundImage: in.BackgroundImage,
		Signature:       in.Signature,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateUserInfoResponse{
		Meta: &v1.Metadata{
			BizCode: 200,
			Message: "success",
		},
	}, nil
}

func (s *UserApplication) GetUserInfo(ctx context.Context, in *v1.GetUserInfoRequest) (*v1.GetUserInfoResponse, error) {
	user, err := s.userUsecase.GetUserInfo(ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	return &v1.GetUserInfoResponse{
		User: user.ToUserResp(),
		Meta: &v1.Metadata{
			BizCode: 200,
			Message: "success",
		},
	}, nil
}