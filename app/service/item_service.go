package service

import (
	"gin-gonic-api/app/constant"
	dao "gin-gonic-api/app/domain/dao/question_item"
	"gin-gonic-api/app/domain/dto"
	"gin-gonic-api/app/pkg"
	"gin-gonic-api/app/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ItemService interface {
	GetItemList(c *gin.Context)
	CreateQuestionItem(item dto.CreateItemRequest) dao.QuestionItem
}

type ItemServiceImpl struct {
	itemRepository repository.ItemRepository
}

func (i ItemServiceImpl) GetItemList(c *gin.Context) {
	data := i.itemRepository.GetItemList()
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
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

func ItemServiceInit(itemRepository repository.ItemRepository) *ItemServiceImpl {
	return &ItemServiceImpl{
		itemRepository: itemRepository,
	}
}
