package main

import (
	"fmt"
	"encoding/json"
)

func Dig(h interface{}, keys ...interface{}) (interface{}, error) {
	n := len(keys)
	for e, key := range(keys) {
		arrayKey, ok := key.([]string) //check if list of strings was passed in for keys
		if ok {
			n += len(arrayKey) - e - 1 //n increases based on len of array minus ones we have done
			for j, k := range arrayKey {
				raw, ok := h.(map[string]interface{})
				if !ok {
					return nil, fmt.Errorf("%v -> not a map", h)
				}

				h, ok = raw[k]
				if !ok {
					return nil, fmt.Errorf("key %v is not found in %v", k, h)
				}

				if j == n - 1 {
					return h, nil
				}
				continue
			}
		}
		stringKey, ok := key.(string) //string is passed in
		if ok {
			raw, ok := h.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("%v -> not a map", h)
			}

			h, ok = raw[stringKey]
			if !ok {
				return nil, fmt.Errorf("key %v is not found in %v", stringKey, h)
			}

			if e == n - 1 {
				return h, nil
			}
			continue
		}
		intKey, ok := key.(int) //int is passed in
		if ok {
			raw, ok := h.([]interface{})
			if !ok {
				return nil, fmt.Errorf("%v -> not a slice", h)
			}

			if intKey < 0 || intKey >= len(raw) {
				return nil, fmt.Errorf("index out of range %v : %v", intKey, raw)
			}

			h = raw[intKey]
			if e == n - 1 {
				return h, nil
			}
			continue
		}
		return nil, fmt.Errorf("key is not supported: %v", key)
	}
	return nil, fmt.Errorf("Key is missing")
}


func main() {
	j := []byte(`{
		"apple" : "pear",
   	"menu":{
	      "header":"SVG Viewer",
	      "items":[
	         {
	            "id":"Open1243636"
	         },
	         {
	            "id":"OpenNew",
	            "label":"Open New"
	         },
	         {
	            "id":"Quality"
	         },
	         {
	            "servlet-name":"cofaxEmail",
	            "servlet-class":"org.cofax.cds.EmailServlet",
	            "init-param":{
	               "mailHost":"mail1",
	               "mailHostOverride":"mail2"
	            }
	         }
	      ]
   	},
   	"more":[
	      [
	         {
	            "id":"0001",
	            "type":"donut",
	            "name":"Cake",
	            "ppu":0.55,
	            "batters":{
	               "batter":[
	                  {
	                     "id":"1001",
	                     "type":"Regular"
	                  },
	                  {
	                     "id":"1002",
	                     "type":"Chocolate"
	                  },
	                  {
	                     "id":"1003",
	                     "type":"Blueberry"
	                  },
	                  {
	                     "id":"1004",
	                     "type":"Devil's Food"
	                  }
	               ]
	            }
	         }
	      ]
   	]
	}`)

	var a interface{}
	err := json.Unmarshal(j, &a)
	if err != nil {
		fmt.Println(err)
	}

	b, err := Dig(a, "apple")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)

	b, err = Dig(a, "menu", "header")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)

	b, err = Dig(a, "menu", "items", 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)

	b, err = Dig(a, "menu", "items", 0, "id")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)

	b, err = Dig(a, "menu", "items", 3, "servlet-name")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)

	b, err = Dig(a, "menu", "items", 3, "init-param", "mailHostOverride")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)

	b, err = Dig(a, "more", 0, 0, "batters", "batter", 1, "type")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)

	b, err = Dig(a, "more", 0, 0, "name")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)

	//nil
	b, err = Dig(a, "strawberry")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)
}