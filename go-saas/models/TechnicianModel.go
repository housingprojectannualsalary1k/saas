package models

import "fmt"

type TechnicianModel struct {
	CommonModel
	Number		string		`gorm:"type:varchar(50);" json:"number"`
	Name		string		`gorm:"type:varchar(50);" json:"name"`
	Job		    string		`gorm:"type:varchar(30);" json:"Job"`
	Phone		string		`gorm:"type:char(30);" json:"Phone"`
	Status		int			`gorm:"type:tinyint(3); NOT NULL;" json:"status"`

}


func (this *TechnicianModel) TableName() string{
	return "Technician"
}


func NewTechnicianModel() *TechnicianModel{
	o := TechnicianModel{}
	return &o
}

func (this *TechnicianModel) View(){
	fmt.Println("123")
}
