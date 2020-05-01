package error

import (
	"fmt"
)

const (
	noElementErrorMessage = "Element with id %s was not found"
	cantListErrorMessage = "Cannot get element list: %s"
)

// An shorthand error to say that the desired element is not existing
type NoSuchElement struct {
	elementId string
}

// Create a new NoSuchElement error, with the associated id
func NewNoSuchElement(elementId string) *NoSuchElement {
	return &NoSuchElement{elementId: elementId}
}

// Prints the error. See noElementErrorMessage to get the output of the
// error itself
func (e *NoSuchElement) Error() string {
	return fmt.Sprintf(noElementErrorMessage, e.elementId)
}

type CantGetListing struct {
	reason string
}

func NewCantGetListing(reason string) *CantGetListing {
	return &CantGetListing{reason: reason}
}

func (e *CantGetListing) Error() string {
	return fmt.Sprintf(cantListErrorMessage, e.reason)
}
