package main

import (
	"io/ioutil"
	"log"
	"fmt"
)

type Torrent struct {
	announce string
	announce_list [][]string
	comment string
	creation_date int64
	info struct {
		length int64
		name string
		piece_length int64
		pieces [][20]byte
	}
}

func (this Torrent) DecodeFromFile(filename string) (err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Error Opening ", filename)
	}
	fmt.Print(data)
	for i := 0; i < len(data); i++ {
		switch data[i] {
		case 'd':
			//fmt.Println("dict")
		}
		//fmt.Println(data[i])
	}
	return nil
}