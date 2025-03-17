package router

import (
	"gin-ranking/api/config"
	"gin-ranking/api/controllers"
	"gin-ranking/api/pkg/logger"

	"github.com/gin-contrib/sessions"
	session_redis "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)

	// 配置Redis会话存储
	store, _ := session_redis.NewStore(10, "tcp", config.RedisAddress, "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// 添加静态文件服务
	r.Static("/css", "../web/css")
	r.Static("/js", "../web/js")
	r.Static("/images", "../web/images")

	// 添加默认路由，提供index.html
	r.GET("/", func(c *gin.Context) {
		c.File("../web/index.html")
	})

	// 添加登录页面路由
	r.GET("/login.html", func(c *gin.Context) {
		c.File("../web/login.html")
	})

	// 添加注册页面路由
	r.GET("/register.html", func(c *gin.Context) {
		c.File("../web/register.html")
	})

	// 创建API路由组
	api := r.Group("/api")
	{
		user := api.Group("/user")
		{
			user.POST("/register", controllers.UserController{}.Register)
			user.POST("/login", controllers.UserController{}.Login)
		}

		player := api.Group("/player")
		{
			player.POST("/list", controllers.PlayerController{}.GetPlayers)
		}

		vote := api.Group("/vote")
		{
			vote.POST("/add", controllers.VoteController{}.AddVote)
		}

		api.POST("/ranking", controllers.PlayerController{}.GetRanking)
	}

	return r
}
