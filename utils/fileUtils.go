package utils

import (
	"errors"
	"io/ioutil"
)

func ReadFile(filepath string) (string, error) {

	if isEmpty(filepath) {
		return "", errors.New("Bos birakma")
	}

	bytes, err := ioutil.ReadFile(filepath)
	checkError(err)

	return string(bytes), nil

}
