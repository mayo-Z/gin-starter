package dao

import (
	"fmt"
	"gorm.io/gorm"
)

type Level struct {
	gorm.Model
	Id      uint64 `  description:"等级ID"`
	Name    string `gorm:"unique" description:"等级名"`
	OrderBy int    ` description:"排序值"`
	Status  bool   ` description:"是否已启用状态"`
}

// TableName 表名
func (receiver *Level) TableName() string {
	return "level"
}

// FindAll 分页查询
func (receiver *Level) FindAll(db *gorm.DB, page int, pageSize int) (res []*Level, count int64, err error) {
	var result *gorm.DB
	result = db.Table(receiver.TableName()).Where("deleted_at is null ").Count(&count).Offset((page - 1) * pageSize).Limit(pageSize).Find(&res)
	return res, count, result.Error
}

// Check 检查数据是否已存在
func (receiver *Level) Check(db *gorm.DB, name string) bool {
	row := db.Where(" name = ? ", name).First(&receiver)
	fmt.Println("err:", row.Error)
	if row.Error == nil {
		return false
	}
	return true
}

// EditCheck 检查数据是否已存在
func (receiver *Level) EditCheck(db *gorm.DB, id uint64, name string) bool {
	row := db.Where(" name = ? AND id != ?", name, id).First(&receiver)
	fmt.Println("err:", row.Error)
	if row.Error == nil {
		return false
	}
	return true
}

// Create 创建
func (receiver *Level) Create(db *gorm.DB, newContent *Level) {

	db.Create(&newContent)
}

// Edit 编辑
func (receiver *Level) Edit(db *gorm.DB, newContent map[string]interface{}) error {
	result := db.Model(receiver).Where(" id= ?", newContent["Id"]).Updates(&newContent)
	return result.Error
}

// SoftDelete 批量软删除，传id数组
func (receiver *Level) SoftDelete(db *gorm.DB, ids []uint64) error {
	for _, id := range ids {
		result := db.Where(" id = ?", id).Delete(&receiver)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}
