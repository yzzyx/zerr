package zerr

import (
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
	if cause, ok := e.err.(*Error); ok {
		// Include fields from causing errors, if any
		return append(cause.Fields(), e.fields...)
	}
	return e.fields
}

// Cause returns the cause of this error
func (e *Error) Cause() error {
	return e.err
}

// Wrap adds zap fields to an error
func Wrap(err error, fields ...zap.Field) error {
	return &Error{
		err:    err,
		fields: fields,
	}
}

// WrapStack wraps error with fields and a stacktrace
func WrapStack(err error, fields ...zap.Field) error {
	// Check if we've already wrapped with a stack.
	// If that's the case, we won't add another stacktrace
	if e, ok := err.(*Error); ok {
		if e.hasStack {
			return Wrap(err, fields...)
		}
	}

	fields = append(fields, zap.Stack("stacktrace"))

	// Tag that this error has a stacktrace
	e := Wrap(err, fields...).(*Error)
	e.hasStack = true
	return e
}

// Sugar is a sugared version of the 'Wrap'  function above.
// It allows the user to add fields without using strongly typed fields, e.g.
// errors.Wraps(err, "key-1", 12, "key-2", "some string", "key-3", value)
func Sugar(err error, args ...interface{}) error {
	fields := sugarFields(args...)
	return Wrap(err, fields...)
}

// SugarStack is exactly like the 'Sugar' function with an additional stacktrace
func SugarStack(err error, args ...interface{}) error {
	fields := sugarFields(args...)
	return WrapStack(err, fields...)
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
