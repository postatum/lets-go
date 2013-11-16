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
	mongourl        = "127.0.0.1:27017"
	dbname          = "godb"
	collection_name = "people"
	perpage         = 5
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
	collection := session.DB(dbname).C(collection_name)

	people := make([]map[string]interface{}, perpage)
	_ = collection.Find(bson.M{}).Sort("-_id").Limit(perpage).All(&people)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, JsonResponse{"people": people})
}

func PersonAddResource(w http.ResponseWriter, r *http.Request) {
	person := make(map[string]interface{})
	person["name"] = r.FormValue("name")
	person["email"] = r.FormValue("email")

	session, _ := mgo.Dial(mongourl)
	defer session.Close()
	collection := session.DB(dbname).C(collection_name)

	if err := collection.Insert(&person); err != nil {
		fmt.Fprint(w, JsonResponse{"err": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, JsonResponse{"success": "true"})
}

func LikePersonResource(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
	session, _ := mgo.Dial(mongourl)
	defer session.Close()
	collection := session.DB(dbname).C(collection_name)

	pid := bson.ObjectIdHex(r.FormValue("pid"))
	selector := bson.M{"_id": pid}

	person := make(map[string]interface{})
	if err := collection.Find(selector).One(&person); err != nil {
		fmt.Fprint(w, JsonResponse{"err": err.Error()})
		return
	}
	change := mgo.Change{
		Update: bson.M{"$inc": bson.M{"rating": 1}},
	}
	if _, err := collection.Find(selector).Apply(change, &person); err != nil {
		fmt.Fprint(w, JsonResponse{"err": err.Error()})
		return
	}

	fmt.Fprint(w, JsonResponse{"success": "true"})
}

func PeopleView(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")

	context := make(map[string]interface{})

	w.Header().Set("Content-type", "text/html")
	t.Execute(w, &context)
}
