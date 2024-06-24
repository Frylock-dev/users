package stub

import (
	"github.com/Frylock-dev/users/internal/model"
	"github.com/brianvoe/gofakeit/v7"
	"time"
)

func NewUserInfoStub() *model.UserInfo {
	return &model.UserInfo{
		Email:             gofakeit.Email(),
		FirstName:         gofakeit.FirstName(),
		SecondName:        gofakeit.Name(),
		LastName:          gofakeit.LastName(),
		Phone:             gofakeit.Phone(),
		PassportNumber:    uint32(gofakeit.RandomUint(make([]uint, 0000000000))),
		PassportCode:      uint8(gofakeit.RandomUint(make([]uint, 000))),
		PassportIssueDate: time.Date(2018, time.November, 10, 23, 0, 0, 0, time.UTC),
		Birthday: gofakeit.DateRange(
			time.Date(1980, time.November, 10, 23, 0, 0, 0, time.UTC),
			time.Date(2005, time.November, 10, 23, 0, 0, 0, time.UTC),
		),
	}
}

func NewUserStub() *model.User {
	return &model.User{
		ID:   1,
		UUID: "44c16a1f-test-40ce-9b19-a5bf962c8bf1",
		Info: NewUserInfoStub(),
	}
}
