package models

//疗程记录表
type CourseModel struct {
	CommonModel
	Store       int 	    `gorm:"type:int(10);" json:"store" form:"store"`	//服务门店
	StoreTime	int		    `gorm:"type:timestamp;" json:"store_time" form:"store_time"`	//到店时间
	Serve		string		`gorm:"type:varchar(50);" json:"serve" form:"serve"`		//服务项目
	Buy		    string		`gorm:"type:varchar(50);" json:"buy" form:"buy"`		//已购买
	NowStage	string	    `gorm:"type:varchar(50);" json:"now_stage" form:"now_stage"`	//当前疗程阶段
	Weight  	string		`gorm:"type:varchar(50);" json:"weight" form:"weight"`	//体重
	CardTime	string		`gorm:"type:varchar(30);" json:"card_time" form:"card_time"`	//美体卡时间
	Business    int         `gorm:"type:int(30);" json:"business" form:"business"`		//是否正常营业
    Vindicate   string     `gorm:"type:varchar(50);" json:"vindicate" form:"vindicate"`	//维护保养卡
    Present     string     `gorm:"type:varchar(50);" json:"present" form:"present"`	//当前体重
    Stage       string     `gorm:"type:varchar(50);" json:"stage" form:"stage"`		//阶段
    FinishTime   int       `gorm:"type:timestamp;" json:"finish_time" form:"finish_time"`	//完成时间
}

func (this *CourseModel) TableName() string{
	return "course"
}


func NewCourseModel() *CourseModel{
	o := CourseModel{}
	return &o
}

//获取数据
func (this *CourseModel) View()(_Course *CourseModel, err error){
	if res := db.Find(this); res.Error != nil{
		err = res.Error
	}

	return
}

//添加方法
func (this *CourseModel) Create()(ID uint, err error){
	if res := db.Create(this); res.Error != nil{
		err = res.Error
	}
	ID = this.ID
	return
}

//修改方法
func (this *CourseModel) Update()(ID uint, err error){
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
func (this *CourseModel) Find()(_Course CourseModel, err error){
	if res := db.Where("id = ?", this.ID).Take(&_Course); res.Error != nil{
		err = res.Error
	}
	return
}

//删除方法
func (this *CourseModel) Delete()(err error){
	//判断id 是否存在
	_course,err  := this.Find()
	if err != nil{
		return
	}

	//删除
	if res := db.Where("id = ?", this.ID).Delete(&_course); res.Error != nil{
		err = res.Error
	}
	return
}
