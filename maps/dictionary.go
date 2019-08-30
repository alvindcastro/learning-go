package maps

import (
	"errors"
)

//Dictionary wrapper
type Dictionary map[string]string

//ErrNotFound var
var ErrNotFound = errors.New("could not find the word you were looking for")

// Add function
func (d Dictionary) Add(word, definition string) {
	d[word] = definition 
}

// Search using Dictionary
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Search - initial function
func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}
