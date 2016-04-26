package main

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var pages *template.Template

func init() {
	router := httprouter.New()
	router.GET("/", showIndex) // Local links
	router.GET("/resume", showResume)
	router.GET("/tutoring", showTutoring)
	router.GET("/projects", showProjects)

	router.GET("/github", showGithub) // Outbound links
	router.GET("/linkedin", showLinkedin)

	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))
	http.Handle("/", router)

	pages = template.Must(pages.ParseGlob("html/*.html"))
}

//// -------------------------
// Helper Functions
/////

func HandleError(res http.ResponseWriter, e error) {
	// generic error handling for any error we encounter.
	if e != nil {
		http.Error(res, e.Error(), http.StatusInternalServerError)
	}
}

func ServeTemplateWithParams(res http.ResponseWriter, req *http.Request, templateName string, params interface{}) {
	// simple func to cut down on repeating code.
	err := pages.ExecuteTemplate(res, templateName, &params)
	HandleError(res, err)
}

//// -------------------------
// Local Pages/Files
/////

func showIndex(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	ServeTemplateWithParams(res, req, "index.html", nil)
}

func showTutoring(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	ServeTemplateWithParams(res, req, "tutoring.html", nil)
}

func showProjects(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	http.Redirect(res, req, "/", http.StatusTemporaryRedirect) // Leaving a spot for later. we'll get to this.
}

func showResume(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	http.Redirect(res, req, "public/Resume-Allen_Mills.pdf", 302)
}

//// -------------------------
// Outbound Links
/////

func showGithub(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	http.Redirect(res, req, "https://github.com/FelixVicis", http.StatusTemporaryRedirect)
}

func showLinkedin(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	http.Redirect(res, req, "https://www.linkedin.com/in/allenjmills", http.StatusTemporaryRedirect)
}
