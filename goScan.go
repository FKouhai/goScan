package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
)

func main() {
	//	tcp_scan("nmap.org", "80")
	var H = flag.String("h", "host", "localhost")
	var P = flag.String("p", "port", "80")
	flag.Parse()
	if *P == "a" {
		for i := 1; i < 65535; i++ {
			tcp_scan(*H, strconv.Itoa(i))
		}
	} else {
		tcp_scan(*H, *P)
	}
}

func tcp_scan(ip_addr string, port string) {
	_, err := net.DialTimeout("tcp", ip_addr+":"+port, 5*1000*1000*100)
	if err != nil {
		//		fmt.Println(err)
		//fmt.Println(port, " Is not reachable")
	} else {
		fmt.Println("The port", port, "is open")
	}
}
