package model

import "time"

type Tag struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	CreateTime time.Time `json:"create_time"`
}

type TagInfoRes struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	CreateTime time.Time `json:"create_time"`
	Count      int       `json:"count"`
}
