package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"rest-api-go/internal/user"
)

func main() {
	log.Println("create router")
	router := httprouter.New()

	log.Println("create handler")
	handler := user.NewHandler()
	handler.Register(router)

	startServer(router)
}

func startServer(router *httprouter.Router) {
	log.Println("start server")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler: router,
	}

	log.Fatal(server.Serve(listener))
}
