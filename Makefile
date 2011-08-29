all: test
	
test: test.6 mvc.6 page.6
	6l -o test test.6
test.6: test.go
	6g -o test.6 test.go

mvc.6: mvc/mvc.go
	cd mvc && make
page.6: app/page/controller.go app/page/model.go
	cd app/page && make