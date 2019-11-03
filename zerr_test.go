package zerr

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestWrapNoStack(t *testing.T) {
	err := errors.New("test")

	// When
	// we wrap an error
	err = WrapNoStack(err, zap.Int("intfield", 1))

	// Then
	// it should be of our error type
	_, ok := err.(*Error)
	require.True(t, ok)

	// And
	// one extra field should be available
	fields := Fields(err)
	require.Len(t, fields, 1)

	// When
	// We wrap the error again
	err = WrapNoStack(err, zap.String("stringfield", "abc"))

	// Then
	// Two fields should be returned (int and string)
	fields = Fields(err)
	require.Len(t, fields, 2)
}

func TestWrap(t *testing.T) {
	err := errors.New("test")

	// When
	// we wrap an error
	err = Wrap(err, zap.Int("intfield", 1))

	// Then
	// it should be of our error type
	_, ok := err.(*Error)
	require.True(t, ok)

	// And
	// two field should be available - one int and one stack
	fields := Fields(err)
	require.Len(t, fields, 2)
	for _, field := range fields {
		require.IsType(t, zap.Field{}, field)
	}

	// When
	// we try to add another stack trace by wrapping the same error
	err = Wrap(err)

	// Then
	// the original stacktrace should still be retained
	newFields := Fields(err)
	require.Len(t, newFields, 2)
	for i := range newFields {
		require.True(t, zap.Field.Equals(newFields[i], fields[i]))
	}
}

func TestCause(t *testing.T) {
	// When
	// we wrap an error
	originalError := errors.New("original error")
	err := Wrap(originalError)

	// Then
	// Cause() should return the original error
	e := Cause(err)
	require.Equal(t, e, originalError)

	// When
	// we wrap the error once more
	err = Wrap(err)

	// Then
	// Cause() should still return the original error
	e = Cause(err)
	require.Equal(t, e, originalError)
}

func TestSugarNoStack(t *testing.T) {
	originalError := errors.New("original error")

	// When
	// we wrap an error with SugarNoStack() and non-strongly typed keys
	err := SugarNoStack(originalError, "intvalue", 1)

	// And
	// one field should be available
	fields := Fields(err)
	require.Len(t, fields, 1)
	for _, field := range fields {
		require.IsType(t, zap.Field{}, field)
	}

	// When
	// we wrap an error with Sugar() using strongly typed keys
	err = SugarNoStack(originalError, zap.Int("intvalue", 1), zap.String("stringvalue", " abc"))

	// Then
	// two field should be available
	fields = Fields(err)
	require.Len(t, fields, 2)
	for _, field := range fields {
		require.IsType(t, zap.Field{}, field)
	}

	// When
	// we wrap an error with Sugar() using a mix of strongly typed keys and non-types keys
	err = SugarNoStack(originalError, zap.Int("intvalue", 1), zap.String("stringvalue", " abc"), "value3", "test")

	// Then
	// three field should be available
	fields = Fields(err)
	require.Len(t, fields, 3)
	for _, field := range fields {
		require.IsType(t, zap.Field{}, field)
	}

	// When
	// we wrap an error with a dangling key, it should be added as an "any"-field
	err = SugarNoStack(originalError, "dangling key")

	// Then
	// we should have one field (with key "dangling-key")
	fields = Fields(err)
	require.Len(t, fields, 1)
}

func TestAddFields(t *testing.T) {
	// When
	// we start with a plain error
	originalError := errors.New("original error")

	// Then
	// we get a wrapper error in return
	e := Wrap(originalError)
	require.NotNil(t, e)
	require.IsType(t, &Error{}, e)

	// When
	// we add fields to the error
	e2 := e.WithField(zap.Int("test", 1))

	// Then
	// A new instance is returned
	require.NotNil(t, e)
	require.IsType(t, &Error{}, e2)
	require.NotEqual(t, e, e2)
	require.Equal(t, e2.err, e)

	// When
	// logging errors in level Debug - Error
	// Then
	// no panics occur
	logger := zap.NewNop()
	e2.LogDebug(logger)
	e2.LogInfo(logger)
	e2.LogWarn(logger)
	e2.LogError(logger)

	checkPanic := func(fn func()) (hasPanic bool) {
		defer func() {
			if x := recover(); x != nil {
				hasPanic = true
			}
		}()
		fn()
		return hasPanic
	}

	// When
	// Logging error to DPanic level
	// Then
	// No panic occur
	require.False(t, checkPanic(func() { e2.LogDPanic(logger) }))

	// When
	// Logging error to Panic level
	// Then
	// Panic occurs
	require.True(t, checkPanic(func() { e2.LogPanic(logger) }))
}
