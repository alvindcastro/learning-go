package maps

//Dictionary wrapper
type Dictionary map[string]string

// Search using Dictionary
func (d Dictionary) Search(word string) string {
	return d[word]
}

// Search - initial function
func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}
