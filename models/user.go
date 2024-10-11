package models

import (
	g "sail-chat/global"
	"sail-chat/types"
	"sail-chat/utils"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"
)

type User struct {
	Model
	Username 	 string       `json:"username" gorm:"index:idx_name,unique; not null; comment:用户名称"`
	Password 	 string       `json:"password" gorm:"not null; comment:用户密码"`
	Nickname 	 string       `json:"nickname" gorm:"not null; default:''; comment:昵称"`
	Avator 	 	 string       `json:"avator" gorm:"not null; default:''; comment:头像"`
	// Role 			 []Role       `json:"-" gorm:"many2many:user_roles;"`
	// Blacklist  []Blacklist  `json:"-" gorm:"foreignKey:UserId"`
	// Message 	 []Message 		`json:"-" gorm:"foreignKey:UserId"`
	// ReplyGroup []ReplyGroup `json:"-" gorm:"foreignKey:UserId"`
	// ReplyItem	 []ReplyItem 	`json:"-" gorm:"foreignKey:UserId"`
	// Visitor 	 []Visitor	 	`json:"-" gorm:"foreignKey:UserId"`
	// Welcome 	 []Welcome 		`json:"-" gorm:"foreignKey:UserId"`
}
type NotPass struct {
	Model
	Username string `json:"username" gorm:"not null; comment:用户名称"`
	Nickname string `json:"nickname" gorm:"not null; default:''; comment:昵称"`
	Avator 	 string `json:"avator" gorm:"not null; default:''; comment:头像"`
	RoleName string `json:"role_name" sql:"-"`
	RoleId 	 uint 	`json:"role_id" sql:"-"`

}

// CreateUser 创建用户
func CreateUser(d *gorm.DB, u *types.User) (uint, error) {
	user := User{
		Username: u.Username,
		Password: u.Password,
		Avator: u.Avator,
		Nickname: u.Nickname,
	}
	user.CreatedAt = LocalTime{time.Now()}
	// g.App.DB.Create(&user)
	err := d.Create(&user).Error
	return user.Id, err
}

// GetUsers 查询所有用户
func GetUsers(meta *types.Meta) ([]NotPass, int64, error) {
	var db = g.App.DB.Model(&User{})
	var count int64
	var userList []NotPass
	var err error
	db.Where("username like ?", "%"+meta.Value+"%").Count(&count)
	if count < meta.Size {
		err = db.Select("users.*,roles.name role_name,roles.id role_id").Joins("join user_roles on users.id=user_roles.user_id").Joins("join roles on user_roles.role_id=roles.id").Having("username like ?", "%"+meta.Value+"%").Find(&userList).Error
	} else {
		err = db.Select("users.*,roles.name role_name,roles.id role_id").Joins("join user_roles on users.id=user_roles.user_id").Joins("join roles on user_roles.role_id=roles.id").Scopes(utils.Paginate(meta)).Find(&userList).Error
	}
	return userList, count, err
}

// GetUserByID 查询用户
func GetUserByID(id uint) (NotPass, error) {
	var user NotPass
	err := g.App.DB.Table("users").Select("users.*,roles.name role_name,roles.id role_id").Joins("join user_roles on users.id=user_roles.user_id").Joins("join roles on user_roles.role_id=roles.id").Where("users.id = ?",id).First(&user).Error
	return user, err
}

// GetUserByName 查询用户
func GetUserByName(name string) (User, error) {
	var user User
	err := g.App.DB.Where("username = ?", name).First(&user).Error
	return user, err
}

// GetUserByNameAndPassword 查询用户
func GetUserByNameAndPassword(name string, password string) User {
	var user User
	g.App.DB.Model(&user).Where("username = ? and password = ?", name, password).First(&user)
	return user
}

// DeleteUser 删除用户
func DeleteUser(user User) error {
	err := g.App.DB.Model(&user).Updates(map[string]interface{}{
		//"deleted_at": time.Now(),
		"is_deleted": true,
	}).Error
	return err
}

// UpdateUser 更新用户
func UpdateUser(d *gorm.DB, u *NotPass) error {
	omit := []string{"created_at", "deleted_at", "id"}
	u.UpdatedAt = LocalTime{time.Now()}
	err := d.Model(&User{}).Omit(omit...).Where("id = ?", u.Id).Updates(&u).Error
	return err
}

// UpdatePass 修改密码
func UpdatePass(user User) (string, error) {
	user.UpdatedAt = LocalTime{time.Now()}
	omit := []string{"update_at", "deleted_at", "id"}
	err := g.App.DB.Omit(omit...).Save(&user).Error
	return user.Username, err
}

func GetUserByCustom(name string) (User, error) {
	wrapper := map[string]interface{}{
		"username": "%" + name + "%",
	}
	p := &utils.Regular[User]{}
	u, err := p.SelectOne(wrapper)
	if err != nil {
		return User{}, err
	} else {
		return u, err
	}
}
