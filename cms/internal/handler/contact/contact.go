package contact

import (
	"CMS/internal/model"
	"CMS/internal/service"
	"CMS/pkg/utils"

	"github.com/gin-gonic/gin"
)

type CreateContactData struct {
	OwnerID   uint   `json:"owner_id" binding:"required"`
	StudentID string `json:"student_id"`
	Name      string `json:"name" binding:"required"`
	Sex       string `json:"sex"`
	PhoneNum  string `json:"phone_num" binding:"required"`
	Major     string `json:"major"`
	Blacklist bool   `json:"blacklist"`
}

func CreateContact(c *gin.Context) {
	var data CreateContactData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonFail(c, 200501, "参数错误")
		return
	}
	err = service.CreateContact(data.OwnerID, data.StudentID, data.Name, data.Sex, data.PhoneNum, data.Major, data.Blacklist)
	if err != nil {
		utils.JsonFail(c, 200504, "添加联系人失败")
		return
	}
	utils.JsonSuccess(c, nil)

}

type UpdateContactData struct {
	ID        uint   `json:"id" binding:"required"`
	StudentID string `json:"student_id"`
	Name      string `json:"name" binding:"required"`
	Sex       string `json:"sex"`
	PhoneNum  string `json:"phone_num" binding:"required"`
	Major     string `json:"major"`
	Blacklist bool   `json:"blacklist"`
}

func UpdateContact(c *gin.Context){
	var data UpdateContactData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonFail(c, 200501, "参数错误")
		return
	}
	err=service.UpdateContact(data.ID, data.StudentID, data.Name, data.Sex, data.PhoneNum, data.Major, data.Blacklist)
	if err != nil {
		utils.JsonFail(c, 200504, "更新联系人失败")
		return
	}
	utils.JsonSuccess(c, nil)
}

type DeleteContactData struct{
	ID uint `json:"id" binding:"required"`
}

func DeleteContact(c *gin.Context){
	var data DeleteContactData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonFail(c, 200501, "参数错误")
		return
	}
	err =service.DeleteContact(data.ID)
	if err!=nil{
		utils.JsonFail(c,200504,"删除联系人失败")
		return
	}
	utils.JsonSuccess(c,nil)
}

type GetContactData struct{
	OwnerID uint `form:"owner_id" binding:"required"`
}

func GetContact(c *gin.Context){
	var data GetContactData
	err := c.ShouldBindQuery(&data)
	if err != nil {
		utils.JsonFail(c, 200501, "参数错误")
		return
	}

	var contactList []model.Contact
	contactList,err=service.GetContactList(data.OwnerID)
	if err!=nil{
		utils.JsonFail(c,200504,"获取列表错误")
	}
	utils.JsonSuccess(c,gin.H{
		"contact_list":contactList,
		"num":len(contactList),
	})
}