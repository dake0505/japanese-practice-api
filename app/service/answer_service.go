package service

import "gin-gonic-api/app/repository"

type AnswerService interface{}

type AnswerServiceImpl struct {
	answerRepository repository.AnswerRepository
}

func AnswerServiceInit(answerRepository repository.AnswerRepository) *AnswerServiceImpl {
	return &AnswerServiceImpl{
		answerRepository: answerRepository,
	}
}
