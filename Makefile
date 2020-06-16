bookdori-server: main.go
	go build -o bin/bookdori-server main.go

install-pkg:
	go get "github.com/codegangsta/negroni" 
	go get "github.com/julienschmidt/httprouter"