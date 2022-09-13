package model

import (
	"fmt"
	"log"
	"strconv"
)

//注册账号
func Register(phone string, password string) string {
	user := User{Phone: phone, Password: password}
	if err := DB.Table("users").Create(&user).Error; err != nil {
		fmt.Println("注册出错" + err.Error()) //err.Error打印错误
		return err.Error()
	}
	Id := strconv.Itoa(user.Id)
	return Id
}

//获取用户信息
func GetUserId(phone string) (User, error) {
	var user User
	if err := DB.Table("users").Where("phone=?", phone).Find(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

//防止电话重复绑定331,如果有这条数据则说明该电话号码已被注册
func IfExistUserPhone(phone string) (error, int) {
	var temp User
	if err := DB.Table("users").Where("phone = ?", phone).Find(&temp).Error; temp.Phone == "" { //电话为空说明数据库中没有这个电话
		log.Println(err) //比fmt.Println多时间戳
		// fmt.Println("hh", err)
		return err, 1
	}
	fmt.Println(temp)
	return nil, 0
}

//验证用户是否存在29
func VerifyPhone(phone string) bool {
	var user = make([]User, 1) //分配一个结构体
	if err := DB.Table("users").Where("phone=?", phone).Find(&user).Error; err != nil {
		log.Println(err.Error())
		return false
	}
	if len(user) != 1 {
		fmt.Println(len(user))
		return true
	}
	return false
}

//验证密码
func VerifyPassword(phone string, password string) bool {
	var user User
	if err := DB.Table("users").Where("phone = ? and password = ?", phone, password).Find(&user).Error; err != nil {
		log.Println(err.Error())
		return false
	}
	//觉得这里不需要再验证密码和电话了
	return true
}

//获取用户信息
func GetUserInfo(uid int) (User, error) {
	var user User
	if err := DB.Table("users").Where("id=?", uid).Find(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

//用电话获取用户信息
func GetUserInfoByPhone(phone string) (User, error) {
	var user User
	if err := DB.Table("users").Where("phone=?", phone).Find(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

//修改密码
func ModifyPassword(id int, newPassword string) error {
	err := DB.Table("users").Where("id=?", id).Updates(map[string]interface{}{"password": newPassword}).Error
	if err != nil {
		return err
	}
	return nil
}

//修改用户头像
func UpdateAvator(avatar User) error {
	if err := DB.Table("users").Where("id = ?", avatar.Id).Updates(map[string]interface{}{"avatar": avatar.Avatar, "sha": avatar.Sha, "path": avatar.Path}).Error; err != nil {
		return err
	}
	return nil
}

//修改用户信息 -fyc
func ChangeUserInfor(userinfor User) error {
	if err := DB.Table("users").Where("id=?", userinfor.Id).Updates(map[string]interface{}{"nickname": userinfor.Nickname, "phone": userinfor.Phone, "email": userinfor.Email, "idcard": userinfor.Idcard, "realname": userinfor.Realname, "gender": userinfor.Gender}).Error; err != nil {
		return err
	}

	return nil
}

//根据剧本Id查询剧本信息 -fyc
func GetScriptInfor(ScriptId string) (Script, error) {
	var scriptinfor Script
	if err := DB.Table("scripts").Where("id=?", ScriptId).Find(&scriptinfor).Error; err != nil {
		return Script{}, err
	}
	return scriptinfor, nil
}

//根据地点Id查询地点信息 -fyc
func GetPlaceInforbyId(PlaceId string) (Place, error) {
	var placeinfor Place
	if err := DB.Table("places").Where("id=?", PlaceId).Find(&placeinfor).Error; err != nil {
		return Place{}, err
	}
	return placeinfor, nil
}

//更改剧本封面 --fyc
func UpdateScriptCover(script Script) error {
	if err := DB.Table("scripts").Where("id=?", script.Id).Updates(Script{Avatar: script.Avatar, Sha: script.Sha, Path: script.Path}).Error; err != nil {
		return err
	}
	return nil
}

//更改地点图片1 --fyc
func UpdatePlacePictureone(place Place) error {
	if err := DB.Table("places").Where("id=?", place.Id).Updates(Place{Picture1: place.Picture1, Sha1: place.Sha1, Path1: place.Path1}).Error; err != nil {
		return err
	}
	return nil
}

//更改地点图片2 --fyc
func UpdatePlacePicturetwo(place Place) error {
	if err := DB.Table("places").Where("id=?", place.Id).Updates(Place{Picture2: place.Picture2, Sha2: place.Sha2, Path2: place.Path2}).Error; err != nil {
		return err
	}
	return nil
}

//更改地点图片3 --fyc
func UpdatePlacePicturethree(place Place) error {
	if err := DB.Table("places").Where("id=?", place.Id).Updates(Place{Picture3: place.Picture3, Sha3: place.Sha3, Path3: place.Path3}).Error; err != nil {
		return err
	}
	return nil
}

//返回剧本的封面、id、名字、简介 --fyc
func GetScriptCoverandNameandBreifIntro() ([]Scriptinterface, error) {
	var scripts []Scriptinterface
	if err := DB.Table("scripts").Find(&scripts).Error; err != nil {
		return []Scriptinterface{}, err
	}
	return scripts, nil
}

//根据地点名称返回地点信息 --fyc
func GetPlaceInfor(placename string) (Place, error) {
	var place Place
	if err := DB.Table("places").Where("name=?", placename).Find(&place).Error; err != nil {
		return Place{}, err
	}
	return place, nil
}

//根据标签搜索剧本 --fyc
func GetInforbyTag(Tag string) ([]Scriptinterface, error) {
	var scripts []Scriptinterface
	if err := DB.Table("scripts").Where("tag1 LIKE(?) or tag2 LIKE(?) or tag3 LIKE(?) or tag4 LIKE(?) or tag5 LIKE(?) ", "%"+Tag+"%", "%"+Tag+"%", "%"+Tag+"%", "%"+Tag+"%", "%"+Tag+"%").Find(&scripts).Error; err != nil {
		return []Scriptinterface{}, err
	}
	return scripts, nil
}

//修改剧本信息 --fyc
func ChangeScriptInfor(script Script) error {
	if err := DB.Table("scripts").Where("id=?", script.Id).Updates(map[string]interface{}{"script_name": script.ScriptName, "introduction": script.Introduction, "brief_intro": script.BriefIntro, "price": script.Price,"tag1": script.Tag1, "tag2": script.Tag2, "tag3": script.Tag3, "tag4": script.Tag4, "tag5": script.Tag5, "time": script.Time, "place": script.Place}).Error; err != nil {
		return err
	}
	return nil
}

//返回收藏的剧本 --fyc
func GetCollectScript(id int) ([]Scriptinterface, error) {
	var collections []ScriptCollections
	if err := DB.Table("script_collections").Where("users_id=?", id).Find(&collections).Error; err != nil {
		return []Scriptinterface{}, err
	}

	var Result []Scriptinterface

	for _, v := range collections {
		var result = Scriptinterface{}
		scriptid := strconv.Itoa(v.ScriptsId)
		scriptInfo, err := GetScriptInfor(scriptid)
		if err != nil {
			return []Scriptinterface{}, err
		}
		result.Id = scriptInfo.Id
		result.Avatar = scriptInfo.Avatar
		result.BriefIntro = scriptInfo.BriefIntro
		result.ScriptName = scriptInfo.ScriptName

		Result = append(Result, result)
	}

	return Result, nil
}

//取消剧本收藏 --fyc
func CancelScript(user_id int, script_id int) error {
	if err := DB.Table("script_collections").Where("users_id=? and scripts_id=?",user_id,script_id).Delete(ScriptCollections{}).Error; err != nil {
		return err
	}
	return nil
}

//取消剧本预约 --fyc
func cancelappoint(user_id int, script_id int) error{
	if err := DB.Table("script_appointments").Where("users_id=? and scripts_id=?",user_id,script_id).Delete(ScriptAppointments{}).Error; err != nil {
		return err
	}
	return nil
}

//返回所有剧本订单 --fyc
func GetScriptOrder(user_id int) ([]ScriptOrders,error) {
	var orders []ScriptOrders
	if err := DB.Table("script_orders").Where("user_id=?",user_id).Order("id Desc").Find(&orders).Error; err != nil{
		return []ScriptOrders{}, err
	}
	return orders,nil
}

//返回所有剧本预约 --fyc
func GetScriptAppointment(user_id int)([]ScriptAppointments,error) {
	var appointments []ScriptAppointments
	if err := DB.Table("script_appointments").Where("users_id=?",user_id).Order("time Desc").Find(&appointments).Error; err != nil{
		return []ScriptAppointments{}, err
	}
	return appointments,nil
}

//修改商店图片
func UpdateShopAvator(avatar Shop) error {
	if err := DB.Table("shops").Where("id = ?", avatar.Id).Updates(map[string]interface{}{"picture": avatar.Picture, "sha": avatar.Sha, "path": avatar.Path}).Error; err != nil {
		return err
	}
	return nil
}

//修改精彩放送图片
func UpdateBroAvator(avatar Broadcast) error {
	if err := DB.Table("broadcasts").Where("id = ?", avatar.Id).Updates(map[string]interface{}{"picture": avatar.Picture, "sha": avatar.Sha, "path": avatar.Path}).Error; err != nil {
		return err
	}
	return nil
}

//获取精彩放送的信息
func GetBroadcastInfo() ([]Broadcast, error) {
	var Bro []Broadcast
	if err := DB.Table("broadcasts").Find(&Bro).Error; err != nil {
		return []Broadcast{}, err
	}
	return Bro, nil
}

//获取所有商店的信息
func GetShopInfo() ([]Shop, error) {
	var shop []Shop
	if err := DB.Table("shops").Find(&shop).Error; err != nil {
		return []Shop{}, err
	}
	return shop, nil
}

//获取单个精彩放送的信息
func GetSingleBroadcastInfo(id string) (Broadcast, error) {
	var Bro Broadcast
	if err := DB.Table("broadcasts").Where("id = ?", id).Find(&Bro).Error; err != nil {
		return Broadcast{}, err
	}
	return Bro, nil
}

//获取单个商店的信息
func GetSingleShopInfo(id string) (Shop, error) {
	var shop Shop
	if err := DB.Table("shops").Where("id = ?", id).Find(&shop).Error; err != nil {
		return Shop{}, err
	}
	return shop, nil
}

//删除单个商店的信息
func DeleteSingleShopInfo(id string) error {
	var shop Shop
	if err := DB.Table("shops").Where("id = ?", id).Delete(&shop).Error; err != nil {
		return err
	}
	return nil
}

//删除单个精彩放送的信息
func DeleteSingleBroInfo(id string) error {
	var broadcasts Broadcast
	if err := DB.Table("broadcasts").Where("id = ?", id).Delete(&broadcasts).Error; err != nil {
		return err
	}
	return nil
}

//根据状态搜索预约 --fyc
func GetAppointmentByStatus(status string,user_id int) ([]ScriptAppointments, error) {
	var appointments []ScriptAppointments
	if err := DB.Table("script_appointments").Where("status = ? and users_id = ? ",status,user_id).Find(&appointments).Error; err != nil {
		return []ScriptAppointments{}, err
	}
	return appointments, nil
}

//修改订单状态 --fyc
func UpdateOrderStatus(order_id int) error {
	if err := DB.Table("script_orders").Where("id=?",order_id).Updates(map[string]interface{}{"status":"已付款"}).Error; err != nil {
		return err
	}

	return nil
}

//修改预约状态 --fyc
func UpdateAppointmentStatus(appointment_id int,status string) error {
	if err := DB.Table("script_appointments").Where("id=?",appointment_id).Updates(map[string]interface{}{"status":status}).Error; err != nil {
		return err
	}
	return nil
}