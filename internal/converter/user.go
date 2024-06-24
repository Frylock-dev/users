package converter

import (
	"github.com/Frylock-dev/users/internal/model"
	"github.com/Frylock-dev/users/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromUserToGRPC(user *model.User) *user_v1.User {
	return &user_v1.User{
		Id:   uint64(user.ID),
		Uuid: user.UUID,
		Info: FromUserInfoToGRPC(user.Info),
	}
}

func FromUserInfoToGRPC(userInfo *model.UserInfo) *user_v1.UserInfo {
	return &user_v1.UserInfo{
		Phone:             userInfo.Phone,
		Email:             userInfo.Email,
		FirstName:         userInfo.FirstName,
		SecondName:        userInfo.SecondName,
		LastName:          userInfo.LastName,
		PassportNumber:    userInfo.PassportNumber,
		PassportCode:      uint32(userInfo.PassportCode),
		PassportIssueDate: timestamppb.New(userInfo.PassportIssueDate),
		Birthday:          timestamppb.New(userInfo.Birthday),
	}
}

func FromUserInfoToService(userInfo *user_v1.UserInfo) *model.UserInfo {
	return &model.UserInfo{
		Phone:             userInfo.Phone,
		Email:             userInfo.Email,
		FirstName:         userInfo.FirstName,
		SecondName:        userInfo.SecondName,
		LastName:          userInfo.LastName,
		PassportNumber:    userInfo.PassportNumber,
		PassportCode:      uint8(userInfo.PassportCode),
		PassportIssueDate: userInfo.PassportIssueDate.AsTime(),
		Birthday:          userInfo.Birthday.AsTime(),
	}
}
