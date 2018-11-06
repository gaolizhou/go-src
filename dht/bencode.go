package main

type ItemType int

const (
	kString    ItemType = 0
	kInt       	ItemType = 1
	kList      	ItemType = 2
	kDict      	ItemType = 3
)

type BItem struct {
	item_type ItemType
	int_value int
	string_value string
	list_value []BItem
	dict_value map[string]BItem
}
/*
func (this BItem) String() string {
	switch this.item_type {
	case kString:
		return this.string_value;
	case kInt:
		return string(this.int_value);
	case kList:
		return this.list_value;
	case kDict:
		return string(this.dict_value);
	} 
}
*/
type DecodeFunc func(data []byte, index int) (BItem, int)

func GetDecoder(bt byte)(DecodeFunc) {
	fun_map := map[byte]DecodeFunc {
		'd': DecodeDict,
		'l': DecodeList,
		'i': DecodeInteger,
		'0': DecodeString,
		'1': DecodeString,
		'2': DecodeString,
		'3': DecodeString,
		'4': DecodeString,
		'5': DecodeString,
		'6': DecodeString,
		'7': DecodeString,
		'8': DecodeString,
		'9': DecodeString,
	}
	return fun_map[bt]
}

//i2097152e
func DecodeInteger(data []byte, index int) (BItem, int) {
	item := BItem{
		item_type:kInt,
	}
	i := index + 1
	int_value := 0
	for ; data[i] != 'e';i++ {
		int_value = int_value * 10
		int_value += int(data[i])-int('0')
	}
	item.int_value = int_value
	return item, i - index + 1
}

//13:announce-list
func DecodeString(data []byte, index int) (BItem, int) {
	item := BItem {
		item_type:kString,
	}
	i := index
	length := 0
	for ; data[i] != ':';i++ {
		length = length * 10
		length += int(data[i])-int('0')
	}
	item.string_value = string(data[i+1:length+i+1])
	return item, length +i+1 - index
}
//ll30:http://henbt.com:2710/announceel38:udp://tracker.publicbt.com:80/announceel44:udp://tracker.openbittorrent.com:80/announceel35:udp://tracker.istole.it:80/announceel38:http://tracker.trackerfix.com/announceel31:udp://9.rarbg.com:2710/announceel29:udp://12.rarbg.me:80/announceel29:udp://10.rarbg.me:80/announceel29:udp://11.rarbg.me:80/announceel30:udp://9.rarbg.me:2710/announceee
func DecodeList(data []byte, index int) (BItem, int) {
	item := BItem {
		item_type:kList,
		list_value:make([]BItem, 0, 0),
	}
	step := 0
	i:= index + 1
	for ; data[i] != 'e'; i += step {
		var itm BItem
		itm, step = GetDecoder(data[i])(data, i)
		item.list_value = append(item.list_value, itm)
	}
	return item,i + 1 - index
}

func DecodeDict(data []byte, index int) (BItem, int) {
	item := BItem {
		item_type:kDict,
		dict_value: make(map[string]BItem),
	}
	step := 0
	i:= index + 1
	for ; data[i] != 'e'; i+=step {
		var key, value BItem
		key, step = GetDecoder(data[i])(data, i)
		i = i + step
		value, step = GetDecoder(data[i])(data, i)
		item.dict_value[key.string_value] = value
	}
	return item,i + 1 - index
}