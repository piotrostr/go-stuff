package dictionary

import (
	"errors"
)

type Dict map[string]string

var NoSuchWordError = errors.New("no such word")

func Search(dict map[string]string, word string) (string, error) {
	value, success := dict[word]
	if !success {
		return "", NoSuchWordError
	}
	return value, nil
}
