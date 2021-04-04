package dao

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"kkako-blog/internal/model"
	"kkako-blog/pkg/gormx"
	"kkako-blog/pkg/mysql"
	"sync"
)

type CategoryDao struct {
	db *gorm.DB
}

var categoryDao *CategoryDao
var categoryOnce sync.Once

func NewCategoryDao() *CategoryDao {
	categoryOnce.Do(func() {
		categoryDao = &CategoryDao{db: mysql.DB()}
	})
	return categoryDao
}

func (item *CategoryDao) Add(category *model.Category) error {
	err := item.db.Create(category).Error
	return errors.Wrap(err, "添加失败")
}

func (item *CategoryDao) Find(opts ...gormx.Option) ([]model.Category, error) {
	db := item.db.Table("categories")
	gormx.DBOpt(db, opts)
	categorys := make([]model.Category, 0)
	err := db.Find(&categorys).Error
	return categorys, errors.Wrap(err, "查找分类列表失败")
}

func (item *CategoryDao) FindCount(opts ...gormx.Option) (int64, error) {
	db := item.db.Table("categories")
	gormx.DBOpt(db, opts)
	var count int64
	err := db.Count(&count).Error
	return count, errors.Wrap(err, "查找分类数量失败")
}


func (item *CategoryDao) First(opts ...gormx.Option) (*model.Category, error) {
	db := item.db.Table("categories")
	gormx.DBOpt(db, opts)
	category := &model.Category{}
	err := db.First(category).Error
	return category, errors.Wrap(err, "查找分类失败")
}

func (item *CategoryDao) Update(update map[string]interface{}, opts ...gormx.Option) error {
	db := item.db.Table("categories")
	gormx.DBOpt(db, opts)
	err := db.Updates(update).Error
	return errors.Wrap(err, "更新列表失败")
}

func (item *CategoryDao) Delete(id int) error {
	category := &model.Category{
		ID: id,
	}
	err := item.db.Table("categories").Delete(category).Error
	return errors.Wrap(err, "删除失败")
}

func (item *CategoryDao) FindInfoRes(opts ...gormx.Option) ([]model.CategoryInfoRes, error) {
	db := item.db.Table("categories")
	gormx.DBOpt(db, opts)
	categorys := make([]model.CategoryInfoRes, 0)
	err := db.Select("categories.id, categories.name, categories.enable, " +
		"DATE_FORMAT(categories.create_time,'%Y-%m-%d %H:%i:%s') as create_time").
		Find(&categorys).Error
	return categorys, errors.Wrap(err, "查找分类列表失败")
}

func (item *CategoryDao) UpdateEnable(id int) error {
	db := item.db.Table("categories")
	err := db.Where("id = ?", id).Update("enable", gorm.Expr(" !enable")).Error
	return errors.Wrap(err, "更新是否启用失败")
}
