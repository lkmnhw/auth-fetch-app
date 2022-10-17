package main

import (
	"fetch-app/servers"
	"log"
)

func main() {
	s, err := servers.Init()
	if err != nil {
		log.Fatal(err)
		return
	}
	err = s.Start()
	if err != nil {
		log.Fatal(err)
	}
}
