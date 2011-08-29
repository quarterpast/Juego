all: test

test: mvc.6 page.6 test.6
	6l -o test test.6
test.6: test.go 
	6g -o test.6 test.go

mvc.6:
	cd mvc && make

page.6:
	cd app/page && make