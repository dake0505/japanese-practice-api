package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type Paginator struct {
	Total  int `json:"total" form:"total"`
	Limit  int `json:"limit" form:"limit"`
	Offset int `json:"offset" form:"offset"`
	Page   int `json:"page" form:"page"`
	Size   int `json:"size" form:"size"`
}

func PaginatorHandler(c *gin.Context) *Paginator {
	var p = &Paginator{}
	p.Page, p.Size = cast.ToInt(c.Query("page")), cast.ToInt(c.Query("size"))
	if p.Page == 0 || p.Size == 0 {
		p.Page, p.Size = 1, 10
	}
	p.Offset = (p.Page - 1) * p.Size
	return p
}

func (p *Paginator) GormPagination() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(p.Offset).Limit(p.Size)
	}
}
