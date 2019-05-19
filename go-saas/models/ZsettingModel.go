package models

import "fmt"

type ZsettingModel struct {
	CommonModel
	Name		string		`gorm:"type:varchar(50);" json:"name"`
	Type		int			`gorm:"type:tinyint(3);" json:"type"`
	Opening		string		`gorm:"type:varchar(30);" json:"opening"`
	Rank		string		`gorm:"type:varchar(30);" json:"pank"`

}


func (this *ZsettingModel) TableName() string{
	return "Zsetting"
}


func NewZsettingModel() *ZsettingModel{
	o := ZsettingModel{}
	return &o
}

func (this *ZsettingModel) View(){
	fmt.Println("123")
}
