package views

import (
    "fmt"
    "net/http"
    "labix.org/v2/mgo"
    "labix.org/v2/mgo/bson"
    "encoding/json"
    "html/template"
)

const (
    mongourl = "127.0.0.1:27017"
    mongodb = "godb"
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
    w.Header().Set("Content-Type", "application/json")

    session, _ := mgo.Dial(mongourl)
    defer session.Close()
    collection := session.DB(mongodb).C("people")

    people := make([]map[string]interface{}, 5)
    _ = collection.Find(bson.M{}).Sort("-_id").Limit(5).All(&people)

    fmt.Fprint(w, JsonResponse{"people": people})
}

func PersonAddResource(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    person := make(map[string]interface{})
    person["name"] = r.FormValue("name")
    person["email"] = r.FormValue("email")

    session, _ := mgo.Dial(mongourl)
    defer session.Close()
    collection := session.DB(mongodb).C("people")

    if err := collection.Insert(&person); err != nil {
        fmt.Fprint(w, JsonResponse{"err": err.Error()})
        return
    }
    fmt.Fprint(w, JsonResponse{"success": "true"})
}

func PeopleView(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type", "text/html")
    t, _ := template.ParseFiles("templates/index.html")

    context := make(map[string]interface{})

    t.Execute(w, &context)
}
