package main

import (
	"CMS/internal/middleware"
	database "CMS/internal/pkg/databse"
	"CMS/internal/router"
	"CMS/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db:=database.Init()
	service.ServiceInit(db)
	r := gin.Default()
	r.NoMethod(middleware.HandleNotFond)
	r.NoRoute(middleware.HandleNotFond)
	router.Init(r)
	err:=r.Run()
	if err!=nil{
		log.Fatal(err)
	}
}