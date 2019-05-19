package controllers

import (
	"github.com/gin-gonic/gin"
)

type InformatimeController struct {
	CommonController
}

func NewInformatimeController() *InformatimeController{
	o := InformatimeController{}
	return &o
}

func (this *InformatimeController) View(ctx *gin.Context){
	this.Response("全部串店信息")
}