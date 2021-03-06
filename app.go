package main

import (
	"github.com/funnythingz/go-plog-api/db"
	"github.com/funnythingz/go-plog-api/handler"
	"net/http"

	"goji.io"
	"goji.io/pat"
)

var (
	exceptionHandler = &handler.ExceptionHandler{}
	plogsHandler     = &handler.PlogsHandler{}
	commentsHandler  = &handler.CommentsHandler{}
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
	mux.HandleFuncC(pat.Get("/api/v1/plogs/:id"), plogsHandler.Plog)

	// Comments
	mux.HandleFuncC(pat.Get("/api/v1/comments"), commentsHandler.Comments)
	mux.HandleFuncC(pat.Post("/api/v1/comments"), commentsHandler.CreateComment)
	mux.HandleFuncC(pat.Get("/api/v1/comments/:id"), commentsHandler.Comment)

	// Colors
	mux.HandleFuncC(pat.Get("/api/v1/colors"), colorsHandler.Colors)
	mux.HandleFuncC(pat.Post("/api/v1/colors"), colorsHandler.CreateColor)
	mux.HandleFuncC(pat.Get("/api/v1/colors/:id"), colorsHandler.Color)

	// Serve
	http.Handle("/", mux)
}
