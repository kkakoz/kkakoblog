package dao

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"kkako-blog/internal/model"
	"kkako-blog/pkg/gormx"
	"kkako-blog/pkg/mysql"
	"sync"
)

type ArticleDao struct {
	db *gorm.DB
}

var articleDao *ArticleDao
var articleOnce sync.Once

func NewArticleDao() *ArticleDao {
	articleOnce.Do(func() {
		articleDao = &ArticleDao{db: mysql.DB()}
	})
	return articleDao
}

func (item *ArticleDao) Add(article *model.Article) error {
	err := item.db.Create(article).Error
	return errors.Wrap(err, "添加失败")
}

func (item *ArticleDao) Find(opts ...gormx.Option) ([]model.Article, error) {
	db := item.db.Table("articles")
	gormx.DBOpt(db, opts)
	articles := make([]model.Article, 0)
	err := db.Find(&articles).Error
	return articles, errors.Wrap(err, "查找列表失败")
}

func (item *ArticleDao) FindCount(opts ...gormx.Option) (int64, error) {
	db := item.db.Table("articles")
	gormx.DBOpt(db, opts)
	var count int64
	err := db.Count(&count).Error
	return count, errors.Wrap(err, "查找数量失败")
}

func (item *ArticleDao) First(opts ...gormx.Option) (*model.Article, error) {
	db := item.db.Table("articles")
	gormx.DBOpt(db, opts)
	article := &model.Article{}
	err := db.First(article).Error
	return article, errors.Wrap(err, "查找失败")
}

func (item *ArticleDao) Update(update map[string]interface{}, opts ...gormx.Option) error {
	db := item.db.Table("articles")
	gormx.DBOpt(db, opts)
	err := db.Updates(update).Error
	return errors.Wrap(err, "更新失败")
}

func (item *ArticleDao) Delete(id int) error {
	article := &model.Article{
		ID: id,
	}
	err := item.db.Table("articles").Delete(article).Error
	return errors.Wrap(err, "删除失败")
}

func (item *ArticleDao) FirstBack(opts ...gormx.Option) (*model.ArticleRes, error) {
	db := item.db.Table("articles")
	db.Joins("left join join categories on articles.category_id = categories.id").
		Select("articles.*, categories.name")
	gormx.DBOpt(db, opts)
	res := &model.ArticleRes{}
	err := db.First(res).Error
	if err != nil {
		return res, errors.Wrap(err, "查询失败")
	}
	tags := make([]model.Tag, 0)
	err = item.db.Table("article_tags").
		Joins("left join tag on article_tags.tag_id = tag.id").
		Where("article_tags.article_id = ?", res.ID).
		Select("tag.*").Find(&tags).Error
	res.Tags = tags
	return res, errors.Wrap(err, "查询失败")
}
