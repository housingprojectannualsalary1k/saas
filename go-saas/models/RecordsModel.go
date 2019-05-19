package models

import "fmt"

type RecordsModel struct {
	CommonModel
	Store		string		`gorm:"type:varchar(50);" json:"store"`
	Time		string		`gorm:"type:datetime;" json:"time"`
	Num		    int		    `gorm:"type:int(11);" json:"num"`

}


func (this *RecordsModel) TableName() string{
	return "Records"
}


	func NewRecordsModel() *RecordsModel{
	o := RecordsModel{}
	return &o
}

func (this *RecordsModel) View(){
	fmt.Println("123")
}
