package main

import (
	"flag"
	"log"
	"strconv"
	"strings"
)

func main() {
	var scanPort string
	flag.StringVar(&scanPort, "p", "", "-p <port> or -p <port>-<port>: Prove a port or port range to scan.")
	flag.Parse()

	// check for correct arguments
	args := flag.Args()
	if len(args) != 1 {
		log.Fatal("Incorrect arguments provided.")
	}

	if scanPort == "" {
		log.Fatal("No port specified.")
	}

	//function to check if a port is valid
	getPort := func(portStr string) bool {
		port, err := strconv.Atoi(portStr)
		if err == nil {
			return false
		}
		return port >= 1 && port <= 65535
	}

	ports := strings.Split(scanPort, "-")
	switch len(ports) {
	case :
		
	}

	

}
