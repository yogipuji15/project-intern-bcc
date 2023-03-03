package entity

import (
	"math"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message   string `json:"message"`
	Code      int    `json:"code"`
	IsSuccess bool   `json:"isSuccess"`
}

type Pagination struct {
	Limit          int64 `form:"limit" json:"limit" gorm:"-"`
	Page           int64 `form:"page" json:"-" gorm:"-"`
	Offset         int64 `json:"-" gorm:"-"`
	CurrentPage    int64 `json:"currentPage" gorm:"-"`
	TotalPage      int64 `json:"totalPage" gorm:"-"`
	CurrentElement int64 `json:"currentElement" gorm:"-"`
	TotalElement   int64 `json:"totalElement" gorm:"-"`
}

func (pg *Pagination) ProcessPagination(rowsAffected int64) {
	pg.CurrentPage = pg.Page
	pg.TotalPage = int64(math.Ceil(float64(pg.TotalElement) / float64(pg.Limit)))
	pg.CurrentElement = rowsAffected
}

func FormatPaginationParam(params Pagination) Pagination {
	paginationParam := params

	if params.Limit == 0 {
		paginationParam.Limit = 10
	}

	if params.Page == 0 {
		paginationParam.Page = 1
	}

	paginationParam.Offset = (params.Page - 1) * paginationParam.Limit

	return paginationParam
}

type FilterParam struct{
	Keyword		string	`form:"keyword" json:"-" gorm:"-"`
	Category	string	`form:"category" json:"-" gorm:"-"`
	Location	string  `form:"location" json:"-" gorm:"-"`
}