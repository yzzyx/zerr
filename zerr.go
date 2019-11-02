package zerr

import (
	"errors"
	"go.uber.org/zap"
)

// Error is the type used to wrap other errors with additional fields
type Error struct {
	err      error
	fields   []zap.Field
	hasStack bool
}

// Error makes us implement the standard error interface
func (e *Error) Error() string {
	return e.err.Error()
}

// Fields returns all fields attached to this error, and all fields attached to previous errors
func (e *Error) Fields() []zap.Field {
	var ok bool
	fields := e.fields

	err := e
	for {
		err, ok = err.err.(*Error)
		if !ok {
			break
		}
		fields = append(fields, err.fields...)
	}
	return fields
}

// Unwrap returns the cause of this error
func (e *Error) Unwrap() error {
	return e.err
}

// Cause returns the cause of this error
// Note that Unwrap should be used, this is included for
// compatibility with pkg/errors
func (e *Error) Cause() error {
	return e.err
}

// Wrap adds zap fields to an error
func Wrap(err error, fields ...zap.Field) error {
	// Check if we've already wrapped with a stack.
	// If that's the case, we won't add another stacktrace
	var e *Error
	if errors.As(err, &e) {
		if !e.hasStack {
			fields = append(fields, zap.Stack("stacktrace"))
		}
	}

	return &Error{
		err:      err,
		fields:   fields,
		hasStack: true,
	}
}

// WrapNoStack wraps error with fields, but always excludes the stack trace
func WrapNoStack(err error, fields ...zap.Field) error {
	// Tag that this error has a stacktrace
	return &Error{
		err:    err,
		fields: fields,
	}
}

// Sugar is a sugared version of the 'Wrap'  function above.
// It allows the user to add fields without using strongly typed fields, e.g.
// errors.Wraps(err, "key-1", 12, "key-2", "some string", "key-3", value)
func Sugar(err error, args ...interface{}) error {
	fields := sugarFields(args...)
	return Wrap(err, fields...)
}

// SugarNoStack is exactly like the 'Sugar' function but without an additional stacktrace
func SugarNoStack(err error, args ...interface{}) error {
	fields := sugarFields(args...)
	return WrapNoStack(err, fields...)
}

// Fields returns any/all fields that are attached to an error
func Fields(err error) []zap.Field {
	if e, ok := err.(*Error); ok {
		return e.Fields()
	}
	return nil
}

// Cause returns the original cause for an error, if available.
func Cause(err error) error {
	type causer interface {
		Cause() error
	}

	for err != nil {
		e, ok := err.(causer)
		if !ok {
			return err
		}
		err = e.Cause()
	}
	return nil
}

// sugarFields converts arguments to zap fields
func sugarFields(args ...interface{}) []zap.Field {
	fields := make([]zap.Field, 0, len(args))
	for i := 0; i < len(args); {
		// This is a strongly-typed field. Consume it and move on.
		if f, ok := args[i].(zap.Field); ok {
			fields = append(fields, f)
			i++
			continue
		}

		// Make sure this element isn't a dangling key.
		if i == len(args)-1 {
			// Log as dangling key
			fields = append(fields, zap.Any("dangling-key", args[i]))
			break
		}

		// Consume this value and the next, treating them as a key-value pair. If the
		// key isn't a string, add this pair to the slice of invalid pairs.
		key, val := args[i], args[i+1]
		if keyStr, ok := key.(string); !ok {
			invalid := struct {
				Pos   int
				Key   interface{}
				Value interface{}
			}{i, key, val}
			fields = append(fields, zap.Any("invalid", invalid))
		} else {
			fields = append(fields, zap.Any(keyStr, val))
		}
		i += 2
	}
	return fields
}
