all: mvc
	gomake -f go.Makefile
mvc:
	cd mvc && gomake