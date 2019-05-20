package models

import "fmt"

type ProductModel struct {
	CommonModel
	Name		string		`gorm:"type:varchar(50);" json:"name"`
	Course		int			`gorm:"type:tinyint(3);" json:"course"`
	Price		string		`gorm:"type:varchar(30);" json:"price"`
	ReturnDay	string		`gorm:"type:varchar(30);" json:"return_day"`
	ReturnMoney	string		`gorm:"type:varchar(100);" json:"return_money"`
	Order	    string		`gorm:"type:varchar(30);" json:"order"`
	Putaway 	string		`gorm:"type:varchar(50);" json:"putaway"`

}


func (this *ProductModel) TableName() string{
	return "Product"
}


func NewProductModel() *ProductModel{
	o := ProductModel{}
	return &o
}

func (this *ProductModel) View(){
	fmt.Println("123")
}
