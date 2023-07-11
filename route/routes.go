package route

import (
	"gin-starter/controller"
	"gin-starter/dao"
	"gin-starter/docs"
	"gin-starter/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
)

// InitConfig 配置文件初始化R
func InitConfig() {
	work, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(work + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("err")
	}

	docs.SwaggerInfo.Title = viper.GetString("swagger.title")
	docs.SwaggerInfo.Description = viper.GetString("swagger.desc")
	docs.SwaggerInfo.Host = viper.GetString("swagger.host")
	docs.SwaggerInfo.BasePath = viper.GetString("swagger.base_path")
}

// InitRouter 路由初始化
func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares...)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//获取session
	store, err := dao.GetSessionStore()
	router.Use(sessions.Sessions("mySession", store))
	if err != nil {
		log.Fatalf("sessions.NewRedisStoreerr:%v", err)
	}

	//会员接口-----------------------------------------------------------------
	memberRouter := router.Group("/member")
	memberRouter.Use(
		middleware.TranslationMiddleware())
	{
		controller.MemberRegister(memberRouter)
	}
	//非登陆接口-----------------------------------------------------------------
	adminLoginRouter := router.Group("/auth")
	// mySession是返回給前端的sessionId名
	adminLoginRouter.Use(
		middleware.TranslationMiddleware())
	{
		controller.AdminLoginRegister(adminLoginRouter)
	}
	//用户接口-----------------------------------------------------------------
	adminRouter := router.Group("/admin")
	adminRouter.Use(
		middleware.SessionAuthMiddleware(),
		middleware.TranslationMiddleware())
	{
		controller.AdminRegister(adminRouter)
	}

	return router
}
