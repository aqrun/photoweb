package helpers

import (
	"os"
	"net/http"
)

func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func RenderHtml(w http.ResponseWriter, tmpl string,
	locals map[string]interface{}) (err error) {
	err = Templates[tmpl].Execute(w, locals)
	return
}