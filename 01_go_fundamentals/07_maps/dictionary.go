package _7_maps

import (
	"errors"
)

// DictionaryErr represents custom error messages related to dictionary operations.
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Predefined error constants for dictionary operations.
const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

// Dictionary is a map that holds word definitions.
type Dictionary map[string]string

// Search looks up a word in the dictionary and returns its definition.
// Returns ErrNotFound if the word is not found.
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add inserts a new word and its definition into the dictionary.
// Returns ErrWordExists if the word already exists.
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch {
	case errors.Is(err, ErrNotFound):
		d[word] = definition
	case err == nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

// Update modifies the definition of an existing word in the dictionary.
// Returns ErrWordDoesNotExist if the word does not exist.
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch {
	case errors.Is(err, ErrNotFound):
		return ErrWordDoesNotExist
	case err == nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

// Delete removes a word and its definition from the dictionary.
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
