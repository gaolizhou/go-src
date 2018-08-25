package main

import (
	"fmt"
)

func main() {
	//bencode("ubuntu.torrent")
	var seed Torrent
	seed.DecodeFromFile("ubuntu.torrent")

	//json_test()
	fmt.Println("Hello World!")
}

