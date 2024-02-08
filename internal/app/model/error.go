package model

import "sync"

// ValidateError ...
type ValidateError struct {
	Attribute   string
	Description string
}

// ErrorList ...
type ErrorList struct {
	List []ValidateError
	mu   sync.Mutex
}

// NewErrorList ...
func NewErrorList() *ErrorList {
	return &ErrorList{mu: sync.Mutex{}}
}

// Add ...
func (el *ErrorList) Add(attr, desc string) {
	el.mu.Lock()
	defer el.mu.Unlock()

	el.List = append(el.List, ValidateError{Attribute: attr, Description: desc})
}

// AddList ...
func (el *ErrorList) AddList(l *ErrorList) {
	el.mu.Lock()
	defer el.mu.Unlock()

	el.List = append(el.List, l.List...)
}
