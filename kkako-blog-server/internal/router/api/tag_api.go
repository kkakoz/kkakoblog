package api

import (
	"kkako-blog/internal/dao"
	"kkako-blog/internal/model"
	"kkako-blog/pkg/errcode"
	"kkako-blog/pkg/gormx"
	"time"
)

type TagApi struct {
	tagDao *dao.TagDao
}

func NewTagApi() *TagApi {
	return &TagApi{
		tagDao:     dao.NewTagDao(),
	}
}

func (item *TagApi) GetList(ctx *context) error {
	page := ctx.GetPage()
	count, err := item.tagDao.FindCount()
	if err != nil {
		return err
	}
	tags, err := item.tagDao.FindInfoRes(gormx.WithPage(page.Page, page.PageSize))
	if err != nil {
		return err
	}
	return ctx.ToResponseList(tags, count)
}

func (item *TagApi) Add(ctx *context) error {
	name := ctx.PostForm("name")
	count, err := item.tagDao.FindCount(gormx.WithEQForce("name", name))
	if err != nil {
		return err
	}
	if count > 0 {
		return errcode.NewErr("已经有该名称的标签")
	}
	tag := &model.Tag{
		Name:       name,
		CreateTime: time.Now(),
	}

	err = item.tagDao.Add(tag)
	if err != nil {
		return err
	}
	return ctx.ToOK()
}

func (item *TagApi) Delete(ctx *context) error {
	id, err := ctx.PostInt("id")
	if err != nil {
		return err
	}
	err = item.tagDao.Delete(id)
	if err != nil {
		return err
	}
	return ctx.ToOK()
}

func (item *TagApi) Edit(ctx *context) error {
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
	err = item.tagDao.Update(update, gormx.WithEQForce("id", id))
	if err != nil {
		return err
	}
	return ctx.ToOK()
}