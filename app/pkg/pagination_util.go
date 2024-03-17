package pkg

import (
	"errors"

	"gorm.io/gorm"
)

type Paginator struct {
	Total  int `json:"total" form:"total"`
	Limit  int `json:"limit" form:"limit"`
	Offset int `json:"offset" form:"offset"`
	Page   int `json:"page" form:"page"`
}

func PaginatorHandler(paginator *Paginator) error {
	if paginator == nil {
		return errors.New("params error")
	}
	if paginator.Limit == 0 {
		paginator.Limit = 10
	}
	if paginator.Page == 0 {
		paginator.Offset = 0
	} else if paginator.Page > 0 {
		paginator.Offset = (paginator.Page - 1) * paginator.Limit
	}
	return nil
}

func (p *Paginator) GormPagination() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(p.Offset).Limit(p.Limit)
	}
}
