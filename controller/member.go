package controller

import (
	"errors"
	"gin-starter/dao"
	"gin-starter/dto"
	"gin-starter/middleware"
	"gin-starter/public"
	"github.com/gin-gonic/gin"
)

type MemberController struct{}

func MemberRegister(group *gin.RouterGroup) {
	content := &MemberController{}
	//----------------------------会员等级----------------------------------------------
	group.GET("/getLevels", content.Get)
	//---------------------------------------------------------------------------------
	group.POST("/createLevels", content.Create)
	group.PUT("/editLevels", content.Edit)
	group.DELETE("/delLevels", content.Delete)

}

//----------------------------会员等级----------------------------------------------

// Get godoc
// @Summary 分页获取所有等级
// @Tags 会员等级接口
// @ID /member/getLevels
// @Param input query dto.GetLevelsInput true "分页查询数据"
// @Success 200 {object} middleware.Response{data=dto.LevelPaginationOutput} "success"
// @Router /member/getLevels [get]
func (*MemberController) Get(ctx *gin.Context) {
	params := &dto.GetLevelsInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, 2001, err)
		return
	}
	db := dao.GetDB()
	content := &dao.Level{}
	detail, count, _ := content.FindAll(db, int(params.Page), int(params.Size))
	outData := []dto.LevelOutput{}
	for _, d := range detail {
		outData = append(outData, dto.LevelOutput{
			Id:         d.Id,
			Name:       d.Name,
			OrderBy:    d.OrderBy,
			Status:     d.Status,
			UpdateTime: public.TimeToStr(d.UpdatedAt),
			CreateTime: public.TimeToStr(d.CreatedAt),
		})
	}
	out := &dto.LevelPaginationOutput{
		LevelOutput: outData,
		Page:        params.Page,
		Size:        params.Size,
		Count:       count,
	}
	middleware.ResponseSuccess(ctx, out)
}

// Create godoc
// @Summary 创建等级
// @Tags 会员等级接口
// @ID /member/createLevels
// @Param param body dto.CreateLevelInput true "创建数据"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /member/createLevels [post]
func (*MemberController) Create(ctx *gin.Context) {
	params := &dto.CreateLevelInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, 2001, err)
		return
	}
	db := dao.GetDB()
	content := &dao.Level{}
	//检查会员等级名是否存在
	if !content.Check(db, params.Name) {
		middleware.ResponseError(ctx, 2002, errors.New("栏目已存在"))
		return
	}
	newContent := &dao.Level{
		Name:    params.Name,
		OrderBy: params.OrderBy,
		Status:  params.Status,
	}
	content.Create(db, newContent)
	middleware.ResponseSuccess(ctx, "创建成功")
}

// Edit godoc
// @Summary 修改等级
// @Tags 会员等级接口
// @ID /member/editLevels
// @Param param body dto.EditLevelInput true "数据"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /member/editLevels [put]
func (*MemberController) Edit(ctx *gin.Context) {
	params := &dto.EditLevelInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, 2001, err)
		return
	}
	db := dao.GetDB()
	content := &dao.Level{}
	if !content.EditCheck(db, params.Id, params.Name) {
		middleware.ResponseError(ctx, 2002, errors.New("栏目已存在"))
		return
	}
	newContent := map[string]interface{}{
		"Id":      params.Id,
		"Name":    params.Name,
		"OrderBy": params.OrderBy,
		"Status":  params.Status,
	}
	err := content.Edit(db, newContent)
	if err != nil {
		middleware.ResponseError(ctx, 2003, err)
		return
	}
	middleware.ResponseSuccess(ctx, "编辑成功")
}

// Delete godoc
// @Summary 批量删除
// @Tags 会员等级接口
// @ID /member/delLevels
// @Param input body dto.LevelsIdInput true "id数组"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /member/delLevels [delete]
func (*MemberController) Delete(ctx *gin.Context) {
	params := &dto.LevelsIdInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, 2001, err)
		return
	}
	db := dao.GetDB()
	content := &dao.Level{}
	if err := content.SoftDelete(db, params.Ids); err != nil {
		middleware.ResponseError(ctx, 2002, err)
		return
	}
	middleware.ResponseSuccess(ctx, "删除成功")
}
