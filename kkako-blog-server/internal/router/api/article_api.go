package api

import (
	"kkako-blog/internal/dao"
	"kkako-blog/internal/model"
	"kkako-blog/pkg/bean"
	"kkako-blog/pkg/gormx"
)

type ArticleApi struct {
	articleDao *dao.ArticleDao
}

func NewArticleApi() ArticleApi {
	return ArticleApi{
		articleDao: dao.NewArticleDao(),
	}
}

func (item *ArticleApi) ArticleAdd(ctx *context) error {
	articleBo := &model.ArticleBO{}
	err := ctx.Validate(articleBo)
	if err != nil {
		return err
	}
	user, err := ctx.GetUser()
	if err != nil {
		return err
	}
	article := &model.Article{}
	err = bean.BeanCopy(article, articleBo)
	if err != nil {
		return err
	}
	article.UserId = user.ID
	err = item.articleDao.Add(article)
	if err != nil {
		return err
	}
	return ctx.ToResponse(nil)
}

func (item *ArticleApi) ArticlePage(ctx *context) error {
	page := ctx.GetPage()
	categoryId, _ := ctx.PostInt("category_id")
	keyword := ctx.PostForm("keyword")
	count, err := item.articleDao.FindCount(gormx.WithLike("title", keyword),
		gormx.WithEQ("category_id", categoryId))
	if err != nil {
		return err
	}
	articles, err := item.articleDao.Find(gormx.WithLike("title", keyword),
		gormx.WithEQ("category_id", categoryId),
		gormx.WithPage(page.Page, page.PageSize))
	if err != nil {
		return err
	}
	return ctx.ToResponseList(articles, count)
}

func (item *ArticleApi) Info(ctx *context) error {
	id, err := ctx.ParamInt("id")
	if err != nil {
		return err
	}
	article, err := item.articleDao.FirstBack(gormx.WithEQForce("id", id))
	if err != nil {
		return err
	}
	return ctx.ToResponse(article)
}