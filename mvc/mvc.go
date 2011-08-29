package mvc

import (
	"fmt"
//	"reflect"
	"io"
	"runtime"
	"os"
	"strings"
)
type Persistent interface{
	Save() os.Error
}
type Model struct {}
type Controller struct {
	Writer io.Writer
}
func (m Model) Save() os.Error {
	return nil
}
func Caller() (string, bool) {
	if pc,_,_,ok := runtime.Caller(2); ok {
		return strings.Split(runtime.FuncForPC(pc).Name(),"·")[1], true
	}
	return "", false
}
func (c Controller) Render(args map[string]interface{}) os.Error {
	if caller,ok := Caller(); ok {
		//TODO: template shit
		fmt.Fprintln(c.Writer,caller)
		return nil
	}
	return os.NewError("oh snap")
}
func (c Controller) RenderJSON() args map[string]interface{}) os.Error {

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