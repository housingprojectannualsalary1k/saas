package models

import "fmt"

type CustomerModel struct {
	CommonModel
	MemberId	int		    `gorm:"type:int(50);" json:"member_id"`
	HeaderImg   int			`gorm:"type:varchar(300);" json:"header_img"`
	Name  		string		`gorm:"type:varchar(30);" json:"name"`
	Tel		    string		`gorm:"type:varchar(30);" json:"tel"`
	Sex		    string		`gorm:"type:varchar(100);" json:"sex"`
	Industry	string		`gorm:"type:varchar(30);" json:"industry"`
	Birthday	string		`gorm:"type:varchar(50);" json:"birthday"`
	Record	    string		`gorm:"type:varchar(50);" json:"record"`
	Form		string		`gorm:"type:varchar(20);" json:"form"`
	Height	    string		`gorm:"type:varchar(50);" json:"height"`
	BeforWeight	string		`gorm:"type:varchar(32);" json:"befor_weight"`
	AtrWeight	string		`gorm:"type:varchar(32);" json:"atr_weight"`
	StandardWeight	string	`gorm:"type:varchar(32);" json:"standard_weight"`
	StarTime	string		`gorm:"type:datetime;" json:"star_time"`
	VestInShop	string		`gorm:"type:varchar(50);" json:"vest_in_shop"`
	ServerInShop	string	`gorm:"type:varchar(50);" json:"server_in_shop"`
	AccountNumber	string	`gorm:"type:varchar(50);" json:"account_number"`
	PId	            int 	`gorm:"type:int(10);" json:"p_id"`
	Comments	  string	`gorm:"type:varchar(50);" json:"comments"`


}


func (this *CustomerModel) TableName() string{
	return "Customer"
}


func NewCustomerModel() *CustomerModel{
	o := CustomerModel{}
	return &o
}

func (this *CustomerModel) View(){
	fmt.Println("123")
}
