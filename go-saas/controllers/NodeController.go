package controllers

import (
	"github.com/gin-gonic/gin"
)

type NodeController struct {
	CommonController
}

func NewNodeController() *NodeController{
	o := NodeController{}
	return &o
}

func (this *NodeController) View(ctx *gin.Context){
	this.Response("维护记录")
}