package api

import (
	"errors"
	"gorm.io/gorm"
	"kkako-blog/internal/dao"
	"kkako-blog/internal/model"
	"kkako-blog/pkg/bean"
	"kkako-blog/pkg/cryption"
	"kkako-blog/pkg/errcode"
	"kkako-blog/pkg/gormx"
)

type UserApi struct {
	userDao     *dao.UserDao
	userAuthDao *dao.UserAuthDao
}

func NewUserApi() UserApi {
	return UserApi{
		userDao:     dao.NewUserDao(),
		userAuthDao: dao.NewUserAuthDao(),
	}
}

// 用户登录
func (item *UserApi) Login(ctx *context) error {
	auth := &struct {
		Email    string `form:"email" binding:"required,min=4"`
		Password string `form:"pwd" binding:"required,min=6"`
	}{}
	//login := &model.LoginBO{}
	if err := ctx.Validate(auth); err != nil {
		return err
	}
	userAuth, err := item.userAuthDao.First(
		gormx.WithEQForce("identifier", auth.Email),
		gormx.WithEQForce("identity_type", model.IdentityTypeEmail))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errcode.NewErr("该账号不存在")
	}
	if err != nil {
		return err
	}
	if userAuth.Credential != cryption.Md5Str(auth.Password) {
		return errcode.NewErr("密码错误")
	}
	userRes := &model.UserRes{}
	user, err := item.userDao.First(gormx.WithEQ("id", userAuth.UserId))
	if err != nil {
		return err
	}
	token, err := item.userDao.SaveCache(user)
	if err != nil {
		return err
	}
	err = bean.BeanCopy(userRes, user)
	if err != nil {
		return err
	}
	userRes.Token = token
	return ctx.ToResponse(userRes)
}

//func (item *CategoryApi) Register(ctx *context) error {
//	auth := &struct {
//		Email  string `form:"email" binding:"required,min=4"`
//		Password string `form:"pwd" binding:"required,min=6"`
//	}{}
//	if err := ctx.Validate(auth); err != nil {
//		return err
//	}
//	_, err := item.userAuthDao.First(
//		gormx.WithEQForce("identifier", auth.Email),
//		gormx.WithEQForce("identity_type", model.IdentityTypeEmail))
//	if !errors.Is(err, gorm.ErrRecordNotFound) {
//		return errcode.NewErr("该账号已经存在")
//	}
//	if err != nil {
//		return err
//	}
//	user := &model.User{}
//	user.Account = login.Account
//	user.Password = cryption.Md5Str(login.Password)
//	user.CreateAt = times.GetNow()
//	err = item.userDao.Add(user)
//	if err != nil {
//		return err
//	}
//	userRes := &model.UserRes{}
//	err = bean.BeanCopy(userRes, user)
//	if err != nil {
//		return err
//	}
//	return ctx.ToResponse(userRes)
//}

func (item *UserApi) Info(ctx *context) error {
	token := ctx.Query("token")
	user, err := item.userDao.UserCache(token)
	if err != nil {
		return err
	}
	userRes := &model.UserRes{}
	err = bean.BeanCopy(userRes, user)
	if err != nil {
		return err
	}
	return ctx.ToResponse(userRes)
}

func (item *UserApi) AdminLogin(ctx *context) error {
	auth := &struct {
		Email    string `form:"email" binding:"required,min=4"`
		Password string `form:"password" binding:"required,min=6"`
	}{}
	if err := ctx.Validate(auth); err != nil {
		return err
	}
	userAuth, err := item.userAuthDao.First(
		gormx.WithEQForce("identifier", auth.Email),
		gormx.WithEQForce("identity_type", model.IdentityTypeEmail))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errcode.NewErr("该账号不存在")
	}
	if err != nil {
		return err
	}
	if userAuth.Credential != cryption.Md5Str(auth.Password) {
		return errcode.NewErr("密码错误")
	}
	userRes := &model.UserRes{}
	user, err := item.userDao.First(gormx.WithEQ("id", userAuth.UserId))
	if err != nil {
		return err
	}
	if !model.UserAuthManageVali(user.Auth) {
		return errcode.NewErr("没有后台权限")
	}
	token, err := item.userDao.SaveCache(user)
	if err != nil {
		return err
	}
	err = bean.BeanCopy(userRes, user)
	if err != nil {
		return err
	}
	userRes.Token = token
	return ctx.ToResponse(userRes)
}

func (item *UserApi) Logout(ctx *context) error {
	token := ctx.GetHeader("Authorization")
	err := item.userDao.DeleteCache(token)
	if err != nil {
		return err
	}
	return ctx.ToOK()
}
