package dao

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"kkako-blog/internal/model"
	"kkako-blog/pkg/gormx"
	"kkako-blog/pkg/mysql"
	"sync"
)

type TagDao struct {
	db *gorm.DB
}

var tagDao *TagDao
var tagOnce sync.Once

func NewTagDao() *TagDao {
	tagOnce.Do(func() {
		tagDao = &TagDao{db: mysql.DB()}
	})
	return tagDao
}

func (item *TagDao) Add(tag *model.Tag) error {
	err := item.db.Create(tag).Error
	return errors.Wrap(err, "添加失败")
}

func (item *TagDao) Find(opts ...gormx.Option) ([]model.Tag, error) {
	db := item.db.Table("tags")
	gormx.DBOpt(db, opts)
	tags := make([]model.Tag, 0)
	err := db.Find(&tags).Error
	return tags, errors.Wrap(err, "查找列表失败")
}

func (item *TagDao) FindCount(opts ...gormx.Option) (int64, error) {
	db := item.db.Table("tags")
	gormx.DBOpt(db, opts)
	var count int64
	err := db.Count(&count).Error
	return count, errors.Wrap(err, "查找数量失败")
}

func (item *TagDao) First(opts ...gormx.Option) (*model.Tag, error) {
	db := item.db.Table("tags")
	gormx.DBOpt(db, opts)
	tag := &model.Tag{}
	err := db.First(tag).Error
	return tag, errors.Wrap(err, "查找失败")
}

func (item *TagDao) Update(update map[string]interface{}, opts ...gormx.Option) error {
	db := item.db.Table("tags")
	gormx.DBOpt(db, opts)
	err := db.Updates(update).Error
	return errors.Wrap(err, "更新失败")
}

func (item *TagDao) Delete(id int) error {
	tag := &model.Tag{
		ID: id,
	}
	err := item.db.Table("tags").Delete(tag).Error
	return errors.Wrap(err, "删除失败")
}

func (item *TagDao) FindInfoRes(opts ...gormx.Option) ([]model.TagInfoRes, error) {
	db := item.db.Table("tags")
	gormx.DBOpt(db, opts)
	tags := make([]model.TagInfoRes, 0)
	err := db.Select("tags.id, tags.name, " +
		"DATE_FORMAT(tags.create_time,'%Y-%m-%d %H:%i:%s') as create_time").
		Find(&tags).Error
	return tags, errors.Wrap(err, "查找标签列表失败")
}
