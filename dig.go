package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func Dig(h interface{}, keys ...interface{}) (interface{}, error) {
	n := len(keys)

	fmt.Println("num of keys:", n)
	fmt.Println(h)
	//fmt.Println(h.(map[string]interface{}))

	if n == 0 {
		return nil, fmt.Errorf("key is missing")
	}

	for pos, key := range keys {
		fmt.Println(pos, key)

		//the key.(string) is type assertion. It gives access to an interface value's underlying value
		//assigns the underlying string value to the variable keyString. It does not convert
		//using it with ok tests whether the interface value holds that type
		keyString, ok := key.(string) //ok will be true if the key is a string
		fmt.Println("  ", keyString, ok, "string branch")
		if ok {
			//the variable h is type interface{} but we need type map[string]interface{}
			//to access index. interface{} is non indexable
			//the h.(map[string]interface{}) is type assertion. checks and assigns the underlying
			//map[string]interface{} of the h interface{} to variable raw
			raw, ok := h.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("%v is not a string accessable map. Key '%v' at position '%v' not applicable. Map is type: %T", h, keyString, pos+1, h)
			}

			h, ok = raw[keyString]
			if !ok {
				return nil, fmt.Errorf("key '%v' at position '%v' not found in  %v", keyString, pos+1, raw)
			}

			fmt.Println(h)
			fmt.Printf("---%T---\n", h) //prints the var type

			if pos == n-1 {
				return h, nil
			}
			continue
		}

		return nil, fmt.Errorf("key is not supported: '%v' type:%T", key, key)
		//fmt.Println(h.(map[string]interface{}))
	}

	return nil, nil

}

func main() {
	file, err := os.Open("sample copy 2.json")
	if err != nil {
		fmt.Println(err)
	}
	rawBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	var b interface{}
	err = json.Unmarshal(rawBytes, &b)
	if err != nil {
		fmt.Println(err)
	}

	//c, err := Dig(b, "menu", "apple", "banna")
	//c, err := Dig(b, "user", "education", "university", "name")

	//arr := [4]string{"one", "two", "three", "four"}
	//c, err := Dig(b, "menu", "items", arr, 2, "id")

	c, err := Dig(b, "menu", "items", 2, "id")
	//c, err := Dig(b, "menu", "items", 2.0, "id") //not supported
	//c, err := Dig(b, "menu", 2, "id") //not a slice
	//c, err := Dig(b, "menu", "items", "id") //not a string accessable map
	//c, err := Dig(b, "menu", "apple") //key not found in map
	//c, err := Dig(b, "menu", "items", 6) //index out of range
	//c, err := Dig(b)                              //key is missing

	//c, err := Dig(b, "more", 0, 0, "batters", "batter", 2, "type")

	fmt.Println(c)
	fmt.Println(err)
}
