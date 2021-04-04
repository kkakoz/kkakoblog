package dao

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"kkako-blog/internal/model"
	"kkako-blog/pkg/cache"
	"kkako-blog/pkg/cryption"
	"kkako-blog/pkg/gormx"
	"kkako-blog/pkg/mysql"
	"sync"
	"time"
)

type UserDao struct {
	db *gorm.DB
}

var userDao *UserDao
var userOnce sync.Once

func NewUserDao() *UserDao {
	userOnce.Do(func() {
		userDao = &UserDao{db: mysql.DB()}
	})
	return userDao
}

func (item *UserDao) Add(user *model.User) error {
	err := item.db.Create(user).Error
	return errors.Wrap(err, "添加失败")
}

func (item *UserDao) Find(opts ...gormx.Option) ([]model.User, error) {
	db := item.db.Table("users")
	gormx.DBOpt(db, opts)
	users := make([]model.User, 0)
	err := db.Find(&users).Error
	return users, errors.Wrap(err, "查找列表失败")
}

func (item *UserDao) FindCount(opts ...gormx.Option) (int64, error) {
	db := item.db.Table("users")
	gormx.DBOpt(db, opts)
	var count int64
	err := db.Count(&count).Error
	return count, errors.Wrap(err, "查找数量失败")
}

func (item *UserDao) First(opts ...gormx.Option) (*model.User, error) {
	db := item.db.Table("users")
	gormx.DBOpt(db, opts)
	user := &model.User{}
	err := db.First(user).Error
	return user, errors.Wrap(err, "查找失败")
}

func (item *UserDao) Update(update map[string]interface{}, opts ...gormx.Option) error {
	db := item.db.Table("users")
	gormx.DBOpt(db, opts)
	err := db.Updates(update).Error
	return errors.Wrap(err, "更新失败")
}

func (item *UserDao) Delete(id int) error {
	user := &model.User{
		ID: id,
	}
	err := item.db.Table("users").Delete(user).Error
	return errors.Wrap(err, "删除失败")
}

func (item *UserDao) UserCache(token string) (*model.User, error) {
	user := &model.User{}
	cmd := cache.Redis.Get(model.UserTokenKey + token)
	if cmd.Err() != nil {
		return user, cmd.Err()
	}
	data, err := cmd.Bytes()
	if err != nil {
		return user, err
	}
	fmt.Println(string(data))
	err = json.Unmarshal(data, user)
	return user, errors.Wrap(err, "获取用户缓存失败")
}

func (item *UserDao) SaveCache(user *model.User) (string, error) {
	token := cryption.UUID()
	userdata, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	cmd := cache.Redis.Set(model.UserTokenKey + token, userdata, 24 * time.Hour)
	if cmd.Err() != nil {
		return "", errors.Wrap(cmd.Err(), "存储用户信息失败")
	}
	return token, nil
}

func (item *UserDao) DeleteCache(token string) error {
	cmd := cache.Redis.Del(model.UserTokenKey + token)
	if cmd.Err() != nil {
		return errors.Wrap(cmd.Err(), "删除用户缓存失败")
	}
	return nil
}
