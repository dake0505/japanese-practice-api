package service

import (
	dao "gin-gonic-api/app/domain/dao/answer_item"
	"gin-gonic-api/app/domain/dto"
	"gin-gonic-api/app/repository"
)

type AnswerService interface {
	CreateAnswerItem(item dto.CreateAnswerItemDto) dao.AnswerItem
	QueryAnswerList() []dao.AnswerItem
}

type AnswerServiceImpl struct {
	answerRepository repository.AnswerRepository
}

func (a AnswerServiceImpl) CreateAnswerItem(item dto.CreateAnswerItemDto) dao.AnswerItem {
	input := dao.AnswerItem{
		QuestionID: item.QuestionID,
		AnswerDesc: item.AnswerDesc,
	}
	data := a.answerRepository.CreateAnswerItem(&input)
	return data
}

func (a AnswerServiceImpl) QueryAnswerList() []dao.AnswerItem {
	items := a.answerRepository.QueryAnswerList()
	return items
}

func AnswerServiceInit(answerRepository repository.AnswerRepository) *AnswerServiceImpl {
	return &AnswerServiceImpl{
		answerRepository: answerRepository,
	}
}
