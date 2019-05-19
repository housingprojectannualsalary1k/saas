package controllers

import (
	"github.com/gin-gonic/gin"
)

type CourseController struct {
	CommonController
}

func NewCourseController() *CourseController{
	o := CourseController{}
	return &o
}

func (this *CourseController) View(ctx *gin.Context){
	this.Response("聊城记录")
}