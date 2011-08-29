package main

import (
	"./mvc"
	"./page"
//	"os"
	"http"
//	"fmt"
	"strings"
	"reflect"
)
var controllers = map[string] reflect.Type {
	"page": reflect.TypeOf(page.Controller{}),
}
func Handler(w http.ResponseWriter, r *http.Request) {
	comps := strings.Split(r.URL.Path[1:],"/")
	/*fmt.Fprintln(os.Stderr,reflect.ValueOf(w))
	controller := reflect.New(controllers[comps[0]])
	reflect.Indirect(controller).FieldByName("writer").Set(reflect.ValueOf(w))*/
	controller := page.Controller{
		mvc.Controller{w},
		page.Model{
			mvc.Model{},
			"test",
			"this is a test page",
		},
	}
	var values []reflect.Value
	if len(comps)>2 {
		values = make([]reflect.Value,len(comps)-2)
		for k,v := range comps[2:] {
			values[k] = reflect.ValueOf(v)
		}
	} else {
		values = []reflect.Value{}
	}
	reflect.ValueOf(controller).MethodByName(comps[1]).Call(values)
}
func main() {
	http.HandleFunc("/",Handler)
	http.ListenAndServe(":10000",nil)
}