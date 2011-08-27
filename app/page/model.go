package page

import "../../mvc/_obj/mvc"

type Model struct {
	mvc.Model
	title, content string
}