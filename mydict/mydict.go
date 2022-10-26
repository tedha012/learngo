package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errCantUpdate = errors.New("Cant update non-existing")
	errWordExists = errors.New("Theat word already exist")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the Dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = definition
	case nil:
		return errWordExists
	}
	return nil
}

// Update a word to the Dictionary
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word to the Dictionary
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	if err != nil {
		return errNotFound
	}
	delete(d, word)
	return nil
}
