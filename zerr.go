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

// WithField creates a new Error instance, with one or more fields added.
// Note that this is equivalent to calling WrapNoStack(err, f)
func (e *Error) WithField(f zap.Field, additionalFields ...zap.Field) *Error {
	newErr := &Error{
		err:      e,
		fields:   append(additionalFields, f),
		hasStack: e.hasStack,
	}

	return newErr
}

// LogDebug logs an Error with Debug level to a given zap logger
func (e *Error) LogDebug(logger *zap.Logger) {
	logger.Debug(e.Error(), e.Fields()...)
}

// LogInfo logs an Error with Info level to a given zap logger
func (e *Error) LogInfo(logger *zap.Logger) {
	logger.Info(e.Error(), e.Fields()...)
}

// LogWarn logs an Error with Warn level to a given zap logger
func (e *Error) LogWarn(logger *zap.Logger) {
	logger.Warn(e.Error(), e.Fields()...)
}

// LogError logs an Error with Error level to a given zap logger
func (e *Error) LogError(logger *zap.Logger) {
	logger.Error(e.Error(), e.Fields()...)
}

// LogDPanic logs an Error with DPanic level to a given zap logger
func (e *Error) LogDPanic(logger *zap.Logger) {
	logger.DPanic(e.Error(), e.Fields()...)
}

// LogPanic logs an Error with Panic level to a given zap logger
func (e *Error) LogPanic(logger *zap.Logger) {
	logger.Panic(e.Error(), e.Fields()...)
}

// LogFatal logs an Error with Fatal level to a given zap logger
func (e *Error) LogFatal(logger *zap.Logger) {
	logger.Fatal(e.Error(), e.Fields()...)
}

// Wrap adds zap fields to an error
func Wrap(err error, fields ...zap.Field) *Error {
	return wrapWithStack(1, err, fields...)
}

// wrapWithStack wraps the error and attaches a stacktrace, skipping the first 'lvl' levels.
// This function allows us to remove any 'zerr' calls from the stacktraces, and instead
// list the stacktrace of the original caller
func wrapWithStack(lvl int, err error, fields ...zap.Field) *Error {
	// If we're not adding any fields, and the supplied error is already of the correct type,
	// return it directly
	if e, ok := err.(*Error); ok && len(fields) == 0 {
		return e
	}

	// Check if we've already wrapped with a stack.
	// If that's the case, we won't add another stacktrace
	var e *Error
	hasStack := false
	if errors.As(err, &e) {
		hasStack = e.hasStack
	}

	if !hasStack {
		fields = append(fields, zap.StackSkip("stacktrace", lvl+1))
	}

	return &Error{
		err:      err,
		fields:   fields,
		hasStack: true,
	}
}

// WrapNoStack wraps error with fields, but always excludes the stack trace
func WrapNoStack(err error, fields ...zap.Field) *Error {
	// Tag that this error has a stacktrace
	return &Error{
		err:    err,
		fields: fields,
	}
}

// Sugar is a sugared version of the 'Wrap'  function above.
// It allows the user to add fields without using strongly typed fields, e.g.
// errors.Wraps(err, "key-1", 12, "key-2", "some string", "key-3", value)
func Sugar(err error, args ...interface{}) *Error {
	fields := sugarFields(args...)
	return wrapWithStack(1, err, fields...)
}

// SugarNoStack is exactly like the 'Sugar' function but without an additional stacktrace
func SugarNoStack(err error, args ...interface{}) *Error {
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
