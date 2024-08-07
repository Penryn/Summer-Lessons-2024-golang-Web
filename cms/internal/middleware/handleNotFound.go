package middleware

import (
	"CMS/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleNotFond(c *gin.Context){
	utils.JsonResponse(c,404,200404,http.StatusText(http.StatusNotFound),nil)
}