package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"kkako-blog/internal/model"
	"kkako-blog/log"
	"kkako-blog/pkg/errcode"
	"strconv"
)

func Wrap(f func(ctx *context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := NewCtx(c)
		err := f(ctx)
		if err != nil {
			ctx.DealError(err)
		}
	}
}

const (
	BackSuccess = 200
	BackFail = 500
)

type context struct {
	*gin.Context
}

func (item *context) GetPage() *model.Pager {
	pager := &model.Pager{}
	err := item.ShouldBind(pager)
	if err != nil {
		pager.Page = 1
		pager.PageSize = 10
	}
	return pager
}

func (item *context) PostInt(name string) (int, error) {
	idstr := item.PostForm(name)
	return strconv.Atoi(idstr)
}

func (item *context) PostInt64(name string) (int64, error) {
	idstr := item.PostForm(name)
	return strconv.ParseInt(idstr, 10, 64)
}

func NewCtx(ctx *gin.Context) *context {
	return &context{Context: ctx}
}

func (item *context) ToResponse(data interface{}) error {
	item.JSON(200, model.Res{
		Code:  BackSuccess,
		Data:  data,
	})
	return nil
}

func (item *context) ToOK() error {
	item.JSON(200, model.Res{
		Code:  BackSuccess,
	})
	return nil
}

// 返回数据列表
func (item *context) ToResponseList(list interface{}, totalRow int64) error {
	item.JSON(200, model.Res{
		Code:  BackSuccess,
		Data:  list,
		Msg:   "",
		Count: totalRow,
	})
	return nil
}

// 绑定和校验
func (item *context) Validate(data interface{}) error {
	err := item.ShouldBind(data)
	if err != nil {
		return errcode.NewErr("参数错误, err=" + err.Error())
	}
	return nil
}

func (item *context) DealError(err error) {
	var kErr *errcode.KError
	b := errors.As(err, &kErr)
	if b {
		item.JSON(200, kErr)
	} else {
		log.L().Error(fmt.Sprintf("err is %+v",err),
			zap.String("url", item.Request.URL.String()))
		item.JSON(200, model.Res{
			Code: BackFail,
			Msg:  err.Error(),
		})
	}
}

func (item *context) ParamInt(key string) (int, error) {
	value := item.Param(key)
	valueInt, err := strconv.Atoi(value)
	if err != nil {
		return 0, errcode.NewErr("参数错误")
	}
	return valueInt, err
}

func (item *context) GetUser() (*model.User, error){
	value, exists := item.Get("user")
	if !exists {
		return nil, errors.New("用户登录信息已经失效")
	}
	user, ok := value.(*model.User)
	if !ok {
		return nil, errors.New("用户登录信息已经失效")
	}
	return user, nil
}