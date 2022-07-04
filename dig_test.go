package dig

import (
	"encoding/json"
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
	"fmt"
	"ioutil"
	"Dig"
)

func loadJson() interface{} {
	j, err := os.Open("example.json")
	if err != nil {
		fmt.Println(err)
	}

	bytes, _ := ioutil.ReadAll(j)

	var a interface{}
	err = json.Unmarshal([]byte(bytes), &a)
	if err != nil {
		fmt.Println(err)
	}
	return a
}

func TestDig(t *testing.T) {
	a := loadJson()

	//test strings
	result, err := Dig(a, "apple")
	assert.Equal(t, "pear", result)
	assert.Nil(t, err)

	result, err = Dig(a, "menu", "header")
	assert.Equal(t, "SVG Viewer", result)
	assert.Nil(t, err)

	//test strings and array access
	result, err = Dig(a, "menu", "items", 0, "id")
	assert.Equal(t, "Open1243636", result)
	assert.Nil(err)

	//nested testing
	result, err = Dig(a, "menu", "items", 3, "servlet-name")
	assert.Equal(t, "cofaxEmail", result)
	assert.Nil(err)

	result, err = Dig(a, "menu", "items", 3, "init-param", "mailHostOverride")
	assert.Equal(t, "mail2", result)
	assert.Nil(err)

	result, err = Dig(a, "more", 0, 0, "batters", "batter", 1, "type")
	assert.Equal(t, "Chocolate", result)
	assert.Nil(err)

	//test array in param
	b := []string{"menu", "header"}
	result, err = Dig(a, b)
	assert.Equal(t, "SVG Viewer", result)
	assert.Nil(t, err)

	//test non existant element
	result, err = Dig(a, "banana")
	assert.Nil(t, result)
	assert.NotNil(t, err)
}