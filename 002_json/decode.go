package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type person struct {
	First string
	Last  string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/ret", ret)
	http.HandleFunc("/send", send)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First: "Tyler",
		Last:  "Souza",
	}

	p2 := person{
		First: "Sydeny",
		Last:  "Aussie",
	}

	xp := []person{p1, p2}

	err := json.NewEncoder(w).Encode(xp)
	if err != nil {
		fmt.Println(err)
	}
}

func ret(w http.ResponseWriter, r *http.Request) {

	xp := []person{}

	err := json.NewDecoder(r.Body).Decode(&xp)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(xp)
	tpl.ExecuteTemplate(w, "catch.html", xp)
}

func send(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "send.html", nil)
}
