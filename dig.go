package dig

import "fmt"

func Dig(hash interface{}, keys ...interface{}) (interface{}, error) {
	n := len(keys)

	for pos, key := range keys {
		//the key.(string) is type assertion. It gives access to an interface value's underlying value
		//assigns the underlying string value to the variable keyString. It does not convert
		//using it with ok tests whether the interface value holds that type
		keyString, ok := key.(string)
		if ok {
			//the variable h is type interface{} but we need type map[string]interface{}
			//to access index. interface{} is non indexable
			//the h.(map[string]interface{}) is type assertion. checks and assigns the underlying
			//map[string]interface{} of the h interface{} to variable inside
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
		keyInt, ok := key.(int)
		if ok {
			//the variable h is type interface{} but we need  type []interface{}
			//to access index. We cannot index type interface{}
			//type assertion to get the underlying type []interface{} from hash interface{}
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
