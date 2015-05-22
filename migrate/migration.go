package main

import (
	"../db"
	"../mapper"
	"fmt"
	"log"
	"os"
)

func main() {
	db.Connect()

	action := "migrate"
	if len(os.Args) >= 2 {
		action = os.Args[1]
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
	log.Println(db.Dbmap.DropTableIfExists(&mapper.Iro{}))
	log.Println(db.Dbmap.DropTableIfExists(&mapper.Color{}))
	Create()
}

func Create() {
	log.Println(db.Dbmap.CreateTable(&mapper.Iro{}))
	log.Println(db.Dbmap.CreateTable(&mapper.Color{}))
}

func Migrate() {
	log.Println(db.Dbmap.AutoMigrate(&mapper.Iro{}))
	log.Println(db.Dbmap.AutoMigrate(&mapper.Color{}))
}
