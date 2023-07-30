package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"gin-starter/dao"
	"gin-starter/dto"
	"gin-starter/middleware"
	"gin-starter/public"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterController struct{}

func AdminLoginRegister(group *gin.RouterGroup) {
	adminLogin := &RegisterController{}
	group.POST("/register", adminLogin.Register)
	group.POST("/login", adminLogin.Login)
	group.GET("/quit", adminLogin.AdminQuit)
	group.GET("/logout", adminLogin.Logout)
}

// Register godoc
// @Summary 用户注册
// @Description 用户注册
// @Tags 用户接口
// @ID /auth/register
// @Param param body dto.RegisterInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /auth/register [post]
func (*RegisterController) Register(ctx *gin.Context) {
	params := &dto.RegisterInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, 2001, err)
		return
	}
	db := dao.GetDB()
	if err := dao.RegisterCheck(db, params); err != nil {
		middleware.ResponseError(ctx, 2002, err)
		return
	}

	//	创建用户
	hashedPassword, err := public.SetHashedPassword(params.Password)
	if err != nil {
		middleware.ResponseError(ctx, 2004, errors.New("加密错误"))
		return
	}
	newUser := dao.Admin{
		Username: params.Username,
		Password: hashedPassword,
	}
	if !newUser.Check(db, params.Username) {
		middleware.ResponseError(ctx, 2003, errors.New("用户名已存在"))
		return
	}
	db.Create(&newUser)
	//	返回结果
	middleware.ResponseSuccess(ctx, "注册成功")
}

// Login godoc
// @Summary 用户登录
// @Description 用户登录
// @Tags 用户接口
// @ID /auth/login
// @Param param body dto.AdminLoginInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /auth/login [post]
func (*RegisterController) Login(ctx *gin.Context) {
	//	获取参数
	params := &dto.AdminLoginInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, 2001, err)
		return
	}
	//数据库验证
	db := dao.GetDB()
	admin := &dao.Admin{}
	admin, err := admin.LoginCheck(db, params)
	if err != nil {
		middleware.ResponseError(ctx, 2002, err)
		return
	}

	//session
	sessInfo := &dto.AdminSessionInfo{
		Uid:      admin.Id,
		Username: admin.Username,
	}

	sessBts, err := json.Marshal(sessInfo)
	if err != nil {
		middleware.ResponseError(ctx, 2003, err)
		return
	}
	sess := sessions.Default(ctx)
	sess.Options(sessions.Options{
		Path:     "/",
		MaxAge:   648000,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	})
	sess.Set(public.AdminSessionInfoKey, string(sessBts))

	err = sess.Save()
	if err != nil {
		middleware.ResponseError(ctx, 2004, err)
		return
	}
	middleware.ResponseSuccess(ctx, "登陆成功")
}

// AdminQuit godoc
// @Summary 用户退出
// @Description 用户退出
// @Tags 用户接口
// @ID /auth/quit
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /auth/quit [get]
func (*RegisterController) AdminQuit(ctx *gin.Context) {
	sess := sessions.Default(ctx)
	sess.Delete(public.AdminSessionInfoKey)
	sess.Save()
	middleware.ResponseSuccess(ctx, "退出成功")
}

// Logout godoc
// @Summary 用户注销
// @Description 用户注销
// @Tags 用户接口
// @ID /auth/logout
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /auth/logout [get]
func (*RegisterController) Logout(ctx *gin.Context) {
	sess := sessions.Default(ctx)
	sessInfo := sess.Get(public.AdminSessionInfoKey)
	adminSessionInfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(fmt.Sprint(sessInfo)), adminSessionInfo); err != nil {
		middleware.ResponseError(ctx, 2001, err)
		return
	}
	db := dao.GetDB()
	admin := &dao.Admin{}
	err := admin.Delete(db, adminSessionInfo.Uid)
	if err != nil {
		middleware.ResponseError(ctx, 2002, err)
		return
	}
	middleware.ResponseSuccess(ctx, "注销成功")
}
