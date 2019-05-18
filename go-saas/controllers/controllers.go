package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"saas/go-saas/models"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HttpStatus int

const (
	HttpStatusOk             HttpStatus = 200
	HttpStatusOkTwo			 HttpStatus = 201
	HttpStatusErr	 	     HttpStatus = 500
	HttpStatusHas	 	     HttpStatus = 501
	HttpStatusNotFond	     HttpStatus = 404
)

var (
	Ctx *gin.Context
	Session sessions.Session
	RedisClient *redis.Pool
)

type CommonController struct {
	ctx *gin.Context
}

func (this *CommonController) Response(data interface{}) {
	h := gin.H{
		"code": HttpStatusOk,
		"data": data}
	Ctx.JSON(http.StatusOK, h)
}

func (this *CommonController) ResponseError(code HttpStatus, e error) {
	h := gin.H{
		"code": code,
		"msg":  e.Error()}
	Ctx.JSON(http.StatusOK, h)
}

func (this CommonController) CreateTables(ctx *gin.Context){
	models.CreateTables()
	this.Response("ok")
}


func (this CommonController) md5V(str string) string  {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}