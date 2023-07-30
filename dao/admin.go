package dao

import (
	"errors"
	"gin-starter/dto"
	"gin-starter/public"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Id       uint   `json:"uid"  description:"唯一id"`
	Username string `gorm:"unique" description:"昵称"`
	Password string `description:"密码"`
}

func (admin *Admin) Find(uid uint) (*Admin, error) {
	result := db.Where("id = ?", uid).First(admin)
	return admin, result.Error
}

func (admin *Admin) Update(newContent map[string]interface{}) error {
	result := db.Model(admin).Where("id= ?", newContent["Id"]).Updates(&newContent)
	return result.Error
}

func (admin *Admin) Delete(uid uint) error {
	result := db.Where("id = ?", uid).Delete(&admin)
	return result.Error
}

// Check 检查数据是否已存在
func (admin *Admin) Check(username string) bool {
	row := db.Where("username = ?", username).First(&admin)
	if row.Error == nil {
		return false
	}
	return true
}
func (admin *Admin) LoginCheck(param *dto.AdminLoginInput) (*Admin, error) {
	row := db.Where("username = ?", param.Username).First(&admin)
	if row.Error != nil {
		return nil, errors.New("用户名不存在")
	}
	if !public.ValidPassword(admin.Password, param.Password) {
		return nil, errors.New("密码错误，请重新输入")
	}
	return admin, nil
}

func RegisterCheck(param *dto.RegisterInput) error {
	admin := &Admin{}
	db.Where("username = ?", param.Username).First(admin)
	if admin.ID != 0 {
		return errors.New("用户已存在")
	}
	return nil
}
