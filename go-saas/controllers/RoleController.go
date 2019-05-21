package controllers

import (
	"github.com/gin-gonic/gin"
)

type RoleController struct {
	CommonController
}

func NewRoleController() *RoleController{
	o := RoleController{}
	return &o
}

func (this *RoleController) View(ctx *gin.Context){
	this.Response("角色")
}