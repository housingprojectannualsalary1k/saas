package controllers

import (
	"github.com/gin-gonic/gin"
)

type OffLineStoreController struct {
	CommonController
}

func NewOffLineStoreController() *OffLineStoreController{
	o := OffLineStoreController{}
	return &o
}

func (this *OffLineStoreController) Copy() *OffLineStoreController{
	cp := *this
	cp.ctx = nil
	return &cp
}

func (this *OffLineStoreController) View(ctx *gin.Context){
	this.Response("你好啊")
}