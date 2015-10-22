package main

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var pages *template.Template

func init() {
	router := httprouter.New()
	router.GET("/", showIndex)
	router.GET("/resume", showResume)

	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))
	http.Handle("/", router)
	pages = template.Must(template.ParseFiles("index.html"))
}

func showIndex(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	err := pages.ExecuteTemplate(res, "index.html", nil)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func showResume(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	http.Redirect(res, req, "public/Resume-Allen_Mills.pdf", 302)
}
