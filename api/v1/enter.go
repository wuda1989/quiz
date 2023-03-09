package v1

import "quiz/api/v1/quiz"

type ApiGroup struct {
	CommentApi quiz.CommentApi
}

var ApiGroupApp = new(ApiGroup)
