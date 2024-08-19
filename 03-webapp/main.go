package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Todo struct {
	ID    string
	Title string
	Done  bool
}

var todos = []Todo{}

func main() {
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/add", AddHandler)
	http.HandleFunc("/remove", RemoveHandler)
	http.HandleFunc("/active", ActiveHandler)

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		todo := Todo{
			ID:    fmt.Sprint(time.Now().Unix()),
			Title: title,
		}
		todos = append(todos, todo)
	}

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func RemoveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := r.URL.Query().Get("id")

		for i, todo := range todos {
			if todo.ID == id {
				todos = append(todos[:i], todos[i+1:]...)
				break
			}
		}
	}

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func ActiveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := r.URL.Query().Get("id")

		for i, todo := range todos {
			if todo.ID == id {
				todos[i].Done = !todo.Done
				break
			}
		}

		fmt.Println(todos)
	}

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/home" {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}

	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
