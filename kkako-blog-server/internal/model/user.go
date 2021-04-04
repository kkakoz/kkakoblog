package model

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	AvatarUrl string    `json:"avatar_url"`
	State     int       `json:"state"`
	Auth      int       `json:"auth"`
	CreateAt  time.Time `json:"create_at"`
}

const (
	UserAuthManage = 1 >> iota
)

func UserAuthManageVali(auth int) bool {
	return UserAuthManage|auth != 0
}

type UserRes struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	AvatarUrl string    `json:"avatar"`
	CreateAt  time.Time `json:"create_at"`
	LoginTime time.Time `json:"login_time"`
	Token     string    `json:"token"`
}

const (
	STATUS_NORMAL = iota + 1
	STATUS_BAN
	STATUS_DELETE
)

const (
	UserTokenKey = "user:token:"
)
