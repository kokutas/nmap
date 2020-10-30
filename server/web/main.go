package main

import (
	"log"
	"nmap/server/web/service/network"
)

func main() {
	adaptors,err:=network.Adaptors()
	if err != nil {
		log.Fatal(err)
	}
	for i, adaptor := range adaptors {
		log.Println(i,adaptor)
	}
}
