package api

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"kkako-blog/internal/dao"
	"kkako-blog/internal/model"
	"kkako-blog/pkg/bean"
	"kkako-blog/pkg/gormx"
	"kkako-blog/pkg/oauth"
	"strconv"
	"time"
)

type OauthApi struct {
	userDao     *dao.UserDao
	userAuthDao *dao.UserAuthDao
}

func NewOauthApi() *OauthApi {
	return &OauthApi{
		userDao:     dao.NewUserDao(),
		userAuthDao: dao.NewUserAuthDao(),
	}
}

func (item *OauthApi) GithubConfig(ctx *context) error {
	return ctx.ToResponse(map[string]interface{}{
		"client_id":    oauth.GithubConf.ClientId,
		"redirect_url": oauth.GithubConf.RedirectUrl,
	})
}

func (item *OauthApi) GithubAuth(ctx *context) error {
	code := ctx.PostForm("code")
	authType, err := ctx.PostInt("auth_type")
	if err != nil {
		return err
	}
	info, err := oauth.GetUserByAuth(authType, code)
	if err != nil {
		return err
	}
	userAuth, err := item.userAuthDao.First(
		gormx.WithEQForce("identifier", info.ID),
		gormx.WithEQForce("identity_type", model.IdentityTypeGithub))
	// 查询错误
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	// 未找到对应的auth，第一次登陆
	user := &model.User{}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user.Username = info.Username
		user.AvatarUrl = info.AvatarUrl
		user.CreateAt = time.Now()
		user.State = model.STATUS_NORMAL
		err := item.userDao.Add(user)
		if err != nil {
			return err
		}
		newAuth := &model.UserAuth{}
		newAuth.Identifier = strconv.Itoa(info.ID)
		newAuth.IdentityType = model.IdentityTypeGithub
		newAuth.UserId = user.ID
		err = item.userAuthDao.Add(newAuth)
		if err != nil {
			return err
		}
	} else {
		user, err = item.userDao.First(gormx.WithEQForce("id", userAuth.UserId))
		if err != nil {
			return err
		}
	}
	userRes := &model.UserRes{}
	userToken, err := item.userDao.SaveCache(user)
	if err != nil {
		return err
	}
	err = bean.BeanCopy(userRes, user)
	if err != nil {
		return err
	}
	userRes.Token = userToken
	return ctx.ToResponse(userRes)
}
