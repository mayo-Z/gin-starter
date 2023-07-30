package public

import (
	"github.com/gin-gonic/gin"
)

type Input interface {
	BindValidParam(ctx *gin.Context) error
}

// CRUD 使用了泛型
type CRUD[T any] interface {
	TableName() string
	FindAll(page int, pageSize int) (res []*T, count int64, err error)
	Check(name string) bool
	EditCheck(id uint64, name string) bool
	Create(newContent *T)
	Edit(newContent map[string]interface{}) error
	SoftDelete(ids []uint64) error
}
type APICRUD interface {
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Edit(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
