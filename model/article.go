package model

import "time"

func articleList() []Article {
	mods := make([]Article, 10)
	DB.Select(&mods, `select * from article`)
	return mods
}

type Article struct {
	id         int
	Name       string
	CreateTime time.Time
	Author     string
}
