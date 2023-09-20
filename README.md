# Go-Dig
go-dig is an implementation of Ruby's dig methods in GoLang. To handle data such as json files you need to create a struct which can get very complex. go-dig allows you to dig thorugh an object using a set of keys and will return the data or object in its path. It will return nil and an error if the value is not found.


## Installation
To install run
```
go get github.com/sauravc99/go-dig
```


Include this import statement in your code 
```
import "github.com/sauravc99/go-dig"
```


## Usage
It's easy to get started with go-dig. There are two ways of using it.

### Data loaded from a file
If the data you want to dig through exists in file:

Open the file, read the data, unmarshal it into an interface
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
After running, 'result' will have the data that exists at the end of the path and 'err' will be nil
