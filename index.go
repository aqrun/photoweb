package main

import (
	"log"
	"net/http"

	"github.com/aqrun/photoweb/controllers"
)

func main(){
	http.HandleFunc("/", controllers.ActionIndex)
	http.HandleFunc("/upload", controllers.ActionUpload)
	http.HandleFunc("/view", controllers.ActionView)


	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
	log.Println("Listen at http://localhost:8081")
}
