package main

import (
	"flag"
)

func main() {
	port := flag.Uint("port", 5001, "TCP Port Number for Blockchain Server")
	flag.Parse()
	app := NewBlockchainServer(uint16(*port))
	app.Run()
}
