package formatter

import (
	"review_movie/model"
	"time"
)

type UserFormatter struct {
	ID         int    `json:"id"`
	FullName       string `json:"fulName"`
	Role string `json:"role"`
	Email      string `json:"email"`
	
}

type UserFormatterRegister struct {
	ID         int    `json:"id"`
	FullName       string `json:"fullName"`
	Role string `json:"role"`
	Email      string `json:"email"`
}

type UserFormatterLogin struct {
	Token      string `json:"token"`
	Expire time.Time `json:"expire"`
	Code int `json:"code"`
}

func FormatUserRegisterResponse(user model.User) UserFormatterRegister {
	formatter := UserFormatterRegister{
		ID:         user.ID,
		FullName:       user.FullName,
		Role: user.Role,
		Email:      user.Email,
	}

	return formatter
}

func FormatUserResponse(user model.User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:         user.ID,
		Email:      user.Email,
		FullName:       user.FullName,
		Role: user.Role,
	}

	return formatter
}

func FormatUserLoginResponse(user model.User, token string, expire time.Time, statusCode int) UserFormatterLogin {
	formatter := UserFormatterLogin {
		Code: statusCode,
		Expire: expire,
		Token: token,
	}

	return formatter
}