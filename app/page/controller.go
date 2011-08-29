package page

import "../../mvc"

type Controller struct{
	mvc.Controller
	Model
}

func (this Controller) Index() {
	this.Render(this.Title,this.Content)
}
func (this Controller) Stuff(a string,b int) {
	this.Render(a,b)
}