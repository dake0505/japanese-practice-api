package service

import (
	"gin-gonic-api/app/constant"
	dao "gin-gonic-api/app/domain/dao/question_item"
	"gin-gonic-api/app/domain/dao/record"
	"gin-gonic-api/app/domain/dto"
	"gin-gonic-api/app/pkg"
	"gin-gonic-api/app/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ItemService interface {
	GetItemList(c *gin.Context)
	CreateQuestionItem(item dto.CreateItemRequest) dao.QuestionItem
	UpdateQuestionItem(item dto.UpdateItemRequest) dao.QuestionItem
	QueryQuestionDetail(questionId string, createdBy string) dto.QuestionDetailDto
}

type ItemServiceImpl struct {
	itemRepository   repository.ItemRepository
	answerRepository repository.AnswerRepository
	recordRepository repository.RecordRepository
}

func (i ItemServiceImpl) GetItemList(c *gin.Context) {
	data := i.itemRepository.GetItemList()
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (i ItemServiceImpl) QueryQuestionDetail(questionId string, createdBy string) dto.QuestionDetailDto {
	questionInfo := i.itemRepository.QueryQuestionDetail(questionId)
	answerList := i.answerRepository.QueryAnswerListByQuestionId(questionInfo.QuestionID)
	nextQuestionId := i.itemRepository.QueryNextQuestionId(questionId)
	preQuestionId := i.itemRepository.QueryPreviousQuestionId(questionId)
	input := record.Record{
		QuestionId: questionId,
		RecordType: "favorite",
	}
	records, err := i.recordRepository.QueryRecordList(&input, createdBy)
	if err != nil {

	}
	isFavorite := len(records) > 0
	answerDtos := make([]dto.AnswerItem, len(answerList))
	for i, answer := range answerList {
		answerDtos[i] = dto.AnswerItem{
			AnswerId:   answer.AnswerId,
			AnswerDesc: answer.AnswerDesc,
		}
	}
	questionDetail := dto.QuestionDetailDto{
		ID:             questionInfo.ID,
		QuestionID:     questionInfo.QuestionID,
		QuestionTitle:  questionInfo.QuestionTitle,
		AnswerItems:    answerDtos,
		NextQuestionId: nextQuestionId,
		PreQuestionId:  preQuestionId,
		IsFavorite:     isFavorite,
	}
	return questionDetail
}

func (i ItemServiceImpl) CreateQuestionItem(item dto.CreateItemRequest) dao.QuestionItem {
	// 处理数据并调用数据层
	input := dao.QuestionItem{
		QuestionTitle: item.QuestionTitle,
		AnswerId:      item.AnswerId,
	}
	data := i.itemRepository.CreateQuestionItem(&input)
	return data
}

func (i ItemServiceImpl) UpdateQuestionItem(item dto.UpdateItemRequest) dao.QuestionItem {
	input := dao.QuestionItem{
		QuestionID:    item.QuestionID,
		ID:            item.ID,
		AnswerId:      item.AnswerId,
		QuestionTitle: item.QuestionTitle,
	}
	data := i.itemRepository.UpdateQuestionItem(&input)
	return data
}

func ItemServiceInit(itemRepository repository.ItemRepository, answerRepository repository.AnswerRepository, recordRepository repository.RecordRepository) *ItemServiceImpl {
	return &ItemServiceImpl{
		itemRepository:   itemRepository,
		answerRepository: answerRepository,
		recordRepository: recordRepository,
	}
}
