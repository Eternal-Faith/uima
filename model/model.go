package model

type User struct {
	Id              int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Phone           string `gorm:"column:phone;type:varchar(11);unique;NOT NULL" json:"phone"` // 电话号码
	Nickname        string `gorm:"column:nickname;type:varchar(255)" json:"nickname"`          // 用户名
	Password        string `gorm:"column:password;type:varchar(255);NOT NULL" json:"password"` // 密码
	Avatar          string `gorm:"column:avatar;type:varchar(255)" json:"avatar"`              // 头像
	Gender          string `gorm:"column:gender;type:varchar(255)" json:"gender"`              // 性别
	Email           string `gorm:"column:email;type:varchar(255)" json:"email"`                // 邮箱
	Realname        string `gorm:"column:realname;type:varchar(255)" json:"realname"`          // 真实姓名
	Idcard          string `gorm:"column:idcard;type:varchar(255)" json:"idcard"`              // 身份证号
	Sha             string `gorm:"column:sha;type:varchar(255)" json:"sha"`
	Path            string `gorm:"column:path;type:varchar(255)" json:"path"`
	ConfirmPassword string `gorm:"-" json:"confirm_password"`
}

type Shop struct {
	Id           int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Picture      string `gorm:"column:picture;type:varchar(255)" json:"picture"`                     // 商店图片
	ShopName     string `gorm:"column:shop_name;type:varchar(255);NOT NULL" json:"shop_name"`        // 商店名称
	FieryNum     string `gorm:"column:fiery_num;type:varchar(255);unique;NOT NULL" json:"fiery_num"` // 火热指数
	OpeningTime  string `gorm:"column:opening_time;type:varchar(255)" json:"opening_time"`           // 营业时间
	CurrentNum   string `gorm:"column:current_num;type:varchar(255)" json:"current_num"`
	HotLine      string `gorm:"column:hot_line;type:varchar(255)" json:"hot_line"`
	ServiceIntro string `gorm:"column:service_intro;type:text" json:"service_intro"` // 服务介绍
	VipService   string `gorm:"column:vip_service;type:text" json:"vip_service"`     // vip服务
	ShouldKnow   string `gorm:"column:should_know;type:text" json:"should_know"`     // 需知
	Sha          string `gorm:"column:sha;type:varchar(255)" json:"sha"`
	Path         string `gorm:"column:path;type:varchar(255)" json:"path"`
}

func (m *Shop) TableName() string {
	return "shop"
}

type Broadcast struct {
	Id         int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Content    string `gorm:"column:content;type:text" json:"content"`
	Title      string `gorm:"column:title;type:varchar(255)" json:"title"`
	CreateTime string `gorm:"column:create_time;type:varchar(255)" json:"create_time"`
	Picture    string `gorm:"column:picture;type:varchar(255)" json:"picture"`
	Sha        string `gorm:"column:sha;type:varchar(255)" json:"sha"`
	Path       string `gorm:"column:path;type:varchar(255)" json:"path"`
}

func (m *Broadcast) TableName() string {
	return "broadcast"
}

type ScenicSpot struct {
	Id           int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Name         string `gorm:"column:name;type:varchar(255);NOT NULL" json:"name"`
	Picture      string `gorm:"column:picture;type:varchar(255);NOT NULL" json:"picture"`
	DownloadTime string `gorm:"column:download_time;type:varchar(255);NOT NULL" json:"download_time"`
	Sha          string `gorm:"column:sha;type:varchar(255)" json:"sha"`
	Path         string `gorm:"column:path;type:varchar(255)" json:"path"`
}

func (m *ScenicSpot) TableName() string {
	return "scenic_spot"
}

type Script struct {
	Id           int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	ScriptName   string `gorm:"column:script_name;type:varchar(255);NOT NULL" json:"script_name"`    // 剧本名称
	Introduction string `gorm:"column:introduction;type:varchar(1000);NOT NULL" json:"introduction"` // 剧本介绍
	BriefIntro   string `gorm:"column:brief_intro;type:varchar(255);NOT NULL" json:"brief_intro"`    // 剧本简介
	Time         string `gorm:"column:time;type:varchar(255);NOT NULL" json:"time"`                  // 剧本时长
	Place        string `gorm:"column:place;type:varchar(255);NOT NULL" json:"place"`                // 地点
	Avatar       string `gorm:"column:avatar;type:varchar(255);NOT NULL" json:"avatar"`              // 剧本封面
	Price        int    `gorm:"column:price;type:int(11);NOT NULL" json:"price"`                     // 价格
	Tag1         string `gorm:"column:tag1;type:varchar(100)" json:"tag1"`                           // 标签一
	Tag2         string `gorm:"column:tag2;type:varchar(100)" json:"tag2"`                           // 标签二
	Tag3         string `gorm:"column:tag3;type:varchar(100)" json:"tag3"`                           // 标签三
	Tag4         string `gorm:"column:tag4;type:varchar(100)" json:"tag4"`                           // 标签四
	Tag5         string `gorm:"column:tag5;type:varchar(100)" json:"tag5"`                           // 标签五
	Sha          string `gorm:"column:sha;type:varchar(255)" json:"sha"`
	Path         string `gorm:"column:path;type:varchar(255)" json:"path"`
}

