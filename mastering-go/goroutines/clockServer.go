package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	abort := make(chan bool)

	go startServer()
	go getConsoleCommand(abort)

	time.Sleep(2 * time.Second)

	go startClient()

	<-abort
}

func startServer() {
	listener, err := net.Listen("tcp", "localhost:8091")
	if err != nil {
		log.Fatal("Error listening.", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting.", err)
			continue
		}
		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))
		if err != nil {
			log.Print("Error Sending.", err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func getConsoleCommand(abort chan<- bool) {
	os.Stdin.Read(make([]byte, 1))
	abort <- true
}

func startClient() {
	conn, err := net.Dial("tcp", "localhost:8091")
	if err != nil {
		log.Fatal("Error connecting.", err)
	}
	defer conn.Close()
	_, err = io.Copy(os.Stdout, conn)
	if err != nil {
		log.Fatal("Error receiving.", err)
	}
}
