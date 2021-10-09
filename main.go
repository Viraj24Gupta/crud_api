package main

import (
	"flag"
	"log"
	"net"

	"crud_api/server"
)

func main() {
	var host = flag.String("host", "", "host http address to listen on")
	var	port = flag.String("port", "3000", "port number for http listener")
	flag.Parse()
	addr := net.JoinHostPort(*host, *port)
	if err := server.StartAPIServer(addr); err != nil {
		log.Fatal(err)
	}
}
