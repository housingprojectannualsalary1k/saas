package controllers

import (
	"github.com/gin-gonic/gin"
)

type CashbackController struct {
	CommonController
}

func NewCashbackController() *CashbackController{
	o := CashbackController{}
	return &o
}

func (this *CashbackController) View(ctx *gin.Context){
	this.Response("返现客户表")
}