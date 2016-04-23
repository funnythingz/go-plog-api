package main

import (
	"github.com/funnythingz/go-plog-api/db"
	"github.com/funnythingz/go-plog-api/handler"
	"net/http"
	"regexp"

	"goji.io"
	"goji.io/pat"
)

var (
	exceptionHandler = &handler.ExceptionHandler{}
	plogsHandler     = &handler.PlogsHandler{}
	colorsHandler    = &handler.ColorsHandler{}
)

func init() {
	// Database
	db.Connect()

	// Goji
	mux := goji.NewMux()

	// Plogs
	mux.HandleFuncC(pat.Get("/api/v1/plogs"), plogsHandler.Plogs)
	mux.HandleFuncC(pat.Post("/api/v1/plogs"), plogsHandler.CreatePlog)
	mux.HandleFuncC(pat.Get(regexp.MustCompile(`^/api/v1/plogs/(?P<id>\d+)$`)), plogsHandler.Plog)

	// Colors
	mux.HandleFuncC(pat.Get("/v1/colors"), colorsHandler.Colors)
	mux.HandleFuncC(pat.Post("/v1/colors"), colorsHandler.CreateColor)
	mux.HandleFuncC(pat.Get(regexp.MustCompile(`^/api/v1/colors/(?P<id>\d+)$`)), colorsHandler.Color)

	// Serve
	http.Handle("/", mux)
}
