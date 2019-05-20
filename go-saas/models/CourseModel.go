package models

import "fmt"

type CourseModel struct {
	CommonModel
	Store       string      `gorm:"type:varchar(11);" json:"store"`
	StoreTime	int		    `gorm:"type:int(50);" json:"store_time"`
	Serve		string		`gorm:"type:varchar(50);" json:"serve"`
	Buy		    string		`gorm:"type:varchar(50);" json:"buy"`
	NowStage	string	    `gorm:"type:varchar(50);" json:"now_stage"`
	Weight  	string		`gorm:"type:varchar(50);" json:"weight"`
	CardTime	string		`gorm:"type:varchar(30);" json:"card_time"`
	Business    int         `gorm:"type:int(30);" json:"business"`
	Vindicate   string     `gorm:"type:varchar(50);" json:"vindicate"`
	Present     string     `gorm:"type:varchar(50);" json:"present"`
	Stage       string     `gorm:"type:varchar(50);" json:"stage"`
	FinishTime   int       `gorm:"type:int(11);" json:"finish_time"`
}


func (this *CourseModel) TableName() string{
	return "Course"
}


func NewCourseModel() *CourseModel{
	o := CourseModel{}
	return &o
}

func (this *CourseModel) View(){
	fmt.Println("123")
}
