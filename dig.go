package main

import (
	"encoding/json"
	"fmt"
)

func Dig(h interface{}, keys ...interface{}) (interface{}, error) {
	n := len(keys)
	fmt.Println(n)
	for e, key := range keys {
		fmt.Println(e, key)
	}
	return nil, fmt.Errorf("key is missing")
}

func main() {
	a := []byte(`{
		"menu":{
			"one":"oneeee",
			"item":1
		}
	}`)
	var b interface{}
	err := json.Unmarshal(a, &b)
	if err != nil {
		fmt.Println(err)
	}

	c, err := Dig(b, "menu", "apple", "banna")
	fmt.Println(c)
}
