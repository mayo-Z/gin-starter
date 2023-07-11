package dto

import (
	"gin-starter/public"
	"github.com/gin-gonic/gin"
)

type GetLevelsInput struct {
	Page uint `json:"page" form:"page"  comment:"页码" validate:"required"`
	Size uint `json:"size" form:"size"  comment:"每页数量" validate:"required"`
}

type CreateLevelInput struct {
	Name    string `form:"name" comment:"等级名"  example:"name" validate:"required"`
	OrderBy int    `form:"orderBy" comment:"排序值"  example:"1" `
	Status  bool   `form:"status" comment:"是否已启用状态"  example:"true" `
}
type EditLevelInput struct {
	Id uint64 `form:"id"  comment:"等级ID" example:"1" validate:"required"`
	CreateLevelInput
}

type LevelsIdInput struct {
	Ids []uint64 `json:"ids" form:"ids"   `
}

//--------------------返回接口-----------------------------------------

type LevelOutput struct {
	Id         uint64 `  comment:"等级ID"`
	Name       string `comment:"等级名"`
	OrderBy    int    ` comment:"排序值"`
	Status     bool   ` comment:"是否已启用状态"`
	CreateTime string ` comment:"创建时间"`
	UpdateTime string ` comment:"更新时间" `
}

type LevelPaginationOutput struct {
	LevelOutput []LevelOutput
	Page        uint  `json:"page" form:"page"  comment:"页码" `
	Size        uint  `json:"size" form:"size"  comment:"数量" `
	Count       int64 `json:"count" form:"count"  comment:"总数" `
}

func (a *GetLevelsInput) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, a)
}

func (a *CreateLevelInput) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, a)
}
func (a *EditLevelInput) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, a)
}

func (a *LevelsIdInput) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, a)
}
