package models

import "fmt"

type NodeModel struct {
	CommonModel
	Name		string		`gorm:"type:varchar(50);" json:"name"`
	Route		string		`gorm:"type:varchar(50);" json:"route"`
	Pid		    int		    `gorm:"type:int(11);" json:"pid"`
	Status		int			`gorm:"type:tinyint(3); NOT NULL;" json:"status"`
}


func (this *NodeModel) TableName() string{
	return "Node"
}


func NewNodeModel() *NodeModel{
	o := NodeModel{}
	return &o
}

func (this *NodeModel) View(){
	fmt.Println("123")
}
