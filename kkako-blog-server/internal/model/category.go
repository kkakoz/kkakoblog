package model

import "time"

type Category struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	CreateTime time.Time `json:"create_time"`
	Enable     bool      `json:"enable"`
}

type CategoryInfoRes struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	CreateTime string `json:"create_time"`
	Enable     bool      `json:"enable"`
	Count      int    `json:"count"`
}
