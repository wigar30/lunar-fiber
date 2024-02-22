package utils

import (
	"math"

	"gorm.io/gorm"
)

type Pagination struct {
	Items       interface{} `json:"items"`
	Limit       int         `json:"limit,omitempty"`
	Page        int         `json:"page,omitempty"`
	Sort        string      `json:"sort,omitempty"`
	TotalData   int64       `json:"totalData"`
	TotalPage   int         `json:"totalPage"`
	AnyNextPage bool        `json:"anyNextPage"`
	AnyPrevPage bool        `json:"anyPrevPage"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetAnyNextPage() bool {
	return p.GetPage() < p.TotalPage
}

func (p *Pagination) GetAnyPrevPage() bool {
	return p.GetPage() > 1
}

func Paginate(p *Pagination, db *gorm.DB, model interface{}) func(db *gorm.DB) *gorm.DB {
	var totalData int64
	db.Model(model).Count(&totalData)

	p.TotalData = totalData

	totalPage := int(math.Ceil(float64(totalData) / float64(p.Limit)))
	p.TotalPage = totalPage

	p.AnyNextPage = p.GetAnyNextPage()
	p.AnyPrevPage = p.GetAnyPrevPage()
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(p.GetOffset()).Limit(p.GetLimit())
	}
}
