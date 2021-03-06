package models

import (
	"fmt"
	"saas/go-saas/libs/config"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
	statusValue = []string {"禁用", "启用"}
)

const (
	STATAUS_ON 	= 1
	STATAUS_OFF = 0
)

type CommonModel struct {
	gorm.Model
}

func CreateTables(){
}

func Init(){
	var err error
	c := config.Get().Database

	db,err = gorm.Open("mysql",fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		c.Username,
		c.Password,
		c.Server,
		c.Port,
		c.Dbname,
		c.Charset,
	))
	if err != nil{
		panic(fmt.Sprintf("account slave db connect errr: %s", err))
	}

	fmt.Println("models init action success...")
}
