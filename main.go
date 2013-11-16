package main

import (
	"lets-go/views"
	"net/http"
)

func ConnectHandlers() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.HandleFunc("/api/people", views.PeopleResource)
	http.HandleFunc("/api/people/add", views.PersonAddResource)
	http.HandleFunc("/api/people/like", views.PersonLikeResource)
	http.HandleFunc("/", views.PeopleView)
}

func main() {
	ConnectHandlers()
	http.ListenAndServe(":8000", nil)
}
