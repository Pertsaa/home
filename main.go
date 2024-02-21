package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Project struct {
	Title       string
	Description string
	Link        string
}

var projects = []Project{
	{
		Title:       "This Website",
		Description: "Web Server",
		Link:        "https://github.com/Pertsaa/home",
	},
	{
		Title:       "Docs",
		Description: "SW Documentation CLI Tool",
		Link:        "https://github.com/Pertsaa/docs",
	},
	{
		Title:       "Full-stack Starter",
		Description: "Code Template",
		Link:        "https://github.com/Pertsaa/full-stack-starter",
	},
	{
		Title:       "TEA5767",
		Description: "Arduino Library",
		Link:        "https://github.com/Pertsaa/TEA5767",
	},
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5))

	r.Get("/robots.txt", handleRobots)
	r.Get("/static/*", handleStatic)

	r.Get("/", handleIndex)
	r.Get("/cv", handleCV)

	r.NotFound(handleRedirect)

	log.Fatal(http.ListenAndServe(":3000", r))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	mustExecuteTemplate(w, projects, "template/index.html")
}

func handleCV(w http.ResponseWriter, r *http.Request) {
	mustExecuteTemplate(w, nil, "template/cv.html")
}

func handleStatic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "max-age=31536000")
	http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))).ServeHTTP(w, r)
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusFound)
}

func handleRobots(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(""))
}

func mustExecuteTemplate(w http.ResponseWriter, data any, filenames ...string) {
	w.Header().Set("Content-Type", "text/html")

	tmpl, err := template.ParseFiles(filenames...)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}
