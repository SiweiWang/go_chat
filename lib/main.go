package lib

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const port = "8080"

// RunGuest takes an destination ip and connect to it
func RunGuest(ip string) {
	ipAndPort := ip + ":" + port
	conn, dialErr := net.Dial("tcp", ipAndPort)
	if dialErr != nil {
		log.Fatal("Error:", dialErr)
	}
	for {
		handleGuest(conn)
	}
}

func handleGuest(conn net.Conn) {
	fmt.Print("Send message: ")
	reader := bufio.NewReader(os.Stdin)
	message, readErr := reader.ReadString('\n')
	if readErr != nil {
		log.Fatal("Error:", readErr)
	}
	fmt.Fprint(conn, message)

	replyReader := bufio.NewReader(conn)
	replyMessage, replyErr := replyReader.ReadString('\n')
	if replyErr != nil {
		log.Fatal("Error: ", replyErr)
	}

	fmt.Println("Message recived: ", replyMessage)

}

// RunHost takes an host ip as agrument and listens for connections
func RunHost(ip string) {
	ipAndPort := ip + ":" + port
	listener, listenErr := net.Listen("tcp", ipAndPort)
	if listenErr != nil {
		log.Fatal("Error:", listenErr)
	}

	fmt.Println("listening on ", ipAndPort)

	conn, acceptErr := listener.Accept()
	if acceptErr != nil {
		log.Fatal("Error:", listenErr)
	}

	fmt.Println("New connection accpeted")
	for {
		handleHost(conn)
	}
}

func handleHost(conn net.Conn) {
	reader := bufio.NewReader(conn)
	message, readErr := reader.ReadString('\n')
	if readErr != nil {
		log.Fatal("Error:", readErr)
	}
	fmt.Println("Message recorded: ", message)
	fmt.Print("Send message: ")
	replyreader := bufio.NewReader(os.Stdin)
	replyMessage, replyErr := replyreader.ReadString('\n')
	if replyErr != nil {
		log.Fatal("Error: ", replyErr)
	}

	fmt.Fprint(conn, replyMessage)
}
