package page

import "../../mvc/_obj/mvc"

type Controller mvc.Controller

func (this Controller) Index() {
	this.Render(this.title,this.content)
}