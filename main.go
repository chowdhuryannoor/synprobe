package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

func main() {
	var portArg string
	flag.StringVar(&portArg, "p", "", "-p <port> or -p <port>-<port>: Prove a port or port range to scan.")
	flag.Parse()

	//check for correct arguments
	args := flag.Args()
	if len(args) != 1 {
		log.Fatal("Incorrect arguments provided.")
	}

	//get ip to scan
	ip := args[0]
	if portArg == "" {
		log.Fatal("No port specified.")
	}
	if ip != "localhost" && net.ParseIP(ip) == nil {
		log.Fatal("Invalid IP Address.")
	}

	//if -p not provided, scan common ports
	if flag.Lookup("p") == nil {
		ports := []int{21, 22, 23, 25, 80, 110, 143, 443, 587, 853, 993, 3389, 8080}

		for _, port := range ports {
			scanPort(ip, port)
		}
		return
	} else {
		ports := strings.Split(portArg, "-")
		switch len(ports) {
		case 1:
			//scan one port
			port, err := getPort(ports[0])
			if err != nil {
				log.Fatal(err)
			}
			scanPort(ip, port)
		case 2:
			//get start port
			startPort, err := getPort(ports[0])
			if err != nil {
				log.Fatal(err)
			}

			//get end port
			endPort, err := getPort(ports[1])
			if err != nil {
				log.Fatal(err)
			}

			//check for valid range
			if startPort > endPort {
				log.Fatal("Invalid port range.")
			}

			//scan port range
			for i := startPort; i <= endPort; i++ {
				scanPort(ip, i)
			}

		default:
			log.Fatal("Too many ports found")
		}
	}
}

func getPort(portStr string) (int, error) {
	port, err := strconv.Atoi(portStr)
	if err == nil {
		return -1, err
	}
	if port >= 1 && port <= 65535 {
		return -1, errors.New("invalid port")
	}

	return port, nil
}

func scanPort(ip string, port int) {
	//try to establish a conneciton
	conn, err := net.Dial("tcp", ip+":"+strconv.Itoa(port))
	if err != nil {
		return
	}
	defer conn.Close()

	timeout := time.After(3 * time.Second)
	select {
	case <-timeout:
		//probe the server
		fmt.Println("Timed out.")
	default:
		//read the first 1024 bytes from server
		data := make([]byte, 1024)
		n, err := conn.Read(data)
		if err != nil {
			fmt.Println("Error reading data from server: ", err)
		}
		fmt.Println("Data recieved from server: ", string(data[:n]))
	}
}
