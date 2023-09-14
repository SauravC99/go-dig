package dig

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func Dig(hash interface{}, keys ...interface{}) (interface{}, error) {
	n := len(keys)

	for pos, key := range keys {
		//the key.(string) is type assertion. It gives access to an interface value's underlying value
		//assigns the underlying string value to the variable keyString. It does not convert
		//using it with ok tests whether the interface value holds that type
		keyString, ok := key.(string) //ok will be true if the key is a string
		if ok {
			//the variable h is type interface{} but we need type map[string]interface{}
			//to access index. interface{} is non indexable
			//the h.(map[string]interface{}) is type assertion. checks and assigns the underlying
			//map[string]interface{} of the h interface{} to variable raw
			inside, ok := hash.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("%v is not a string accessable map. Key '%v' at position '%v' not applicable. Map is type: %T", hash, keyString, pos+1, hash)
			}

			hash, ok = inside[keyString]
			if !ok {
				return nil, fmt.Errorf("key '%v' at position '%v' not found in  %v", keyString, pos+1, inside)
			}

			if pos == n-1 {
				return hash, nil
			}
			continue
		}
		//type assertion, using with ok tests if that type is
		keyInt, ok := key.(int) //ok will be true if the key is a int
		if ok {
			//the variable h is type interface{} but we need  type []interface{}
			//to access index. We cannot index type interface{}
			//type assertion to get the underlying type []interface{} from h interface{}
			inside, ok := hash.([]interface{})
			if !ok {
				return nil, fmt.Errorf("%v is not a slice/int accessable map. Key '%v' at position '%v' not applicable. Map is type: %T", hash, keyInt, pos+1, hash)
			}

			if keyInt < 0 || keyInt >= len(inside) {
				return nil, fmt.Errorf("index '%v' at position '%v' out of range  %v has length '%v'", keyInt, pos+1, inside, len(inside))
			}

			//assign the result back to the hash
			hash = inside[keyInt]

			if pos == n-1 {
				return hash, nil
			}
			continue
		}

		return nil, fmt.Errorf("key is not supported: '%v' type:%T", key, key)
	}

	return nil, fmt.Errorf("key is missing")

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
	//c, err := Dig(b, "menu", "items", 2.0, "id") //not supported
	//c, err := Dig(b, "menu", 2, "id") //not a slice
	//c, err := Dig(b, "menu", "items", "id") //not a string accessable map
	//c, err := Dig(b, "menu", "apple") //key not found in map
	//c, err := Dig(b, "menu", "items", 6) //index out of range
	//c, err := Dig(b)                              //key is missing

	c, err := Dig(b, "more", 0, 0, "batters", "batter", 2, "type")

	//arr := []string{"one", "two", "three", "four"}
	//arr := []string{"menu", "items"}
	//arr := []string{"batters", "batter"}

	//arr := [4]int{1, 2, 3, 4}

	//c, err := Dig(b, "menu", "items", arr, 2, "id")
	//c, err := Dig(b, "menu", "items")
	//c, err := Dig(b, arr)
	//c, err := Dig(b, "more", 0, 0, arr, 2, "type")

	//c, err := Dig(b, "user", "education", "university", "name")
	//arr := []string{"user", "education"}
	//c, err := Dig(b, arr, "university", "name") //NOT WORK
	//arr := []string{"education", "university"}
	//c, err := Dig(b, "user", arr, "name")
	//arr := []string{"university", "name"}
	//c, err := Dig(b, "user", "education", arr)

	fmt.Println()
	fmt.Println(c)
	fmt.Println(err)
}
