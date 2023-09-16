package dig

import (
	"encoding/json"
	"fmt"
	"testing"
)

func returnInterface() interface{} {
	b := []byte(`{
		"apple": "pear",
		"two": 2,
		"three": 3.3,
		"menu": {
		  "header": "SVG Viewer",
		  "items": [
			{
			  "id": "Open1243636"
			},
			{
			  "id": "OpenNew",
			  "label": "Open New",
			  "desc" : "Opens new thing"
			}
		  ]
		},
		"more": [
		  [
			{
			  "id": "0001",
			  "type": "donut",
			  "ppu": 0.55,
			  "batters": {
				"batter": [
				  {
					"id": "1001",
					"type": "Regular"
				  },
				  {
					"id": "1003",
					"type": "Blueberry"
				  }
				]
			  }
			},
			{
			  "id": "0002",
			  "type": "donut",
			  "topping": [
				{
				  "id": "5001",
				  "type": "None"
				},
				{
				  "id": "5002",
				  "type": "Glazed"
				}
			  ]
			}
		  ]
		]
	  }
	`)

	var a interface{}
	json.Unmarshal(b, &a)
	return a
}

// TestDig calls dig.Dig with a json document loaded as a interface{} and keys which
// are the path to the desired value, checking for the correct return value and no error.
func TestDig(t *testing.T) {
	json := returnInterface()

	got, err := Dig(json, "menu", "items", 1, "label")
	want := "Open New"

	if got != want || err != nil {
		t.Errorf("got %v, %v - wanted %v, %v", got, err, want, nil)
	}
}

// TestDigStringReturn calls dig.Dig with a json document loaded as a interface{} and a key
// which is the path to a string, checking for the correct return value and no error.
func TestDigStringReturn(t *testing.T) {
	json := returnInterface()

	got, err := Dig(json, "apple")
	want := "pear"

	if got != want || err != nil {
		t.Errorf("got %v, %v - wanted %v, %v", got, err, want, nil)
	}
}

// TestDigFloatReturn calls dig.Dig with a json document loaded as a interface{} and a key which
// is the path to a floating point number, checking for the correct return value and no error.
func TestDigFloatReturn(t *testing.T) {
	json := returnInterface()

	got, err := Dig(json, "three")
	want := 3.3

	if got != want || err != nil {
		t.Errorf("got %v, %v - wanted %v, %v", got, err, want, nil)
	}
}

// TestDigErrorEmpty calls dig.Dig with a json document loaded as a
// interface{} and empty keys, checking for an error.
func TestDigErrorEmptyKey(t *testing.T) {
	json := returnInterface()

	got, err := Dig(json)

	if got != nil || err == nil {
		t.Errorf("got %v, %v - wanted %v, %v", got, err, nil, "key is missing")
	}
}

// TestDigErrorKeyNotFound calls dig.Dig with a json document loaded as a interface{}
// and keys that lead to a value that does not exist, checking for an error.
func TestDigErrorKeyNotFound(t *testing.T) {
	json := returnInterface()

	got, err := Dig(json, "menu", "apples")

	if got != nil || err == nil {
		t.Errorf("got %v, %v - wanted %v, %v", got, err, nil, "key not found")
	}
}

// TestDigErrorUnsupportedKey calls dig.Dig with a json document loaded as a interface{}
// and a float 1.0 as a key parameter which is not supported, checking for an error.
func TestDigErrorUnsupportedKey(t *testing.T) {
	json := returnInterface()

	got, err := Dig(json, "menu", "items", 1.0, "id")

	if got != nil || err == nil {
		t.Errorf("got %v, %v - wanted %v, %v", got, err, nil, "key is not supported")
	}
}

// TestDigErrorNotSlice calls dig.Dig with a json document loaded as a interface{} and
// keys that try to access a string map with an int 2 parameter, checking for an error.
func TestDigErrorNotSlice(t *testing.T) {
	json := returnInterface()

	got, err := Dig(json, "menu", 2, "id")

	if got != nil || err == nil {
		t.Errorf("got %v, %v - wanted %v, %v", got, err, nil, "not a int accessable map")
	}
}

// TestDigErrorNoStringAccess calls dig.Dig with a json document loaded as a interface{} and
// keys that try to access a int map/array with an string 'id' parameter, checking for an error.
func TestDigErrorNoStringAccess(t *testing.T) {
	json := returnInterface()

	got, err := Dig(json, "menu", "items", "id")

	if got != nil || err == nil {
		t.Errorf("got %v, %v - wanted %v, %v", got, err, nil, "not a string accessable map")
	}
}

// TestDigErrorIndexOutOfRange calls dig.Dig with a json document loaded as a interface{}
// and keys that try to access a int map element that is out of range, checking for an error.
func TestDigErrorIndexOutOfRange(t *testing.T) {
	json := returnInterface()

	got, err := Dig(json, "menu", "items", 6)

	if got != nil || err == nil {
		t.Errorf("got %v, %v - wanted %v, %v", got, err, nil, "index out of range")
	}
}

func BenchmarkDigShallow(b *testing.B) {
	json := returnInterface()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Dig(json, "apple")
	}
}

func BenchmarkDigMedium(b *testing.B) {
	json := returnInterface()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Dig(json, "menu", "items", 1, "id")
	}
}

func BenchmarkDigDeep(b *testing.B) {
	json := returnInterface()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Dig(json, "more", 0, 0, "batters", "batter", 1, "type")
	}
}

func ExampleDig() {
	json := returnInterface()

	result, err := Dig(json, "menu", "header")
	result1, err1 := Dig(json, "more", 0, 0, "type")
	fmt.Println(result)
	fmt.Println(err)
	fmt.Println()
	fmt.Println(result1)
	fmt.Println(err1)
	// Output:
	// SVG Viewer
	// <nil>
	//
	// donut
	// <nil>
}
