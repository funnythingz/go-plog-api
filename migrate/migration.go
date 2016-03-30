package main

import (
	"../db"
	"../model"
	"fmt"
	"log"
	"os"
)

func main() {
	db.Connect()

	action := "migrate"
	if len(os.Args) >= 2 {
		action = os.Args[2]
	}

	log.Println(fmt.Sprintf("mode: %s", action))

	switch {
	case action == "migrate":
		Migrate()
		return

	case action == "reset":
		Reset()
		return
	}
}

func Reset() {
	db.Dbmap.DropTableIfExists(&model.Plog{})
	db.Dbmap.DropTableIfExists(&model.Color{})
	Create()
}

func Create() {
	db.Dbmap.CreateTable(&model.Plog{})
	db.Dbmap.CreateTable(&model.Color{})
}

func Migrate() {
	db.Dbmap.AutoMigrate(&model.Plog{})
	db.Dbmap.AutoMigrate(&model.Color{})
}
