package dig

import (
	"encoding/json"
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

// TestDigStringReturn calls dig.Dig with a json document loaded as a interface{} and keys
// which are the path to a string, checking for the correct return value and no error.
func TestDigStringReturn(t *testing.T) {
	json := returnInterface()

	got, err := Dig(json, "apple")
	want := "pear"

	if got != want || err != nil {
		t.Errorf("got %v, %v - wanted %v, %v", got, err, want, nil)
	}
}

// TestDigFloatReturn calls dig.Dig with a json document loaded as a interface{} and keys which
// are the path to a floating point number, checking for the correct return value and no error.
func TestDigFloatReturn(t *testing.T) {
	json := returnInterface()

	got, err := Dig(json, "three")
	want := 3.3

	if got != want || err != nil {
		t.Errorf("got %v, %v - wanted %v, %v", got, err, want, nil)
	}
}