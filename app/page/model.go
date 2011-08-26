package page

import "../../mvc"

type Model struct {
	mvc.Model
	title, content string
}