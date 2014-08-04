package main

import (
	"lets-go/api"
	"lets-go/views"
	"net/http"
)

func ConnectHandlers() {
	http.HandleFunc("/", views.IndexView)

	http.HandleFunc("/api/people", api.PeopleResource)
	http.HandleFunc("/api/people/add", api.PersonAddResource)
	http.HandleFunc("/api/people/like", api.PersonLikeResource)

    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.Handle("/partials/", http.StripPrefix("/partials/", http.FileServer(http.Dir("static/app/partials/"))))
}

func main() {
	ConnectHandlers()
	http.ListenAndServe(":8000", nil)
}
