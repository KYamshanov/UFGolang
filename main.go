package main

import (
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("templates/index.html")
	tmp.Execute(w, index)
}

func request(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	server := http.FileServer(http.Dir("templates/css"))
	prefix := http.StripPrefix("/css/", server)
	mux.Handle("/css/", prefix)

	http.ListenAndServe(":80",mux)
}

func main(){
	request()
}