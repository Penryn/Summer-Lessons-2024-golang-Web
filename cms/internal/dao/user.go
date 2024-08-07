package dao

import (
	"CMS/internal/model"
	"context"
)

func (d *Dao) GetUserByPhoneNum(ctx context.Context,phone string) (*model.User,error){
	var user model.User
	err:=d.orm.WithContext(ctx).Where("phone_num=?",phone).First(&user).Error
	return &user,err
}

func (d *Dao) CreateUser(ctx context.Context,user *model.User) error{
	return d.orm.WithContext(ctx).Create(user).Error
}