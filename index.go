package main

import (
	"log"

	"github.com/aqrun/photoweb/controllers"
	"html/template"
	"io/ioutil"
	"github.com/aqrun/photoweb/configs"
	"path"
	"strings"
	"net/http"
)

var Templates = make(map[string]*template.Template)

func init() {
	fileInfoArr,err := ioutil.ReadDir(configs.TemplateDir)
	if err != nil {
		panic(err)
		return
	}
	var templateName, templatePath string
	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}
		templatePath = configs.TemplateDir + "/" + templateName
		log.Println("loading template:", templatePath)
		t := template.Must(template.ParseFiles(templatePath))
		Templates[strings.Split(templateName, ".")[0]] = t
	}
	log.Println(Templates)
}


func main(){
	http.HandleFunc("/", controllers.ActionFileIndex)
	http.HandleFunc("/upload", controllers.ActionFileUpload)
	http.HandleFunc("/view", controllers.ActionFileView)


	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
	log.Println("Listen at http://localhost:8081")
}
