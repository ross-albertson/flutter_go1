// name_server1 project main.go
package main

import (
	//"fmt"
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Name struct {
	ID        bson.ObjectId `bson:"_id" json:"_id"`
	FirstName string        `bson:"firstName" json:"firstName"`
	LastName  string        `bson:"lastName" json:"lastName"`
}

func main() {
	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		session, err1 := mgo.Dial("localhost")
		if err1 != nil {
			http.Error(w, err1.Error(), 500)
			return
		}
		c := session.DB("rpc1").C("names")
		var result []Name
		err := c.Find(nil).All(&result)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		json.NewEncoder(w).Encode(&result)

	})

	if e := http.ListenAndServe("10.0.0.61:8000", nil); e != nil {
		log.Fatalln(e.Error())
	}

}
