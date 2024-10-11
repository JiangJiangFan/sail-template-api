package models

import (
	"sail-chat/global"
)

type ReplyGroup struct {
	Model
	Name  string `json:"g_name" gorm:"not null; comment: 组"`
	UserId uint   `json:"u_id" gorm:"index:user_id,unique; not null"`
}


// CreateReplyGroup 创建
func CreateReplyGroup(ReplyGroup *ReplyGroup) (uint, error) {
	// 创建国家/城市
	err := global.App.DB.Create(&ReplyGroup)
	return ReplyGroup.Id, err.Error
}

// UpdateReplyGroup 更新
func UpdateReplyGroup(ReplyGroup *ReplyGroup) (uint, error) {
	err := global.App.DB.Save(&ReplyGroup).Error
	return ReplyGroup.Id, err
}

// SelectReplyGroupCustom 自定义查询
// func SelectReplyGroupCustom(m *types.Meta) ([]ReplyGroupList, error) {
// 	var ReplyGroupList []ReplyGroupList
// 	// offset, limit := utils.PaginateString(m)
// 	// const str = `select name , (select name from countries where id = (c.p_id))
// 	// 	 as p_name from countries c where name like ?`
// 	var db = global.App.DB
// 	// err := db.Raw(str, "%"+m.Value+"%").Scopes(utils.Paginate(m)).Scan(&ReplyGroupList).Error
// 	err := db.Table("countries c").Model(&ReplyGroup{}).Select("id, name , (select name from countries where id = (c.p_id)) as p_name, p_id, level").Where("name like ?", "%"+m.Value+"%").Scopes(utils.Paginate(m)).Find(&ReplyGroupList).Error
// 	return ReplyGroupList, err
// }

// SelectReplyGroupCustomByValue  自定义查询
// func SelectReplyGroupCustomByValue(m *types.Meta) ([]ReplyGroupList, error) {
// 	var ReplyGroupList []ReplyGroupList
// 	var db = global.App.DB
// 	err := db.Table("countries c").Model(&ReplyGroup{}).Select("id, name , (select name from countries where id = (c.p_id)) as p_name, level").Where("name like ?", "%"+m.Value+"%").Find(&ReplyGroupList).Error
// 	return ReplyGroupList, err
// }
