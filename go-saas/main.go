package main

import (
	"fmt"
	"saas/go-saas/libs/config"
	"saas/go-saas/models"
)

var (
	cfg = "conf/config.toml"
)

func main(){
	fmt.Println("project initialization....")

	//加载配置文件
	config.GetFileAndLoad(cfg)

	//初始化数据库
	models.Init()

	//初始化路由
	ResgiterRoute()
}