package models

import "fmt"

type CashbackModel struct {
	CommonModel
	Number		string		`gorm:"type:varchar(50);" json:"number"`
	Name		string		`gorm:"type:varchar(50);" json:"name"`
	Job		    string		`gorm:"type:varchar(30);" json:"Job"`
	Phone		string		`gorm:"type:char(30);" json:"Phone"`
	Status		int			`gorm:"type:tinyint(3); NOT NULL;" json:"status"`

}


func (this *CashbackModel) TableName() string{
	return "Cashback"
}


func NewCashbackModel() *CashbackModel{
	o := CashbackModel{}
	return &o
}

func (this *CashbackModel) View(){
	fmt.Println("123")
}
