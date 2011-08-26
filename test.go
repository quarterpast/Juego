package main

import (
	"fmt"
	"reflect"
	"flag"
)

type Object interface{}
type Model struct {}
type Controller struct {}
type Page struct {
	title, content string
}
type Pages struct {Controller
	Page
}

func (c Controller) Render(args ...interface{}) {
	fmt.Println(args...)
}


func (this Pages) Index() {
	this.Render(this.title,this.content)
}

func main() {
	args := []reflect.Value{reflect.ValueOf(1)}
	page := Page{"test","this is a test page"}
	reflect.ValueOf(Pages{Page:page}).MethodByName(flag.Arg(0)).Call(args)
}