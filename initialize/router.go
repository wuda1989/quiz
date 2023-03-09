package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quiz/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	quizRouter := router.InitGroupApp.QuizRouter

	PublicGroup := Router.Group("quiz")
	{
		// 心跳
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})

		quizRouter.InitQuizRouter(PublicGroup)
	}

	return Router
}
