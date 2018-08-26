package main

import (
	"fmt"
)

func main() {
	//bencode("ubuntu.torrent")
	//var seed Torrent
	//seed.DecodeFromFile("ubuntu.torrent")

	//json_test()
	fmt.Println("Hello World!")
	fmt.Println(DecodeInteger([]byte(`i2097152e`), 0))
	fmt.Println(DecodeString([]byte(`13:announce-list`), 0))
}

