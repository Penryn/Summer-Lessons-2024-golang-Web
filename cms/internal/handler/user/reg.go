package user

import (
	"CMS/internal/service"
	"CMS/pkg/utils"

	"github.com/gin-gonic/gin"
)

type RegisterData struct{
	Username        string `json:"username" binding:"required"`  	//用户名
	Sex             string `json:"sex" binding:"required"`			//性别	
	PhoneNum        string `json:"phone_num" binding:"required"`	//电话号码
	Major           string `json:"major" binding:"required"`		//专业
	Password        string `json:"password" binding:"required"`		//密码
	ConfirmPassword string `json:"confirm_password" binding:"required"`//确认密码
}


func Register(c *gin.Context){
	var data RegisterData
	if err:=c.ShouldBindJSON(&data);err!=nil{
		utils.JsonFail(c,200501,"参数错误")
		return
	}
	if data.Password!=data.ConfirmPassword{
		utils.JsonFail(c,200502,"两次密码不一致")
		return
	}
	_,err:=service.GetUserByPhoneNum(data.PhoneNum)
	if err==nil{
		utils.JsonFail(c,200503,"该手机号已注册")
		return
	}
	err=service.Register(data.Username,data.Sex,data.PhoneNum,data.Major,data.Password)
	if err!=nil{
		utils.JsonFail(c,200504,"注册失败")
		return
	}
	utils.JsonSuccess(c,nil)

}