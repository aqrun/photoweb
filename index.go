package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		io.WriteString(w, "This is index")
	})
	http.HandleFunc("/upload", UploadHandler)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
	fmt.Println("Listen at http://localhost:8081")
}
