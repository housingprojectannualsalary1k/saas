package controllers

import (
	"github.com/gin-gonic/gin"
)

type ZsettingController struct {
	CommonController
}

func NewZsettingController() *ZsettingController{
	o := ZsettingController{}
	return &o
}

func (this *ZsettingController) View(ctx *gin.Context){
	this.Response("区域管理")
}