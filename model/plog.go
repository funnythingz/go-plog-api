package model

import (
	"github.com/funnythingz/go-plog-api/db"
	_ "github.com/k0kubun/pp"
)

type Plog struct {
	Entity
	Color   Color  `json:"color"`
	ColorId int    `json:"color_id"`
	Content string `json:"content"`
}

func (m *Plog) Commit() {
	db.Dbmap.NewRecord(m)
	db.Dbmap.Create(&m).Related(&m.Color)
}

func (m *Plog) Update() {
	db.Dbmap.Model(&Plog{}).Update(m)
}

func (m *Plog) Fetch(id int) {
	db.Dbmap.Find(&m, id).First(&m).Related(&m.Color)
}

type PlogList struct {
	PlogList []Plog `json:"plog_list"`
}

func (m *PlogList) Fetch(permit int, page int) {
	db.Dbmap.Order("id desc").Offset((page - 1) * permit).Limit(permit).Find(&m.PlogList).Offset(page * permit).Limit(permit)
	for i, iro := range m.PlogList {
		m.PlogList[i].Fetch(iro.Id)
	}
}
