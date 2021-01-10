package jw

import (
	"fmt"
	"strings"
)

// Keys represents type used for keys in the Result
type Keys map[string]interface{}

// WithError creates new Result with given error
func WithError(err error) *Result {
	return NewResult().WithError(err)
}

// WithKeys creates new Result with given Keys
func WithKeys(keys Keys) *Result {
	return NewResult().WithKeys(keys)
}

// WithKey creates new Result with given value for given key
func WithKey(key string, value interface{}) *Result {
	return NewResult().WithKey(key, value)
}

// Result represents the response from Job Execute method
type Result struct {
	err  error
	keys Keys
}

// NewResult creates empty Result
func NewResult() *Result {
	return &Result{
		keys: make(map[string]interface{}),
	}
}

// String returns the string build in specific format
func (e *Result) String() string {
	var sb strings.Builder

	sb.WriteString("jw error with keys: ")

	for k, v := range e.keys {
		sb.WriteString(fmt.Sprintf("%s: %+v; ", k, v))
	}

	if e.err != nil {
		sb.WriteString("and error: ")
		sb.WriteString(e.err.Error())
	}

	return sb.String()
}

// WithError sets given error to Result
func (e *Result) WithError(err error) *Result {
	e.err = err
	return e
}

// WithKeys sets given Keys to Result
func (e *Result) WithKeys(keys Keys) *Result {
	e.keys = keys
	return e
}

// WithKey sets given value for given key to Result
func (e *Result) WithKey(key string, value interface{}) *Result {
	e.setKey(key, value)
	return e
}

// setKey sets given value for given eky
func (e *Result) setKey(key string, value interface{}) {
	e.keys[key] = value
}

// GetValue returns the value for given key
func (e Result) GetValue(key string) interface{} {
	return e.keys[key]
}

// GetError returns the error
func (e Result) GetError() error {
	return e.err
}
