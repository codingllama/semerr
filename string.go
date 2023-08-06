package semerr

// String converts a string into an non-semantic error.
//
// Useful for creating semantic errors with custom messages.
type String string

// Error returns the error string.
func (s String) Error() string {
	return string(s)
}
