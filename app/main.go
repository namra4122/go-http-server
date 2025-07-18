package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

type Request struct {
	Method string
	Path string
	Version string
	Host string
	Headers map[string]string
}

type Response struct {
	Status int
	Body []byte
	Version string
	Headers map[string]string
}

func extractHeader(hrs []string) map[string]string {
	header_map := make(map[string]string)
	for idx,hr := range(hrs) {
		fmt.Printf("IDX: %v - Headers: %v\n", idx, hr)
	}
	return header_map
}

func parseRequest(request string) *Request{
	lines := strings.Split(request, "\r\n");
	fmt.Printf("==== this is the request Lines: %#v\n", lines)
	reqLine := strings.Split(lines[0], " ")
	fmt.Printf("==== this is the request ReqLine: %#v\n", reqLine)
	header := extractHeader(lines[1:len(lines)])
	return &Request{
		Method: reqLine[0],
		Path: reqLine[1],
		Version: reqLine[2],
		Headers: header,
	}
}

func handleConnection(conn net.Conn) {
	requestBuffer := make([]byte, 512)
	n,err := conn.Read(requestBuffer)

	if err != nil {
		fmt.Println("Error while reading the request: ", err)
	}

	fmt.Println(string(requestBuffer[:n]))

	requestStruct := parseRequest(string(requestBuffer[:n]))
	fmt.Println(requestStruct)
}

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleConnection(conn)
	}
}
