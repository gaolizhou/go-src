package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	/*
	fmt.Println("Hello World!")
	fmt.Println(DecodeInteger([]byte(`i2097152e`), 0))
	fmt.Println(DecodeString([]byte(`13:announce-list`), 0))
	list_str := []byte(`ll30:http://henbt.com:2710/announceel38:udp://tracker.publicbt.com:80/announceel44:udp://tracker.openbittorrent.com:80/announceel35:udp://tracker.istole.it:80/announceel38:http://tracker.trackerfix.com/announceel31:udp://9.rarbg.com:2710/announceel29:udp://12.rarbg.me:80/announceel29:udp://10.rarbg.me:80/announceel29:udp://11.rarbg.me:80/announceel30:udp://9.rarbg.me:2710/announceee`)
	fmt.Println(DecodeList(list_str, 0))
	//fmt.Println(DecodeList([]byte(`l30:http://henbt.com:2710/announcee`), 0))
	*/
	//fmt.Println(DecodeDict([]byte(`d8:announce39:http://torrent.ubuntu.com:6969/announcee`), 0))

	filename := "ubuntu.torrent"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Error Opening ", filename)
	}
	item, _ := GetDecoder(data[0])(data, 0)
	fmt.Println(item.item_type)
	fmt.Println(item.dict_value["comment"])
	fmt.Println(item.dict_value["creation date"])
	fmt.Println(item.dict_value["announce"])
}

