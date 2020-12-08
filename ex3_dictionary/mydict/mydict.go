package mydict

import "errors"

type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errCantUpdate = errors.New("Can't update non-existing word")
var errWordExists = errors.New("Already Exists")

// Search 
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	} 
	return "", errNotFound
}

// Add 
func (d Dictionary) Add(word,def string) error {
	_, err := d.Search(word)
	switch err {
		case errNotFound:
			d[word] = def
		case nil:
			return errWordExists
	}
	return nil
}

// Update
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
		case nil:
			d[word] = def
		case errNotFound:
			return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) {

	delete(d, word)
}

