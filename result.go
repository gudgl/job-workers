package jw

import (
	"fmt"
	"strings"
)

type Keys map[string]interface{}

func WithError(err error) *Result {
	return NewResult().WithError(err)
}

func WithKeys(keys map[string]interface{}) *Result {
	return NewResult().WithKeys(keys)
}

func WithKey(key string, value interface{}) *Result {
	return NewResult().WithKey(key, value)
}

type Result struct {
	err  error
	keys Keys
}

func NewResult() *Result {
	return &Result{
		keys: make(map[string]interface{}),
	}
}

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

func (e *Result) WithError(err error) *Result {
	e.err = err
	return e
}

func (e *Result) WithKeys(keys map[string]interface{}) *Result {
	e.keys = keys
	return e
}

func (e *Result) WithKey(key string, value interface{}) *Result {
	e.setKey(key, value)
	return e
}

func (e *Result) setKey(key string, value interface{}) {
	e.keys[key] = value
}

func (e Result) GetValue(key string) interface{} {
	return e.keys[key]
}

func (e Result) GetError() error {
	return e.err
}
