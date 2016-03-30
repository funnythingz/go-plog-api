package main

import (
	"../db"
	"../handler"
	_ "github.com/zenazn/goji"
	"github.com/zenazn/goji/bind"
	"github.com/zenazn/goji/graceful"
	"github.com/zenazn/goji/web"
	"regexp"
)

var (
	exceptionHandler = &handler.ExceptionHandler{}
	plogsHandler     = &handler.PlogsHandler{}
	colorsHandler    = &handler.ColorsHandler{}
)

func main() {
	// Database
	db.Connect()

	// Goji
	m := web.New()

	// Plogs
	m.Get("/api/v1/plogs", plogsHandler.Plogs)
	m.Post("/api/v1/plogs", plogsHandler.CreatePlog)
	m.Get(regexp.MustCompile(`^/api/v1/plogs/(?P<id>\d+)$`), plogsHandler.Plog)

	// Colors
	m.Get("/v1/colors", colorsHandler.Colors)
	m.Post("/v1/colors", colorsHandler.CreateColor)
	m.Get(regexp.MustCompile(`^/api/v1/colors/(?P<id>\d+)$`), colorsHandler.Color)

	// Exception
	m.NotFound(exceptionHandler.NotFound)

	// Serve
	graceful.Serve(bind.Default(), m)
}