type Scriptinterface struct {
	Id         int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	ScriptName string `gorm:"column:script_name;type:varchar(255);NOT NULL" json:"script_name"` // 剧本名称
	Avatar     string `gorm:"column:avatar;type:varchar(255);NOT NULL" json:"avatar"`           // 剧本封面
	BriefIntro string `gorm:"column:brief_intro;type:varchar(255);NOT NULL" json:"brief_intro"` // 剧本简介
}

type Place struct {
	Id            int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Name          string `gorm:"column:name;type:varchar(255);NOT NULL" json:"name"`                   // 地点名称
	Data          string `gorm:"column:data;type:varchar(1000);NOT NULL" json:"data"`                  // 地点资料
	Area          string `gorm:"column:area;type:varchar(255);NOT NULL" json:"area"`                   // 地点面积
	Visitor       string `gorm:"column:visitor;type:varchar(255);NOT NULL" json:"visitor"`             // 最大游客量
	Entertainment string `gorm:"column:entertainment;type:varchar(255);NOT NULL" json:"entertainment"` // 娱乐项目数量
	ScenicSpot    string `gorm:"column:scenic_spot;type:varchar(255);NOT NULL" json:"scenic_spot"`     // 特色景点数量
	Picture1      string `gorm:"column:picture1;type:varchar(255)" json:"picture1"`                    // 地点轮换封面1
	Sha1          string `gorm:"column:sha1;type:varchar(255)" json:"sha1"`
	Path1         string `gorm:"column:path1;type:varchar(255)" json:"path1"`
	Picture2      string `gorm:"column:picture2;type:varchar(255)" json:"picture2"` // 地点轮换封面2
	Sha2          string `gorm:"column:sha2;type:varchar(255)" json:"sha2"`
	Path2         string `gorm:"column:path2;type:varchar(255)" json:"path2"`
	Picture3      string `gorm:"column:picture3;type:varchar(255)" json:"picture3"` // 地点轮换封面3
	Sha3          string `gorm:"column:sha3;type:varchar(255)" json:"sha3"`
	Path3         string `gorm:"column:path3;type:varchar(255)" json:"path3"`
}

type ScriptCollections struct {
	Id        int `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	UsersId   int `gorm:"column:users_id;type:int(11);NOT NULL" json:"users_id"`     // 用户id
	ScriptsId int `gorm:"column:scripts_id;type:int(11);NOT NULL" json:"scripts_id"` // 剧本id
}

type ScriptOrders struct {
	Id          int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	UserId      int    `gorm:"column:user_id;type:int(11);NOT NULL" json:"user_id"`              // 用户id
	ScriptId    int    `gorm:"column:script_id;type:int(11);NOT NULL" json:"script_id"`          // 剧本id
	ScriptName  string `gorm:"column:script_name;type:varchar(255);NOT NULL" json:"script_name"` // 剧本名称
	Type        string `gorm:"column:type;type:varchar(255);NOT NULL" json:"type"`               // 类型(店铺/线上/线下)/剧本
	Price       int    `gorm:"column:price;type:int(11);NOT NULL" json:"price"`                  // 价格
	Createtime  string `gorm:"column:createtime;type:varchar(255);NOT NULL" json:"createtime"`   // 创建订单时间
	Paymenttime string `gorm:"column:paymenttime;type:varchar(255);NOT NULL" json:"paymenttime"` // 付款时间
	Avatar      string `gorm:"column:avatar;type:varchar(255)" json:"avatar"`                    // 订单图片
	Information string `gorm:"column:information;type:varchar(255);NOT NULL" json:"information"` // 订单内容
	Status      string `gorm:"column:status;type:varchar(255);NOT NULL" json:"status"`           // 订单状态
}

type ScriptAppointments struct {
	Id           int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	UsersId      int    `gorm:"column:users_id;type:int(11);NOT NULL" json:"users_id"`                // 用户id
	ScriptsId    int    `gorm:"column:scripts_id;type:int(11);NOT NULL" json:"scripts_id"`            // 剧本id
	ScriptsName  string `gorm:"column:scripts_name;type:varchar(255);NOT NULL" json:"scripts_name"`   // 剧本名称
	ScriptsCover string `gorm:"column:scripts_cover;type:varchar(255);NOT NULL" json:"scripts_cover"` // 剧本封面
	Time         string `gorm:"column:time;type:varchar(255)" json:"time"`                            // 预约时间
	Status       string `gorm:"column:status;type:varchar(255);NOT NULL" json:"status"`               // 预约状态(预约中/已预约/已完成)
}
