package models

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
	VestInShop		string		`gorm:"type:varchar(50);" json:"vest_in_shop" form:"vest_in_shop"`	//归属门店
	ServerInShop	string		`gorm:"type:varchar(50);" json:"server_in_shop" form:"server_in_shop"`	//服务门店
	AccountNumber	string		`gorm:"type:varchar(50);" json:"account_number" form:"account_number"`	//卡面账号
	Pid				int			`gorm:"type:int(10);" json:"pid" form:"pid"`		//	推荐人id
	Comments		string		`gorm:"type:varchar(50);" json:"comments" form:"comments"`	//备注
}

	CustomerSearch struct {
		Name 		string
		Sex  		int
		Date  		string
		Zcountmin 	string
		Zcountmax  	string
		Lcountmin  	string
		Lcountmax  	string
		Type 		int
		Page		int
	}
)



func (this *CustomerModel) TableName() string{
	return "customer"
}

func NewCustomerModel() *CustomerModel{
	o := CustomerModel{}
	return &o
}

func (this *CustomerModel) View(_search *CustomerSearch)(_customer *CustomerModel, err error){
	if res := db.Find(this); res.Error != nil{
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
