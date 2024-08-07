package service

import "CMS/internal/model"

func GetUserByPhoneNum(phoneNum string) (*model.User, error) {
	return d.GetUserByPhoneNum(ctx,phoneNum)
}

func Register(username string,sex string ,phone string,major string,password string) error {
	return d.CreateUser(ctx,&model.User{
		Username: username,
		Sex: sex,
		PhoneNum: phone,
		Major: major,
		Password: password,
	})
}