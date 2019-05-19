package controllers

import (
	"github.com/gin-gonic/gin"
)

type SinformationController struct {
	CommonController
}

func NewSinformationController() *SinformationController{
	o := SinformationController{}
	return &o
}

func (this *SinformationController) View(ctx *gin.Context){
	this.Response("门店订单管理")
}