package main

import (
	"github.com/funnythingz/go-plog-api/db"
	"github.com/funnythingz/go-plog-api/handler"
	"net/http"
	"regexp"

	"goji.io"
	"goji.io/pat"
	"golang.org/x/net/context"
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
	mux.HandlerFuncC(pat.Get("/api/v1/plogs"), plogsHandler.Plogs)
	mux.HandlerFuncC(pat.Post("/api/v1/plogs"), plogsHandler.CreatePlog)
	mux.HandlerFuncC(pat.Get(regexp.MustCompile(`^/api/v1/plogs/(?P<id>\d+)$`)), plogsHandler.Plog)

	// Colors
	mux.HandlerFuncC(pat.Get("/v1/colors"), colorsHandler.Colors)
	mux.HandlerFuncC(pat.Post("/v1/colors"), colorsHandler.CreateColor)
	mux.HandlerFuncC(pat.Get(regexp.MustCompile(`^/api/v1/colors/(?P<id>\d+)$`)), colorsHandler.Color)

	// Exception
	mux.NotFound(exceptionHandler.NotFound)

	// Serve
	http.Handle("/", mux)
}
