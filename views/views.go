package views

import (
	"encoding/json"
	"fmt"
	"html/template"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"net/http"
)

const (
	mongourl   = "127.0.0.1:27017"
	dbname     = "godb"
	collection = "people"
	PER_PAGE   = 5
)

type JsonResponse map[string]interface{}

func (r JsonResponse) String() (s string) {
	b, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(b)
}

func PeopleResource(w http.ResponseWriter, r *http.Request) {
	session, _ := mgo.Dial(mongourl)
	defer session.Close()
	collection := session.DB(dbname).C(collection)

	people := make([]map[string]interface{}, PER_PAGE)
	_ = collection.Find(bson.M{}).Sort("-_id").Limit(PER_PAGE).All(&people)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, JsonResponse{"people": people})
}

func PersonAddResource(w http.ResponseWriter, r *http.Request) {
	person := make(map[string]interface{})
	person["name"] = r.FormValue("name")
	person["email"] = r.FormValue("email")

	session, _ := mgo.Dial(mongourl)
	defer session.Close()
	collection := session.DB(dbname).C(collection)

	if err := collection.Insert(&person); err != nil {
		fmt.Fprint(w, JsonResponse{"err": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, JsonResponse{"success": "true"})
}

func PeopleView(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")

	context := make(map[string]interface{})

	w.Header().Set("Content-type", "text/html")
	t.Execute(w, &context)
}
