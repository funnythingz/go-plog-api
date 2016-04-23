package model

import (
	"github.com/funnythingz/go-plog-api/db"
	_ "github.com/k0kubun/pp"
)

type Color struct {
	Entity
	Name     string `json:"name"`
	Code     string `json:"color_code"`
	TextCode string `json:"text_code"`
}

func (m *Color) Commit() {
	db.Dbmap.NewRecord(m)
	db.Dbmap.Create(&m)
}

func (m *Color) Fetch(id int) {
	db.Dbmap.Find(&m, id).First(&m)
}

type ColorList struct {
	ColorList []Color `json:"color_list"`
}

func (m *ColorList) Fetch(permit int, page int) {
	db.Dbmap.Order("id asc").Offset((page - 1) * permit).Limit(permit).Find(&m.ColorList).Offset(page * permit).Limit(permit)
}
