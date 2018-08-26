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
	item := BItem{
		item_type:kString,
	}
	i := index
	len := 0
	for ; data[i] != ':';i++ {
		len = len * 10
		len += int(data[i])-int('0')
	}
	item.string_value = string(data[i+1:len+i+1])
	return item, len+i+1 - index
}
/*
func Decode(data []byte, index int) (item BItem, bytes_consume int){
	switch data[index] {
	case 'd':

	case 'i':
		return decodeInteger(input);
	case 'l':
		return decodeList(input);
	case '0':
	case '1':
	case '2':
	case '3':
	case '4':
	case '5':
	case '6':
	case '7':
	case '8':
	case '9':
		return decodeString(input);
	}

	for i:=index; ; {
		switch data[i] {
		case 'd':

		case 'i':
			return decodeInteger(input);
		case 'l':
			return decodeList(input);
		case '0':
		case '1':
		case '2':
		case '3':
		case '4':
		case '5':
		case '6':
		case '7':
		case '8':
		case '9':
			return decodeString(input);
		}
	}
}

func (item BItem) 	DecodeFromBytes(data []byte, index int) (len int) {
	for i:=index; ; {
		switch data[i] {
		case 'd':

		case 'i':
			return decodeInteger(input);
		case 'l':
			return decodeList(input);
		case '0':
		case '1':
		case '2':
		case '3':
		case '4':
		case '5':
		case '6':
		case '7':
		case '8':
		case '9':
			return decodeString(input);
		}
	}
}
*/