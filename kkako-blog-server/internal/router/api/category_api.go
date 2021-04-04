package api

import (
	"kkako-blog/internal/dao"
	"kkako-blog/internal/model"
	"kkako-blog/pkg/errcode"
	"kkako-blog/pkg/gormx"
	"time"
)

type CategoryApi struct {
	categoryDao     *dao.CategoryDao
}

func NewCategoryApi() CategoryApi {
	return CategoryApi{
		categoryDao:     dao.NewCategoryDao(),
	}
}

func (item *CategoryApi) GetList(ctx *context) error {
	page := ctx.GetPage()
	count, err := item.categoryDao.FindCount()
	if err != nil {
		return err
	}
	categories, err := item.categoryDao.FindInfoRes(gormx.WithPage(page.Page, page.PageSize))
	if err != nil {
		return err
	}
	return ctx.ToResponseList(categories, count)
}

func (item *CategoryApi) Add(ctx *context) error {
	name := ctx.PostForm("name")
	count, err := item.categoryDao.FindCount(gormx.WithEQForce("name", name))
	if err != nil {
		return err
	}
	if count > 0 {
		return errcode.NewErr("已经有该名称的分类")
	}
	category := &model.Category{
		Name:       name,
		CreateTime: time.Now(),
	}

	err = item.categoryDao.Add(category)
	if err != nil {
		return err
	}
	return ctx.ToOK()
}

func (item *CategoryApi) Enable(ctx *context) error {
	id, err := ctx.PostInt("id")
	if err != nil {
		return err
	}
	err = item.categoryDao.UpdateEnable(id)
	if err != nil {
		return err
	}
	return ctx.ToOK()
}

func (item *CategoryApi) Delete(ctx *context) error {
	id, err := ctx.PostInt("id")
	if err != nil {
		return err
	}
	err = item.categoryDao.Delete(id)
	if err != nil {
		return err
	}
	return ctx.ToOK()
}

func (item *CategoryApi) Edit(ctx *context) error {
	id, err := ctx.PostInt("id")
	if err != nil {
		return err
	}
	name := ctx.PostForm("name")
	if name == "" {
		return errcode.NewErr("名称不能为空")
	}
	update := map[string]interface{}{
		"name": name,
	}
	err = item.categoryDao.Update(update, gormx.WithEQForce("id", id))
	if err != nil {
		return err
	}
	return ctx.ToOK()
}