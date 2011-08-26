package main

import (
	"fmt"
//	"reflect"
//	"flag"
	"http"
	"io"
)

type Object interface{}
type Persist struct{}
type Model struct {
	Persist
}
type Controller struct {
	writer io.Writer
}
type Page struct {
	Model
	title, content string

}
func NewPage(title, content string) Page {
	return Page{Model{Persist{}},title,content}
}
type Pages struct {
	Controller
	Page
}

func (m Model) Save() {
}
func (c Controller) Render(args ...interface{}) {
	fmt.Fprintln(c.writer,args...)
}


func (this Pages) Index() {
	this.Render(this.title,this.content)
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
func main() {
	//page := NewPage("test","this is a test page")
	//cont := Pages{Page:page}
	http.HandleFunc("/",handler)
	http.ListenAndServe(":10000",nil)
}