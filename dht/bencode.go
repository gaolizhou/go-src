package main

import (
	"io/ioutil"
	"fmt"
	"log"
	//"strconv"
)
func decode_string(data []byte) (error int) {
  return 0;
}

func bencode(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Error Opening ", filename)
	}
	fmt.Print(data)
	for i:=0; i < len(data);i++  {
		switch data[i] {
		case 'd':
			//fmt.Println("dict")
		}
		//fmt.Println(data[i])
	}
	//fmt.Print(string(data))
}
