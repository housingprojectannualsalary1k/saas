package models

import "fmt"

type RoleModel struct {
	CommonModel
	RoleName		string		`gorm:"type:varchar(50);" json:"role_name"`
	Status		    int			`gorm:"type:tinyint(3); NOT NULL;" json:"status"`
}


func (this *RoleModel) TableName() string{
	return "Role"
}


func NewRoleModel() *RoleModel{
	o := RoleModel{}
	return &o
}

func (this *RoleModel) View(){
	fmt.Println("123")
}
