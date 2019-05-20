package models

import "fmt"

type SinformationModel struct {
	CommonModel
	CId         int         `gorm:"type:int(11);" json:"c_id"`
	OrderNumber	string		`gorm:"type:varchar(50);" json:"order_number"`
	Time		string		`gorm:"type:varchar(50);" json:"name"`
	PId		    int		    `gorm:"type:int(11);" json:"p_id"`
	OsId		int		    `gorm:"type:int(11);" json:"os_id"`
	IsEffective	string		`gorm:"type:tinyint(3);" json:"is_effective"`
	Discounts	string		`gorm:"type:varchar(30);" json:"discounts"`
	NowIntegral string      `gorm:"type:varchar(30);" json:"now_integral"`
	NowMoney    float64     `gorm:"type:float(11,0);" json:"now_money"`
}


func (this *SinformationModel) TableName() string{
	return "Sinformation"
}


func NewSinformationModel() *SinformationModel{
	o := SinformationModel{}
	return &o
}

func (this *SinformationModel) View(){
	fmt.Println("123")
}
