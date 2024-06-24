package converter

import (
	"github.com/Frylock-dev/users/internal/model"
	repoModel "github.com/Frylock-dev/users/internal/repository/postgres/user/model"
)

func ToUserInfoFromService(user *model.UserInfo) *repoModel.UserInfo {
	return &repoModel.UserInfo{
		Phone:             user.Phone,
		Email:             user.Email,
		FirstName:         user.FirstName,
		SecondName:        user.SecondName,
		LastName:          user.LastName,
		PassportNumber:    user.PassportNumber,
		PassportCode:      user.PassportCode,
		PassportIssueDate: user.PassportIssueDate,
		Birthday:          user.Birthday,
	}
}

func ToUserInfoFromRepo(info *repoModel.UserInfo) *model.UserInfo {
	return &model.UserInfo{
		Phone:             info.Phone,
		Email:             info.Email,
		FirstName:         info.FirstName,
		SecondName:        info.SecondName,
		LastName:          info.LastName,
		PassportNumber:    info.PassportNumber,
		PassportCode:      info.PassportCode,
		PassportIssueDate: info.PassportIssueDate,
		Birthday:          info.Birthday,
	}
}

func ToUserFromRepo(user *repoModel.User) *model.User {
	return &model.User{
		ID:   user.ID,
		UUID: user.UUID,
		Info: ToUserInfoFromRepo(user.Info),
	}
}
