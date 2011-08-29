package page

import "../../mvc"

type Controller struct{
	mvc.Controller
	Model
}

func (this Controller) Index() {
	this.Render(map[string]interface{}{"title":this.Title,"content":this.Content})
}
func (this Controller) Stuff(a string,b int) {
	this.Render(map[string]interface{}{"a":a,"b":b})
}