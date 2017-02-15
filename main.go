package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	db "upper.io/db.v2"
	"upper.io/db.v2/postgresql"
)

var belts db.Collection

type Belt struct {
	ID          uint64 `db:"id,omitempty"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Fabric      string `db:"fabric"`
	Front       string `db:"front"`
	Back        string `db:"back"`
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var b []Belt
	err := belts.Find().All(&b)
	fmt.Println("%+v", b)
	fmt.Println("%+v", err)

	fmt.Fprint(w, "Welcome!\n")
}

func GetBelt(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var b Belt
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
