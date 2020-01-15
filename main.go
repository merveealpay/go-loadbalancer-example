package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

//config parameters, editable
var (
	counter       int
	listenAddress = "localhost:8080" //loadbalancer will listen to
	//our servers:
	server = []string{
		"localhost:5001",
		"localhost:5002",
		"localhost:5003",
	}
)

func main() {
	listener, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatalf("failed to listen %s", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept connection : %s ", err)
		}

		backend := chooseServer()
		fmt.Printf("counter=%d backend=%s\n", counter, backend)

		go func() {
			err := proxy(backend, conn)
			if err != nil {
				log.Printf("Warnin: proxying failed: %v", err)
			}
		}()
	}

}

func proxy(backend string, incomingConn net.Conn) error {

	backendConn, err := net.Dial("tcp", backend)
	if err != nil {
		return fmt.Errorf("failed to connect to backend: %s: %v", backend, err)
	}

	//arka ön tarafa cevap yollarken ön taraftan da arka tarafa cevap yollamamız gerekiyor
	//bu yüzden bunları go-routine ile yapalim
	//incomingConn ->backendCon
	go io.Copy(backendConn, incomingConn)
	//backendConn -> incomingConn
	go io.Copy(incomingConn, backendConn)

	return nil
}

func chooseServer() string {
	s := server[counter%len(server)]
	counter++
	return s
}
