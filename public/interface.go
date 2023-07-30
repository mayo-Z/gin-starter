package public

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Input interface {
	BindValidParam(ctx *gin.Context) error
}

// CRUD 使用了泛型
type CRUD[T any] interface {
	TableName() string
	FindAll(db *gorm.DB, page int, pageSize int) (res []*T, count int64, err error)
	Check(db *gorm.DB, name string) bool
	EditCheck(db *gorm.DB, id uint64, name string) bool
	Create(db *gorm.DB, newContent *T)
	Edit(db *gorm.DB, newContent map[string]interface{}) error
	SoftDelete(db *gorm.DB, ids []uint64) error
}
type APICRUD interface {
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Edit(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
