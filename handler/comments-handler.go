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

type CommentsHandler struct{}

func (h *CommentsHandler) Comments(c context.Context, w http.ResponseWriter, r *http.Request) {

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

	Comments := model.CommentList{}
	Comments.Fetch(permit, page)
	response, _ := json.Marshal(Comments)
	io.WriteString(w, string(response))
}

func (h *CommentsHandler) Comment(c context.Context, w http.ResponseWriter, r *http.Request) {

	if service.BeforeAuth(w, r) == false {
		return
	}

	id, _ := strconv.Atoi(pat.Param(c, "id"))
	comment := model.Comment{}
	comment.Fetch(id)
	if comment.Id == 0 {
		helper.ResultMessageJSON(w, []string{fmt.Sprintf("Not Found: %d", id)})
		return
	}
	response, _ := json.Marshal(comment)
	io.WriteString(w, string(response))
}

func (h *CommentsHandler) CreateComment(c context.Context, w http.ResponseWriter, r *http.Request) {

	if service.BeforeAuth(w, r) == false {
		return
	}

	content := r.FormValue("comment[content]")
	plogId, _ := strconv.Atoi(r.FormValue("plog[plog_id]"))

	// Validation
	errors := []string{}

	if utf8.RuneCountInString(content) <= 0 {
		errors = append(errors, "input Content must be blank.")
	}
	if utf8.RuneCountInString(content) < 1 || utf8.RuneCountInString(content) > 1000 {
		errors = append(errors, "input Content minimum is 1 and maximum is 1000 character.")
	}

	if len(errors) > 0 {
		helper.ResultMessageJSON(w, errors)
		return
	}

	comment := model.Comment{
		PlogId:  plogId,
		Content: content,
	}

	comment.Commit()
	response, _ := json.Marshal(comment)
	io.WriteString(w, string(response))
}
