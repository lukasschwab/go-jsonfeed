package jsonfeed

import "fmt"

// A MissingRequiredValueError indicates that a required value in some JSON
// feed object is undefined. STRUCTURE denotes the kind of object (for example,
// Item or Hub). KEY is the JSON key for the missing field.
type MissingRequiredValueError struct {
	Structure string
	Key       string
}

func (e MissingRequiredValueError) Error() string {
	return fmt.Sprintf(
		"Required field '%s' missing from '%s'",
		e.Key,
		e.Structure,
	)
}

// IndexedMissingRequiredValueError is a MissingRequiredValueError for an item
// at position INDEX in some iterable within the field (e.g. an array of Items
// or an array of Hubs).
type IndexedMissingRequiredValueError struct {
	error
	Index int
}

func (e IndexedMissingRequiredValueError) Error() string {
	return fmt.Sprintf(
		"%s at index %d",
		e.error.Error(),
		e.Index,
	)
}
