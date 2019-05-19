package controllers

import (
	"github.com/gin-gonic/gin"
)

type TechnicianController struct {
	CommonController
}

func NewTechnicianController() *TechnicianController{
	o := TechnicianController{}
	return &o
}

func (this *TechnicianController) View(ctx *gin.Context){
	this.Response("技师管理")
}