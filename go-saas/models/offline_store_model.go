package models

import "fmt"

type OffLineStoreModel struct {
	CommonModel
	Name		string		`gorm:"type:varchar(50);" json:"name"`
	IsJoin		int			`gorm:"type:tinyint(3);" json:"is_join"`
	City		string		`gorm:"type:varchar(30);" json:"city"`
	Area		string		`gorm:"type:varchar(30);" json:"area"`
	Site		string		`gorm:"type:varchar(100);" json:"site"`
	Coordinate	string		`gorm:"type:varchar(30);" json:"coordinate"`
	OpenTime	string		`gorm:"type:varchar(50);" json:"open_time"`
	WorkTime	string		`gorm:"type:varchar(50);" json:"work_time"`
	Tel			string		`gorm:"type:varchar(20);" json:"tel"`
	UserName	string		`gorm:"type:varchar(50);" json:"username"`
	Password	string		`gorm:"type:char(32);" json:"tel"`
	Status		int			`gorm:"type:tinyint(3); NOT NULL;" json:"status"`
}


func (this *OffLineStoreModel) TableName() string{
	return "OfflineStore"
}


func NewOffLineStoreModel() *OffLineStoreModel{
	o := OffLineStoreModel{}
	return &o
}

func (this *OffLineStoreModel) View(){
	fmt.Println("123")
}
