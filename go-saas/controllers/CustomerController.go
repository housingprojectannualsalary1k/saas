package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"saas/go-saas/models"
	"strconv"
)

type CustomerController struct {
	CommonController
}

func NewCustomerController() *CustomerController{
	o := CustomerController{}
	return &o
}

//客户管理列表
func (this *CustomerController) View(ctx *gin.Context){
	//搜索条件
	_search := models.CustomerSearch{}

	//绑定搜索数据
	err := ctx.Bind(&_search)
	if err != nil{
		this.ResponseError(HttpStatusErr, err)
	}

	//调用查询方法
	_customer := models.NewCustomerModel()
	customer,err := _customer.View(&_search)
	if err != nil{
		this.ResponseError(HttpStatusErr, err)
		return
	}
	this.Response(customer)
}

//客户详情获取
func (this *CustomerController) Find(ctx *gin.Context){
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
	_customer := models.NewCustomerModel()
	_customer.ID = uint(ID)
	customer, err := _customer.Find()
	if err != nil{
		this.ResponseError(HttpStatusErr, err)
		return
	}

	this.Response(customer)
}

//添加客户
func (this *CustomerController) Create(ctx *gin.Context){
	_customer := models.NewCustomerModel()

	//绑定添加
	err := ctx.Bind(_customer)
	if err != nil{
		this.ResponseError(HttpStatusErr, err)
		return
	}

	//调用添加方法
	ID, err := _customer.Create()
	if err != nil{
		this.ResponseError(HttpStatusErr, err)
		return
	}

	this.Response(ID)
}

//修改客户
func (this *CustomerController) Update(ctx *gin.Context){
	_customer := models.NewCustomerModel()

	//绑定数据
	err := ctx.Bind(_customer)
	if err != nil{
		this.ResponseError(HttpStatusErr, err)
		return
	}

	//调用修改方法
	ID, err := _customer.Update()
	if err != nil{
		this.ResponseError(HttpStatusErr, err)
		return
	}

	this.Response(ID)
}

//客户删除
func (this *CustomerController) Delete(ctx *gin.Context){
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
	_customer := models.NewCustomerModel()
	_customer.ID = uint(ID)
	err = _customer.Delete()
	if err != nil{
		this.ResponseError(HttpStatusErr, err)
		return
	}

	this.Response(ID)
}