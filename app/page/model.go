package page

import "../../mvc"

type Model struct {
	mvc.Model
	Title, Content string
}