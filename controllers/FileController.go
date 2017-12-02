package controllers

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/aqrun/photoweb/configs"
	"github.com/aqrun/photoweb/helpers"
)


type FileController struct{

}


// index
func ActionFileIndex(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir(configs.UploadDir)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name())
	}
	locals["images"] = images
	if err := helpers.RenderHtml(w, "list", locals); err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}


// /upload
func ActionFileUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if err := helpers.RenderHtml(w, "upload", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		filename := h.Filename
		defer f.Close()
		t, err := os.Create(configs.UploadDir + "/" + filename)
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
func ActionFileView(w http.ResponseWriter, r *http.Request){
	imageId := r.FormValue("id")
	imagePath := configs.UploadDir + "/" + imageId
	if exists := helpers.IsFileExist(imagePath); !exists {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

