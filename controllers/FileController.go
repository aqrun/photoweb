package controllers

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/aqrun/photoweb/configs"
	"github.com/aqrun/photoweb/helpers"
)



// index
func ActionIndex(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir(configs.UPLOAD_DIR)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var listHtml string
	for _, fileInfo := range fileInfoArr {
		imgid := fileInfo.Name()
		listHtml += "<li><a href=\"/view?id=" + imgid + "\">" + imgid + "</a></li>"
	}
	io.WriteString(w, "<html><ol>"+ listHtml +"</ol></html>")
}


// /upload
func ActionUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		io.WriteString(w, `
<html>
	<form method="post" action="upload" enctype="multipart/form-data">
		Choose and image to upload: <input type="file" name="image"/>
		<input type="submit" value="Upload"/>
	</form>
</html>
		`)
		return
	}

	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		filename := h.Filename
		defer f.Close()
		t, err := os.Create(configs.UPLOAD_DIR + "/" + filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer t.Close()
		if _, err := io.Copy(t, f); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/view?id=" + filename, http.StatusFound)
	}


}

// view
func ActionView(w http.ResponseWriter, r *http.Request){
	imageId := r.FormValue("id")
	imagePath := configs.UPLOAD_DIR + "/" + imageId
	if exists := helpers.IsFileExist(imagePath); !exists {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

