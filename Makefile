export GOPATH=$(shell pwd)
run: 
	go run src/webapp/main.go

edit:
	@code .
