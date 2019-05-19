package controllers

import (
	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	CommonController
}

func NewCustomerController() *CustomerController{
	o := CustomerController{}
	return &o
}

func (this *CustomerController) View(ctx *gin.Context){
	this.Response("客户管理")
}