package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"bufio"
)

var host = flag.String("host", "", "host")
var port = flag.String("port", "1234", "port")

func main() {
	flag.Parse()
	var listener net.Listener
	var err error
	listener, err = net.Listen("tcp", *host + ":" + *port)
	if err != nil {
		fmt.Println("error listening", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Listening on " + *host + ":" + *port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err)
			os.Exit(1)
		}
		fmt.Printf("Recv msg %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
		go EchoByIoCopy(conn)
		//		go EchoByBufio(conn)
	}
}

func EchoByIoCopy(conn net.Conn) {
	defer conn.Close()
	_, err := io.Copy(conn, conn) // A successful Copy returns err == nil, not err == EOF. so just return can work

	fmt.Println("Echo over ", err)
}

func EchoByBufio(conn net.Conn) {
	defer conn.Close()
	var err error
	var line string
	for {
		line, err = bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			break
		}
		_, err = conn.Write([]byte(line))
		if err != nil {
			break
		}
	}
	fmt.Println("Echo over ", err)
}
