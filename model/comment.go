package model

import (
	"github.com/funnythingz/go-plog-api/db"
	_ "github.com/k0kubun/pp"
)

type Comment struct {
	Entity
	Content string `json:"content"`
	PlogId  int    `json:"plog_id"`
}

func (m *Comment) Commit() {
	db.Dbmap.NewRecord(m)
	db.Dbmap.Create(&m)
}

func (m *Comment) Update() {
	db.Dbmap.Model(&Comment{}).Update(m)
}

func (m *Comment) Fetch(id int) {
	db.Dbmap.Find(&m, id).First(&m)
}

type CommentList struct {
	CommentList []Comment `json:"comment_list"`
}

func (m *CommentList) Fetch(permit int, page int) {
	db.Dbmap.Order("id desc").Offset((page - 1) * permit).Limit(permit).Find(&m.CommentList).Offset(page * permit).Limit(permit)
	for i, comment := range m.CommentList {
		m.CommentList[i].Fetch(comment.Id)
	}
}
