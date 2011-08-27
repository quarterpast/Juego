package mvc

import (
	"fmt"
//	"reflect"
//	"flag"
//	"http"
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
func NewModel() *Model {
	return &Model{Persist{}}
}
func (m Model) Save() {
}
func (c Controller) Render(args ...interface{}) {
	fmt.Fprintln(c.writer,args...)
}
/*func handler(w http.ResponseWriter, r *http.Request) {
	for k,v := range r.URL.Path[1:] {
		
	}
}
func main() {
	//page := NewPage("test","this is a test page")
	//cont := Pages{Page:page}
	http.HandleFunc("/",handler)
	http.ListenAndServe(":10000",nil)
}*/