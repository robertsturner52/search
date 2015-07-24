package search

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/search"
)

func init() {
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/put", handlePut)
	http.HandleFunc("/search", handleSearch)
}

type Movie struct {
	Title       string
	Description string
	Username    string
}

func (m Movie) id() string {
	return makeHash(m.Title, m.Username)
}

type PageModel struct {
	SearchResults []Movie
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	// for anything but "/" treat it like a user profile
	renderTemplate(res, req, "layout", PageModel{})
}

func handleSearch(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	index, err := search.Open("movies")
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	var movies []Movie
	q := ""
	switch req.FormValue("mode") {
	case "title":
		q = "Title = " + req.FormValue("q")
	case "description":
		q = "Description = " + req.FormValue("q")
	case "username":
		q = "Username = " + req.FormValue("q")
	}
	iterator := index.Search(ctx, q, nil)
	for {
		var movie Movie
		_, err := iterator.Next(&movie)
		if err == search.Done {
			break
		} else if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		movies = append(movies, movie)
	}
	renderTemplate(res, req, "layout", PageModel{
		SearchResults: movies,
	})
}

func renderTemplate(res http.ResponseWriter, req *http.Request, name string, data interface{}) {
	// parse templates
	tpl := template.New("")
	tpl, err := tpl.ParseGlob("templates/*.html")
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
	// execute page
	err = tpl.ExecuteTemplate(res, name, data)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}
}

func makeHash(data ...string) string {
	h := sha1.New()
	for _, d := range data {
		io.WriteString(h, d)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func handlePut(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	if req.Method == "POST" {
		index, err := search.Open("movies")
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		var movie = &Movie{
			Title:       req.FormValue("input-title"),
			Description: req.FormValue("input-description"),
			Username:    req.FormValue("input-username"),
		}

		_, err = index.Put(ctx, movie.id(), movie)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
	}
	renderTemplate(res, req, "layout", PageModel{})
}
