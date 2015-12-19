package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path"

	"github.com/julienschmidt/httprouter"
)

type Post struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	r := httprouter.New()
	r.GET("/", Home)

	// Post List
	r.GET("/posts", All)
	r.POST("/posts", Create)
	r.GET("/posts/:id", One)
	r.PUT("/posts/:id", Update)
	r.GET("/posts/:id/update", Edit)

	fmt.Println("Starting server on localhost:8080")
	http.ListenAndServe(":8080", r)
}

func Home(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	post := Post{
		"Building Web Apps with Go",
		"Jeremy Saenz",
	}

	fp := path.Join("templates", "index.html")
	tmpl, _ := template.ParseFiles(fp)
	tmpl.Execute(rw, post)
}

func All(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	post := Post{
		"Building Web Apps with Go",
		"Jeremy Saenz",
	}

	js, _ := json.Marshal(post)

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)
}

func Create(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "Create")
}

func One(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(rw, "Here is the id you entered " + p.ByName("id"))
}

func Update(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "Update")
}

func Edit(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "Edit")
}
