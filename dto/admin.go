package dto

import (
	"gin-starter/public"
	"github.com/gin-gonic/gin"
)

type AdminInfoOutput struct {
	Uid  uint   `form:"uid"`
	Name string `form:"name"`
}

type ChangePwdInput struct {
	Password string `json:"password" form:"password" comment:"密码" example:"123456" validate:"required,valid_password"`
}

func (a *ChangePwdInput) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, a)
}
