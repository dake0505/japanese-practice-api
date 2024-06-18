package service

import (
	"gin-gonic-api/app/constant"
	"gin-gonic-api/app/pkg"
	"gin-gonic-api/app/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ItemService interface {
	GetItemList(c *gin.Context)
}

type ItemServiceImpl struct {
	itemRepository repository.ItemRepository
}

func (i ItemServiceImpl) GetItemList(c *gin.Context) {
	data := i.itemRepository.GetItemList()
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func ItemServiceInit(itemRepository repository.ItemRepository) *ItemServiceImpl {
	return &ItemServiceImpl{
		itemRepository: itemRepository,
	}
}
