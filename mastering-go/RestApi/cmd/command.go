package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello RestApi!")
	startServer1()
}

//--------------------------------------- server by implement http.Handler interface

type RestServer struct {
	listenPort	int
}

func (rs RestServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"name" : "ali"}`))
}

func startServer1() {
	server := RestServer{8081}
	log.Fatal(http.ListenAndServe(":8081", server))
}

func startServr2() {
}