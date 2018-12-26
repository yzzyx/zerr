package zerr

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

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
	// one extra field should be available
	fields := Fields(err)
	require.Len(t, fields, 1)

	// When
	// We wrap the error again
	err = Wrap(err, zap.String("stringfield", "abc"))

	// Then
	// Two fields should be returned (int and string)
	fields = Fields(err)
	require.Len(t, fields, 2)
}

func TestWrapStack(t *testing.T) {
	err := errors.New("test")

	// When
	// we wrap an error
	err = WrapStack(err, zap.Int("intfield", 1))

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
	err = WrapStack(err)

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

func TestSugar(t *testing.T) {
	originalError := errors.New("original error")

	// When
	// we wrap an error with Sugar() and non-strongly typed keys
	err := Sugar(originalError, "intvalue", 1)

	// Then
	// returned error should be of type Error
	_, ok := err.(*Error)
	require.True(t, ok)

	// And
	// one field should be available
	fields := Fields(err)
	require.Len(t, fields, 1)
	for _, field := range fields {
		require.IsType(t, zap.Field{}, field)
	}

	// When
	// we wrap an error with Sugar() using strongly typed keys
	err = Sugar(originalError, zap.Int("intvalue", 1), zap.String("stringvalue", " abc"))

	// Then
	// two field should be available
	fields = Fields(err)
	require.Len(t, fields, 2)
	for _, field := range fields {
		require.IsType(t, zap.Field{}, field)
	}

	// When
	// we wrap an error with Sugar() using a mix of strongly typed keys and non-types keys
	err = Sugar(originalError, zap.Int("intvalue", 1), zap.String("stringvalue", " abc"), "value3", "test")

	// Then
	// three field should be available
	fields = Fields(err)
	require.Len(t, fields, 3)
	for _, field := range fields {
		require.IsType(t, zap.Field{}, field)
	}

	// When
	// we wrap an error with a dangling key, it should be added as an "any"-field
	err = Sugar(originalError, "dangling key")

	// Then
	// we should have one field (with key "dangling-key")
	fields = Fields(err)
	require.Len(t, fields, 1)
}
