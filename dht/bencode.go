package main

type BItem interface {
}

type BList struct {
	item_list []BItem
}

type BDict struct {
	item_dict map[string]BItem
}

