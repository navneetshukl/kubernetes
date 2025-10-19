package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDetail(ctx *gin.Context){
	params:=ctx.Param("name")
	log.Println("Params is ",params)
	ctx.JSON(http.StatusOK,gin.H{
		"status":"success",
		"name":params,
	})
}

func main(){
	router:=gin.Default()
	router.GET("/get/:name",GetDetail)
	router.Run()
}