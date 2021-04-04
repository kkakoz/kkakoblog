package gormx

import (
	"gorm.io/gorm"
)

func WithPage(page, pageSize int) Option {
	return func(db *gorm.DB) {
		if page <= 0 {
			return
		}
		if pageSize <= 0 {
			return
		}
		db = db.Limit(pageSize).Offset(pageSize * (page - 1))
	}
}

func WithOrder(order string) Option {
	return func(db *gorm.DB) {
		db = db.Order(order)
	}
}

func WithSelect(statement string) Option {
	return func(db *gorm.DB) {
		db = db.Select(statement)
	}
}
