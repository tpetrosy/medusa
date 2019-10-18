package helpers

import (
	"fmt"
	"io/ioutil"
)

//LoadFile load html template files
func LoadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(bytes), nil
}
