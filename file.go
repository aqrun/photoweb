package main

import (
	"io"
	"net/http"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		io.WriteString(w, `<form method="post" action="upload" enctype="multipart/form-data">
			Choose and image to upload: <input type="file" name="image"/>
			<input type="submit" value="Upload"/>
			</form>
		`)
		return
	}
}