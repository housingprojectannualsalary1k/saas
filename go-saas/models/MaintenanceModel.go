package models

import "fmt"

type MaintenanceModel struct {
	CommonModel
	Store       string      `gorm:"type:varchar(11);" json:"store"`
	StoreTime	int		    `gorm:"type:int(50);" json:"store_time"`
	Serve		string		`gorm:"type:varchar(50);" json:"serve"`
	Current		string		`gorm:"type:varchar(50);" json:"current"`
	CardTime	string		`gorm:"type:varchar(30);" json:"card_time"`
	Reduction	string		`gorm:"type:varchar(30);" json:"reduction"`
	Weight  	string		`gorm:"type:varchar(50);" json:"weight"`
	Vindicate   string     `gorm:"type:varchar(50);" json:"vindicate"`
	Present     string     `gorm:"type:varchar(50);" json:"present"`
	Stage       string     `gorm:"type:varchar(50);" json:"stage"`
	FinishTime   int       `gorm:"type:int(11);" json:"finish_time"`
}


func (this *MaintenanceModel) TableName() string{
	return "Maintenance"
}


func NewMaintenanceModel() *MaintenanceModel{
	o := MaintenanceModel{}
	return &o
}

func (this *MaintenanceModel) View(){
	fmt.Println("123")
}
