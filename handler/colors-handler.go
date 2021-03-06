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
	"regexp"
	"strconv"
	"unicode/utf8"
)

type ColorsHandler struct{}

func (h *ColorsHandler) Colors(c context.Context, w http.ResponseWriter, r *http.Request) {

	if service.BeforeAuth(w, r) == false {
		return
	}

	permit := 10
	urlQuery, _ := url.ParseQuery(r.URL.RawQuery)

	var page int
	if len(urlQuery["page"]) == 0 {
		page = 1
	} else {
		page, _ = strconv.Atoi(urlQuery["page"][0])
	}

	colors := model.ColorList{}
	colors.Fetch(permit, page)
	response, _ := json.Marshal(colors)
	io.WriteString(w, string(response))
}

func (h *ColorsHandler) Color(c context.Context, w http.ResponseWriter, r *http.Request) {

	if service.BeforeAuth(w, r) == false {
		return
	}

	id, _ := strconv.Atoi(pat.Param(c, "id"))
	color := model.Color{}
	color.Fetch(id)
	if color.Id == 0 {
		helper.ResultMessageJSON(w, []string{fmt.Sprintf("Not Found: %d", id)})
		return
	}
	response, _ := json.Marshal(color)
	io.WriteString(w, string(response))
}

func (h *ColorsHandler) CreateColor(c context.Context, w http.ResponseWriter, r *http.Request) {

	if service.BeforeAuth(w, r) == false {
		return
	}

	name := r.FormValue("color[name]")
	code := r.FormValue("color[color_code]")
	textCode := r.FormValue("color[text_code]")

	// Validation
	errors := []string{}

	if utf8.RuneCountInString(code) <= 0 {
		errors = append(errors, "input ColorCode must be blank.")
	}
	if codeMatched, _ := regexp.MatchString("^#[0-9a-fA-F]{6}$", code); !codeMatched {
		errors = append(errors, "input ColorCode ex: #1a2b3c")
	}
	if utf8.RuneCountInString(textCode) <= 0 {
		errors = append(errors, "input TextCode must be blank.")
	}
	if textCodeMatched, _ := regexp.MatchString("^#[0-9a-fA-F]{6}$", textCode); !textCodeMatched {
		errors = append(errors, "input TextCode ex: #1a2b3c")
	}
	if utf8.RuneCountInString(name) <= 0 {
		errors = append(errors, "input Name must be blank.")
	}
	if nameMatched, _ := regexp.MatchString("^[0-9a-zA-Z]*$", name); !nameMatched {
		errors = append(errors, "input Name ex: Red100")
	}

	if len(errors) > 0 {
		helper.ResultMessageJSON(w, errors)
		return
	}

	// New
	color := model.Color{
		Name:     name,
		Code:     code,
		TextCode: textCode,
	}

	color.Commit()
	response, _ := json.Marshal(color)
	io.WriteString(w, string(response))
}
