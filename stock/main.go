package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Error struct  {
	err string;
};

func (e *Error) Error() string {
	return e.err;
}

func MakeQuery(stock_id string) (string, error) {
	if len(stock_id) != 6 {
		return "", &Error{"Incorrect stock number = " + stock_id};
	}

	url := "http://hq.sinajs.cn/list=s_";
	if stock_id[0] == '0' || stock_id[0] == '3' || stock_id[0] == '1' {
		url += "sz" + stock_id;
	} else if stock_id[0] == '6' || stock_id[0] == '5' {
		url += "sh" + stock_id;
	}
	resp, err := http.Get(url);
	if err != nil {
		return "", err;
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err;
	}
	data := string(body);
	s := strings.Split(data, ",")
	return s[1]+","+s[3], nil;
}

func main()  {
	args := os.Args[1:]

	for _, ele := range args {
		ret, err := MakeQuery(ele);
		if err != nil {
			fmt.Println(err);
			continue;
		}
		fmt.Println(ret)
	}
	
}