package gormx

import (
	"gorm.io/gorm"
)

type Context struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewCtx(db *gorm.DB) Context {
	return Context{db: db}
}

func (item *Context) Begin() {
	item.tx = item.db.Begin()
}

func (item *Context) CheckErr(err error) {
	if err == nil {
		return
	}
	if item.tx != nil {
		item.tx.Rollback()
		item.tx = nil
	}
}

func (item *Context) RDB() *gorm.DB {
	return item.db
}

func (item *Context) WDB() *gorm.DB {
 	if item.tx != nil {
 		return item.tx
	}
	return item.db
}

func (item *Context) Commit() {
 	if item.tx != nil {
 		item.tx.Commit()
	}
}
