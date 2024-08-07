package router

import (
	"CMS/internal/handler/contact"
	"CMS/internal/handler/user"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine){
	const pre = "/api"

	api:=r.Group(pre)
	{
		api.POST("register",user.Register)
		api.POST("login",user.Login)

		c:=api.Group("/contact")
		{
			c.POST("",contact.CreateContact)
			c.PUT("",contact.UpdateContact)
			c.DELETE("",contact.DeleteContact)
			c.GET("",contact.GetContact)

		}
	}
	
}