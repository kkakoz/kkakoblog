package model

import "time"

type Article struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	UserId     int       `json:"user_id"`
	CategoryId int       `json:"category_id"`
	Star       int       `json:"star"`
	View       int       `json:"view"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	Enable     bool      `json:"enable"`
}

type ArticleRes struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	UserId       int    `json:"user_id"`
	Username     string `json:"username"`
	CategoryId   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
	Star         int    `json:"star"`
	View         int    `json:"view"`
	CreateTime   string `json:"create_time"`
	UpdateTime   string `json:"update_time"`
	Enable     bool      `json:"enable"`
	Tags         []Tag  `json:"tags"`
}

type ArticleTag struct {
	ArticleId int `json:"article_id"`
	TagId     int `json:"tag_id"`
}

type ArticleBO struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	UserId     int    `json:"user_id"`
	CategoryId int    `json:"category_id"`
}
