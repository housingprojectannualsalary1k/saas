package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"saas/go-saas/models"
	"strconv"
)

type CourseController struct {
	CommonController
}

func NewCourseController() *CourseController{
	o := CourseController{}
	return &o
}

//客户管理列表
func (this *CourseController) View(ctx *gin.Context){
	//调用查询方法
	_Course := models.NewCourseModel()
	_Course.View()
}

//客户详情获取
func (this *CourseController) Find(ctx *gin.Context){
	ID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil{
		this.ResponseError(HttpStatusErr, err)
		return
	}

	//判断参数是否正确
	if ID < 0{
		this.ResponseError(HttpStatusErr, errors.New("缺少必要参数"))
		return
	}

	//调用查询方法
	_Course := models.NewCourseModel()
	_Course.ID = uint(ID)
	Course, err := _Course.Find()
	if err != nil{
		this.ResponseError(HttpStatusErr, err)
		return
	}

	this.Response(Course)
}

//添加客户
func (this *CourseController) Create(ctx *gin.Context){
	_Course := models.NewCourseModel()

	//绑定添加
	err := ctx.Bind(_Course)
	if err != nil{
		this.ResponseError(HttpStatusErr, err)
		return
	}

	//调用添加方法
	ID, err := _Course.Create()
	if err != nil{
		this.ResponseError(HttpStatusErr, err)
		return
	}

	this.Response(ID)
}

//修改客户
func (this *CourseController) Update(ctx *gin.Context){
	_Course := models.NewCourseModel()

	//绑定数据
	err := ctx.Bind(_Course)
	if err != nil{
		this.ResponseError(HttpStatusErr, err)
		return
	}

	//调用修改方法
	ID, err := _Course.Update()
	if err != nil{
		this.ResponseError(HttpStatusErr, err)
		return
	}

	this.Response(ID)
}

//客户删除
func (this *CourseController) Delete(ctx *gin.Context){
	ID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil{
		this.ResponseError(HttpStatusErr, err)
		return
	}

	//判断参数是否正确
	if ID < 0{
		this.ResponseError(HttpStatusErr, errors.New("缺少必要参数"))
		return
	}

	//删除
	_Course := models.NewCourseModel()
	_Course.ID = uint(ID)
	err = _Course.Delete()
	if err != nil{
		this.ResponseError(HttpStatusErr, err)
		return
	}

	this.Response(ID)
}