package controllers

import (
	"github.com/gin-gonic/gin"
)

type MaintenanceController struct {
	CommonController
}

func NewMaintenanceController() *MaintenanceController{
	o := MaintenanceController{}
	return &o
}

func (this *MaintenanceController) View(ctx *gin.Context){
	this.Response("权限")
}