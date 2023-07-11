package dto

import (
	"gin-starter/public"
	"github.com/gin-gonic/gin"
)

type AdminSessionInfo struct {
	Uid      uint   `json:"uid" `
	Username string `json:"username" `
}

type RegisterInput struct {
	Username string `json:"username" comment:"用户名" example:"admin" validate:"required"`
	Password string ` comment:"密码" example:"123456" validate:"required,valid_password"`
}

func (a *RegisterInput) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, a)
}

type AdminLoginInput struct {
	Username string `  comment:"用户名" example:"admin" validate:"required"`
	Password string ` comment:"密码" example:"123456" validate:"required,valid_password"`
}

func (a *AdminLoginInput) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, a)
}

//type AdminLoginOutput struct {
//	Token string `json:"token" form:"token" comment:"token" example:"token" validate:""`
//}
