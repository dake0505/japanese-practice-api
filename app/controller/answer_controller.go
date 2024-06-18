package controller

import "gin-gonic-api/app/service"

type AnswerController interface{}

type AnswerControllerImpl struct {
	answerService service.AnswerService
}

func AnswerControllerInit(answerService service.AnswerService) *AnswerControllerImpl {
	return &AnswerControllerImpl{
		answerService: answerService,
	}
}
