package handler

import (
	"encoding/json"
	"fmt"
	"github.com/funnythingz/go-plog-api/helper"
	"github.com/funnythingz/go-plog-api/model"
	"github.com/funnythingz/go-plog-api/services"
	_ "github.com/k0kubun/pp"
	"goji.io/pat"
	"golang.org/x/net/context"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"unicode/utf8"
)

type PlogsHandler struct{}

func (h *PlogsHandler) Plogs(c context.Context, w http.ResponseWriter, r *http.Request) {

	if service.BeforeAuth(w, r) == false {
		return
	}

	permit := 50
	urlQuery, _ := url.ParseQuery(r.URL.RawQuery)

	var page int
	if len(urlQuery["page"]) == 0 {
		page = 1
	} else {
		page, _ = strconv.Atoi(urlQuery["page"][0])
	}

	plogs := model.PlogList{}
	plogs.Fetch(permit, page)
	response, _ := json.Marshal(plogs)
	io.WriteString(w, string(response))
}

func (h *PlogsHandler) Plog(c context.Context, w http.ResponseWriter, r *http.Request) {

	if service.BeforeAuth(w, r) == false {
		return
	}

	id, _ := strconv.Atoi(pat.Param(c, "id"))
	plog := model.Plog{}
	plog.Fetch(id)
	if plog.Id == 0 {
		helper.ResultMessageJSON(w, []string{fmt.Sprintf("Not Found: %d", id)})
		return
	}
	response, _ := json.Marshal(plog)
	io.WriteString(w, string(response))
}

func (h *PlogsHandler) CreatePlog(c context.Context, w http.ResponseWriter, r *http.Request) {

	if service.BeforeAuth(w, r) == false {
		return
	}

	content := r.FormValue("plog[content]")
	colorId, _ := strconv.Atoi(r.FormValue("color[color_id]"))

	// Validation
	errors := []string{}

	if utf8.RuneCountInString(content) <= 0 {
		errors = append(errors, "input Content must be blank.")
	}
	if utf8.RuneCountInString(content) < 5 || utf8.RuneCountInString(content) > 1000 {
		errors = append(errors, "input Content minimum is 5 and maximum is 1000 character.")
	}

	if len(errors) > 0 {
		helper.ResultMessageJSON(w, errors)
		return
	}

	plog := model.Plog{
		Color: model.Color{
			Entity: model.Entity{
				Id: colorId,
			},
		},
		Content: content,
	}

	plog.Commit()
	response, _ := json.Marshal(plog)
	io.WriteString(w, string(response))
}
