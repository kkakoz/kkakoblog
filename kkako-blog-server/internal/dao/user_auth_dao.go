package dao

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"kkako-blog/internal/model"
	"kkako-blog/pkg/gormx"
	"kkako-blog/pkg/mysql"
	"sync"
)

type UserAuthDao struct {
    db *gorm.DB
}

var userAuthDao *UserAuthDao
var userAuthOnce sync.Once

func NewUserAuthDao() *UserAuthDao {
    userAuthOnce.Do(func() {
        userAuthDao = &UserAuthDao{db: mysql.DB()}
    })
    return userAuthDao
}

func (item *UserAuthDao) Add(userAuth *model.UserAuth) error {
    err := item.db.Create(userAuth).Error
    return errors.Wrap(err, "添加失败")
}

func (item *UserAuthDao) Find(opts ...gormx.Option) ([]model.UserAuth, error) {
    db := item.db.Table("user_auths")
    gormx.DBOpt(db, opts)
    userAuths := make([]model.UserAuth, 0)
    err := db.Find(&userAuths).Error
    return userAuths, errors.Wrap(err, "查找列表失败")
}

func (item *UserAuthDao) FindCount(opts ...gormx.Option) (int64, error) {
    db := item.db.Table("user_auths")
    gormx.DBOpt(db, opts)
    var count int64
    err := db.Count(&count).Error
    return count, errors.Wrap(err, "查找数量失败")
}

func (item *UserAuthDao) First(opts ...gormx.Option) (*model.UserAuth, error) {
    db := item.db.Table("user_auths")
    gormx.DBOpt(db, opts)
    userAuth := &model.UserAuth{}
    err := db.First(userAuth).Error
    return userAuth, errors.Wrap(err, "查找失败")
}

func (item *UserAuthDao) Update(update map[string]interface{},opts ...gormx.Option) error {
    db := item.db.Table("user_auths")
    gormx.DBOpt(db, opts)
    err := db.Updates(update).Error
    return errors.Wrap(err, "更新失败")
}

func (item *UserAuthDao) Delete(id int) error {
    userAuth := &model.UserAuth{
      ID:id,
    }
    err := item.db.Table("user_auths").Delete(userAuth).Error
    return errors.Wrap(err, "删除失败")
}