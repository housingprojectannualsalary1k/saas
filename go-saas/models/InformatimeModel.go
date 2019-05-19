package models

import "fmt"

type InformatimeModel struct {
	CommonModel
	Name		string		`gorm:"type:varchar(50);" json:"name"`
	Across		string		`gorm:"type:varchar(30);" json:"across"`
	Present		string		`gorm:"type:varchar(30);" json:"present"`
	Before		string		`gorm:"type:varchar(30);" json:"before"`
	Service		string		`gorm:"type:varchar(30);" json:"service"`
	Sex		    string		`gorm:"type:varchar(30);" json:"sex"`
	Reducing	string		`gorm:"type:varchar(30);" json:"reducing"`
	Ideality	string		`gorm:"type:varchar(30);" json:"ideality"`
	Now		    string		`gorm:"type:varchar(30);" json:"now"`
	Course		string		`gorm:"type:varchar(30);" json:"course"`
	Treatment	int		    `gorm:"type:int(11);" json:"treatment"`
	AcrossTime	int		    `gorm:"type:int(11);" json:"across_time"`

}


func (this *InformatimeModel) TableName() string{
	return "Informatime"
}


func NewInformatimeModel() *InformatimeModel{
	o := InformatimeModel{}
	return &o
}

func (this *InformatimeModel) View(){
	fmt.Println("123")
}
