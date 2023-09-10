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
		fmt.Println("  ", keyString, ok)
		if ok {
			//copy the hash h into raw bc we want to keep h as type interface{} but we need to
			//convert NOPE to type map[string]interface{} to access index. interface{} is non indexable
			//the h.(map[string]interface{}) is type assertion. checks and assigns the underlying
			//map[string]interface{} of the h interface{} to variable raw
			raw, ok := h.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("'%v' is not a string accessable map", h)
			}

			h, ok = raw[keyString]
			if !ok {
				return nil, fmt.Errorf("key '%v' is not found in %v", keyString, raw)
			}

			fmt.Println(h)

			if pos == n-1 {
				return h, nil
			}
			continue
		}

		return nil, fmt.Errorf("key is not supported: '%v'", key)
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

	//c, err := Dig(b, "menu", "items", 2, "id")
	//c, err := Dig(b, "menu", 2, "id")
	c, err := Dig(b, "menu", "items", "id")

	fmt.Println(c, err)
	//d, err := Dig(b)
	//fmt.Println(d, err)
	//arr := [4]string{"one", "two", "three", "four"}
	//e, err := Dig(b, arr)
	//fmt.Println(e, err)
}
