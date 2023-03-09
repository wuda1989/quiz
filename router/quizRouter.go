package router

import (
	"github.com/gin-gonic/gin"
	v1 "quiz/api/v1"
)

type QuizRouter struct {
}

func (s *QuizRouter) InitQuizRouter(Router *gin.RouterGroup) {
	quizRouter := Router.Group("v1")
	var commentApi = v1.ApiGroupApp.CommentApi
	{
		quizRouter.POST("comment", commentApi.CreateComment)
		quizRouter.GET("comment/:uuid", commentApi.FindComment)
		quizRouter.PUT("comment", commentApi.UpdateComment)
		quizRouter.DELETE("comment", commentApi.DeleteComment)
	}
}
