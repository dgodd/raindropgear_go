package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgodd/raindropgear_go/model"
	"github.com/dgodd/raindropgear_go/tpl"
	"github.com/julienschmidt/httprouter"
	db "upper.io/db.v2"
	"upper.io/db.v2/postgresql"
)

var belts db.Collection

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var b []model.Belt
	err := belts.Find().All(&b)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Could not read database"))
	}
	// fmt.Println("%+v", b)
	// fmt.Println("%+v", err)

	// fmt.Fprint(w, "Welcome!\n")
	w.Write([]byte(tpl.Index(b)))
}

func GetBelt(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var b model.Belt
	err := belts.Find("slug", ps.ByName("slug")).One(&b)
	fmt.Println("%+v", b)
	fmt.Println("%+v", err)

	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("slug"))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/belts/:slug", GetBelt)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func init() {
	settings := postgresql.ConnectionURL{
		Host: "127.0.0.1",
		// User:     "foo",
		// Password: "bar",
		Database: "raindropgear_rust",
	}
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal(err)
	}
	belts = sess.Collection("belts")
}
