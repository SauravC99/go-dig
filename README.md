# Go-Dig
[![Go Reference](https://pkg.go.dev/badge/github.com/sauravc99/go-dig.svg)](https://pkg.go.dev/github.com/sauravc99/go-dig)

go-dig is an implementation of Ruby's dig methods in GoLang. To handle data such as json files you need to create a struct which can get very complex. go-dig allows you to dig thorugh an object using a set of keys and will return the data or object in its path. It will return nil and an error if the value is not found.


## Installation
To install run
```
go get github.com/sauravc99/go-dig
```

Include this import statement in your code 
```go
import "github.com/sauravc99/go-dig"
```


## Usage
It's easy to get started with go-dig. Here are two ways of using it.


### Data loaded from a file
If the data you want to dig through exists in a file:

Open the file, read the data, create an interface, and unmarshal the data into the interface.
```go
file, err := os.Open("path/to/data.json")
if err != nil {
	fmt.Println(err)
}
rawBytes, err := io.ReadAll(file)
if err != nil {
	fmt.Println(err)
}

var jsonData interface{}

err = json.Unmarshal(rawBytes, &jsonData)
if err != nil {
	fmt.Println(err)
}
```

Now you can use the interface as the first parameter and dig to the data you need.
```go
result, err := dig.Dig(jsonData, "path", "to", "data", "you", "want")
if err != nil {
    fmt.Println(err)
}
```

After running, `result` will have the data that exists at the end of the path and `err` will be nil.

### Data in your code 
If the data you want to dig through exists inline in your code:

Create an interface and unmarshal the data into the interface.
```go
data := []byte(`{{"name":"Tiger","id":2454,"information":{"details":{"status":"active","claws":true,"teeth":"many"}}}}`)

var jsonData interface{}

err := json.Unmarshal(data, &jsonData)
if err != nil {
    fmt.Println(err)
}
```

Now you can use the interface as the first parameter and dig to get the data you need.
```go
result, err := dig.Dig(jsonData, "name")
if err != nil {
    fmt.Println(err)
}
// result : Tiger
// err    : nil
```
```go
result, err := dig.Dig(jsonData, "information", "details", "teeth")
if err != nil {
    fmt.Println(err)
}
// result : many
// err    : nil
```


## Examples
Using this json file as our data
```json
{
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
        "desc": "Opens new thing"
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
```
Here's how to dig for different data
```go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/sauravc99/go-dig"
)

func main() {
	file, err := os.Open("sample data.json")
	if err != nil {
		fmt.Println(err)
	}
	rawBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	var jsonData interface{}
	err = json.Unmarshal(rawBytes, &jsonData)
	if err != nil {
		fmt.Println(err)
	}


	result, err := dig.Dig(jsonData, "three")
	if err != nil {
		fmt.Println(err)
	}
	// result : 3.3
	// err    : nil

	result1, err := dig.Dig(jsonData, "menu", "header")
	if err != nil {
		fmt.Println(err)
	}
	// result1 : SVG Viewer
	// err     : nil


	// You can also dig through arrays using a index
	result2, err := dig.Dig(jsonData, "more", 0, 0, "batters", "batter", 1, "type")
	if err != nil {
		fmt.Println(err)
	}
	// result2 : Blueberry
	// err     : nil
}
```


## Errors
If you get a return of nil and the error variable is not nil, there has been an error. go-dig will return a helpful error message to help you diagnose the issue.

The error message will say which key parameter is the problem and it's position. The positions start at 1 after the interface parameter. 
Ex:
```go
result, err := dig.Dig(jsonData, "menu", "items", 1, "id")

//position 1 - menu
//position 2 - items
//position 3 - 1
//position 4 - id
```

An example of a key not found error:
```go
result, err := dig.Dig(jsonData, "menu", "apple")
if err != nil {
    fmt.Println(err)
}
// result : nil
// err    : key 'apple' at position '2' not found in map[...]
```
