package controllers

import (
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	CommonController
}

func NewProductController() *ProductController{
	o := ProductController{}
	return &o
}

func (this *ProductController) View(ctx *gin.Context){
	this.Response("产品管理")
}