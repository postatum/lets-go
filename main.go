package main

import (
    "net/http"
    "lets-go/views"
)

func ConnectHandlers() {
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
    http.HandleFunc("/api/people", views.PeopleResource)
    http.HandleFunc("/api/people/add", views.PersonAddResource)
    http.HandleFunc("/", views.PeopleView)
}

func main() {
    ConnectHandlers()
    http.ListenAndServe(":8000", nil)
}
