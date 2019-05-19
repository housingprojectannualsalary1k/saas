package controllers

import (
	"github.com/gin-gonic/gin"
)

type RecordsController struct {
	CommonController
}

func NewRecordsController() *RecordsController{
	o := RecordsController{}
	return &o
}

func (this *RecordsController) View(ctx *gin.Context){
	this.Response("到店记录")
}