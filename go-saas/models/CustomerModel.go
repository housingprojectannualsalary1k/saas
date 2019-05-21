package models

import (
	"strconv"
)

type (
	CustomerModel struct {
	CommonModel
	MemberId		int			`gorm:"type:int(10);"json:"member_id" form:"member_id"`	//会员卡号
	HeaderImg		string		`gorm:"type:varchar(300);"json:"header_img" form:"header_img"`	//客户头像
	Name			string		`gorm:"type:varchar(20);" json:"name" form:"name"`	//客户姓名
	Tel				string		`gorm:"type:varchar(20);" json:"tel" form:"tel"`	//客户手机
	Sex				int			`gorm:"type:tinyint(3);" json:"sex" form:"sex"`	//性别
	Industry		string		`gorm:"type:varchar(50);" json:"industry" form:"industry"`	//行业
	Birthday		string		`gorm:"type:varchar(50);" json:"birthday" form:"birthday"`	//出生日期
	Record			string		`gorm:"type:varchar(50);" json:"area" form:"area"`		//跨店记录
	Form			string		`gorm:"type:varchar(100);" json:"form" form:"form"`		//来店途径
	Height			string		`gorm:"type:varchar(50);" json:"height" form:"height"`	//身高
	BeforWeight		string		`gorm:"type:varchar(50);" json:"befor_weight" form:"befor_weight"`	//减前记录（斤）
	AtrWeight		string		`gorm:"type:varchar(50);" json:"atr_weight" form:"atr_weight"`	//美学体重（斤）
	StandardWeight	string		`gorm:"type:varchar(50);" json:"standard_weight" form:"standard_weight"`	//标准体重（斤）
	StarTime		string		`gorm:"type:datetime;" json:"star_time" form:"star_time"`		//开始减重时间
	VestInShop		int			`gorm:"type:int(10);" json:"vest_in_shop" form:"vest_in_shop"`	//归属门店id
	ServerInShop	int			`gorm:"type:int(10);" json:"server_in_shop" form:"server_in_shop"`	//服务门店id
	AccountNumber	string		`gorm:"type:varchar(50);" json:"account_number" form:"account_number"`	//卡面账号
	Pid				int			`gorm:"type:int(10);" json:"pid" form:"pid"`		//	推荐人id
	Comments		string		`gorm:"type:varchar(50);" json:"comments" form:"comments"`	//备注
}

	CustomerSearch struct {
		Name 		string
		Sex  		int
		Date  		int
		Zcountmin 	string
		Zcountmax  	string
		Lcountmin  	string
		Lcountmax  	string
		Type 		int
		Page		int
	}
)

var (
	LIMIT = 10

	NOW = 1
	LASTMONTH = 2
	LASTYEAR = 3
)

func (this *CustomerModel) TableName() string{
	return "customer"
}

func NewCustomerModel() *CustomerModel{
	o := CustomerModel{}
	return &o
}

//获取数据
func (this *CustomerModel) View(_search *CustomerSearch)(_customer []CustomerModel, err error){
	offset := (_search.Page - 1) * LIMIT
	sex := strconv.Itoa(_search.Sex)
	if sex == "0"{
		sex = "1,2"
	}
	

	if res := db.Where("name LIKE ? and sex in (?)", "%" + _search.Name + "%", sex).Offset(offset).Limit(LIMIT).Find(&_customer); res.Error != nil{
		err = res.Error
	}

	return
}

//获取总条数
func (this *CustomerModel) Count()(count int, err error){
	if res := db.Find(this).Count(&count); res.Error != nil{
		err = res.Error
	}

	return
}

//添加方法
func (this *CustomerModel) Create()(ID uint, err error){
	if res := db.Create(this); res.Error != nil{
		err = res.Error
	}
	ID = this.ID
	return
}

//修改方法
func (this *CustomerModel) Update()(ID uint, err error){
	//判断id是否存在
	_,err  = this.Find()
	if err != nil{
		return
	}

	//修改
	if res := db.Where("id = ?", this.ID).Save(this); res.Error != nil{
		err = res.Error
	}
	ID = this.ID
	return
}

//查询详情方法
func (this *CustomerModel) Find()(_customer CustomerModel, err error){
	if res := db.Where("id = ?", this.ID).First(&_customer); res.Error != nil{
		err = res.Error
	}
	return
}

//删除方法
func (this *CustomerModel) Delete()(err error){
	//判断id 是否存在
	_customer,err  := this.Find()
	if err != nil{
		return
	}

	//删除
	if res := db.Where("id = ?", this.ID).Delete(_customer); res.Error != nil{
		err = res.Error
	}
	return
}
