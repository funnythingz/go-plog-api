package helper

import (
	"encoding/json"
	"io"
	"net/http"
)

func ResultMessageJSON(w http.ResponseWriter, messages []string) {
	response, _ := json.Marshal(map[string][]string{"message": messages})
	io.WriteString(w, string(response))
}
