package zerr

import (
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// WithRequest adds information about a http request to the given error.
// This is a convenience function that performs the same task as calling
//  err.WithField(zerr.FieldRequest("request", r))
// or
//  WrapNoStack(err, zerr.FieldRequest("request", r))
// WithRequest adds a zap.Request field to Error
func (e *Error) WithRequest(r *http.Request) *Error {
	return e.WithField(FieldRequest("request", r))
}

// WithAny adds a zap.Any field to Error
func (e *Error) WithAny(key string, value interface{}) *Error {
	return e.WithField(zap.Any(key, value))
}

// WithArray adds a zap.Array field to Error
func (e *Error) WithArray(key string, val zapcore.ArrayMarshaler) *Error {
	return e.WithField(zap.Array(key, val))
}

// WithBinary adds a zap.Binary field to Error
func (e *Error) WithBinary(key string, val []byte) *Error { return e.WithField(zap.Binary(key, val)) }

// WithBool adds a zap.Bool field to Error
func (e *Error) WithBool(key string, val bool) *Error { return e.WithField(zap.Bool(key, val)) }

// WithBoolp adds a zap.Boolp field to Error
func (e *Error) WithBoolp(key string, val *bool) *Error { return e.WithField(zap.Boolp(key, val)) }

// WithBools adds a zap.Bools field to Error
func (e *Error) WithBools(key string, bs []bool) *Error { return e.WithField(zap.Bools(key, bs)) }

// WithByteString adds a zap.ByteString field to Error
func (e *Error) WithByteString(key string, val []byte) *Error {
	return e.WithField(zap.ByteString(key, val))
}

// WithByteStrings adds a zap.ByteStrings field to Error
func (e *Error) WithByteStrings(key string, bss [][]byte) *Error {
	return e.WithField(zap.ByteStrings(key, bss))
}

// WithComplex128 adds a zap.Complex128 field to Error
func (e *Error) WithComplex128(key string, val complex128) *Error {
	return e.WithField(zap.Complex128(key, val))
}

// WithComplex128p adds a zap.Complex128p field to Error
func (e *Error) WithComplex128p(key string, val *complex128) *Error {
	return e.WithField(zap.Complex128p(key, val))
}

// WithComplex128s adds a zap.Complex128s field to Error
func (e *Error) WithComplex128s(key string, nums []complex128) *Error {
	return e.WithField(zap.Complex128s(key, nums))
}

// WithComplex64 adds a zap.Complex64 field to Error
func (e *Error) WithComplex64(key string, val complex64) *Error {
	return e.WithField(zap.Complex64(key, val))
}

// WithComplex64p adds a zap.Complex64p field to Error
func (e *Error) WithComplex64p(key string, val *complex64) *Error {
	return e.WithField(zap.Complex64p(key, val))
}

// WithComplex64s adds a zap.Complex64s field to Error
func (e *Error) WithComplex64s(key string, nums []complex64) *Error {
	return e.WithField(zap.Complex64s(key, nums))
}

// WithDuration adds a zap.Duration field to Error
func (e *Error) WithDuration(key string, val time.Duration) *Error {
	return e.WithField(zap.Duration(key, val))
}

// WithDurationp adds a zap.Durationp field to Error
func (e *Error) WithDurationp(key string, val *time.Duration) *Error {
	return e.WithField(zap.Durationp(key, val))
}

// WithDurations adds a zap.Durations field to Error
func (e *Error) WithDurations(key string, ds []time.Duration) *Error {
	return e.WithField(zap.Durations(key, ds))
}

// WithError adds a zap.Error field to Error
func (e *Error) WithError(err error) *Error { return e.WithField(zap.Error(err)) }

// WithErrors adds a zap.Errors field to Error
func (e *Error) WithErrors(key string, errs []error) *Error {
	return e.WithField(zap.Errors(key, errs))
}

// WithFloat32 adds a zap.Float32 field to Error
func (e *Error) WithFloat32(key string, val float32) *Error {
	return e.WithField(zap.Float32(key, val))
}

// WithFloat32p adds a zap.Float32p field to Error
func (e *Error) WithFloat32p(key string, val *float32) *Error {
	return e.WithField(zap.Float32p(key, val))
}

// WithFloat32s adds a zap.Float32s field to Error
func (e *Error) WithFloat32s(key string, nums []float32) *Error {
	return e.WithField(zap.Float32s(key, nums))
}

// WithFloat64 adds a zap.Float64 field to Error
func (e *Error) WithFloat64(key string, val float64) *Error {
	return e.WithField(zap.Float64(key, val))
}

// WithFloat64p adds a zap.Float64p field to Error
func (e *Error) WithFloat64p(key string, val *float64) *Error {
	return e.WithField(zap.Float64p(key, val))
}

// WithFloat64s adds a zap.Float64s field to Error
func (e *Error) WithFloat64s(key string, nums []float64) *Error {
	return e.WithField(zap.Float64s(key, nums))
}

// WithInline adds a zap.Inline field to Error
func (e *Error) WithInline(val zapcore.ObjectMarshaler) *Error { return e.WithField(zap.Inline(val)) }

// WithInt adds a zap.Int field to Error
func (e *Error) WithInt(key string, val int) *Error { return e.WithField(zap.Int(key, val)) }

// WithIntp adds a zap.Intp field to Error
func (e *Error) WithIntp(key string, val *int) *Error { return e.WithField(zap.Intp(key, val)) }

// WithInts adds a zap.Ints field to Error
func (e *Error) WithInts(key string, nums []int) *Error { return e.WithField(zap.Ints(key, nums)) }

// WithInt16 adds a zap.Int16 field to Error
func (e *Error) WithInt16(key string, val int16) *Error { return e.WithField(zap.Int16(key, val)) }

// WithInt16p adds a zap.Int16p field to Error
func (e *Error) WithInt16p(key string, val *int16) *Error { return e.WithField(zap.Int16p(key, val)) }

// WithInt16s adds a zap.Int16s field to Error
func (e *Error) WithInt16s(key string, nums []int16) *Error {
	return e.WithField(zap.Int16s(key, nums))
}

// WithInt32 adds a zap.Int32 field to Error
func (e *Error) WithInt32(key string, val int32) *Error { return e.WithField(zap.Int32(key, val)) }

// WithInt32p adds a zap.Int32p field to Error
func (e *Error) WithInt32p(key string, val *int32) *Error { return e.WithField(zap.Int32p(key, val)) }

// WithInt32s adds a zap.Int32s field to Error
func (e *Error) WithInt32s(key string, nums []int32) *Error {
	return e.WithField(zap.Int32s(key, nums))
}

// WithInt64 adds a zap.Int64 field to Error
func (e *Error) WithInt64(key string, val int64) *Error { return e.WithField(zap.Int64(key, val)) }

// WithInt64p adds a zap.Int64p field to Error
func (e *Error) WithInt64p(key string, val *int64) *Error { return e.WithField(zap.Int64p(key, val)) }

// WithInt64s adds a zap.Int64s field to Error
func (e *Error) WithInt64s(key string, nums []int64) *Error {
	return e.WithField(zap.Int64s(key, nums))
}

// WithInt8 adds a zap.Int8 field to Error
func (e *Error) WithInt8(key string, val int8) *Error { return e.WithField(zap.Int8(key, val)) }

// WithInt8p adds a zap.Int8p field to Error
func (e *Error) WithInt8p(key string, val *int8) *Error { return e.WithField(zap.Int8p(key, val)) }

// WithInt8s adds a zap.Int8s field to Error
func (e *Error) WithInt8s(key string, nums []int8) *Error { return e.WithField(zap.Int8s(key, nums)) }

// WithNamedError adds a zap.NamedError field to Error
func (e *Error) WithNamedError(key string, err error) *Error {
	return e.WithField(zap.NamedError(key, err))
}

// WithNamespace adds a zap.Namespace field to Error
func (e *Error) WithNamespace(key string) *Error { return e.WithField(zap.Namespace(key)) }

// WithObject adds a zap.Object field to Error
func (e *Error) WithObject(key string, val zapcore.ObjectMarshaler) *Error {
	return e.WithField(zap.Object(key, val))
}

// WithReflect adds a zap.Reflect field to Error
func (e *Error) WithReflect(key string, val interface{}) *Error {
	return e.WithField(zap.Reflect(key, val))
}

// WithSkip adds a zap.Skip field to Error
func (e *Error) WithSkip() *Error { return e.WithField(zap.Skip()) }

// WithStack adds a zap.Stack field to Error
func (e *Error) WithStack(key string) *Error { return e.WithField(zap.Stack(key)) }

// WithStackSkip adds a zap.StackSkip field to Error
func (e *Error) WithStackSkip(key string, skip int) *Error {
	return e.WithField(zap.StackSkip(key, skip))
}

// WithString adds a zap.String field to Error
func (e *Error) WithString(key string, val string) *Error { return e.WithField(zap.String(key, val)) }

// WithStringp adds a zap.Stringp field to Error
func (e *Error) WithStringp(key string, val *string) *Error {
	return e.WithField(zap.Stringp(key, val))
}

// WithStringer adds a zap.Stringer field to Error
func (e *Error) WithStringer(key string, val fmt.Stringer) *Error {
	return e.WithField(zap.Stringer(key, val))
}

// WithStrings adds a zap.Strings field to Error
func (e *Error) WithStrings(key string, ss []string) *Error { return e.WithField(zap.Strings(key, ss)) }

// WithTime adds a zap.Time field to Error
func (e *Error) WithTime(key string, val time.Time) *Error { return e.WithField(zap.Time(key, val)) }

// WithTimep adds a zap.Timep field to Error
func (e *Error) WithTimep(key string, val *time.Time) *Error { return e.WithField(zap.Timep(key, val)) }

// WithTimes adds a zap.Times field to Error
func (e *Error) WithTimes(key string, ts []time.Time) *Error { return e.WithField(zap.Times(key, ts)) }

// WithUint adds a zap.Uint field to Error
func (e *Error) WithUint(key string, val uint) *Error { return e.WithField(zap.Uint(key, val)) }

// WithUintp adds a zap.Uintp field to Error
func (e *Error) WithUintp(key string, val *uint) *Error { return e.WithField(zap.Uintp(key, val)) }

// WithUints adds a zap.Uints field to Error
func (e *Error) WithUints(key string, nums []uint) *Error { return e.WithField(zap.Uints(key, nums)) }

// WithUint16 adds a zap.Uint16 field to Error
func (e *Error) WithUint16(key string, val uint16) *Error { return e.WithField(zap.Uint16(key, val)) }

// WithUint16p adds a zap.Uint16p field to Error
func (e *Error) WithUint16p(key string, val *uint16) *Error {
	return e.WithField(zap.Uint16p(key, val))
}

// WithUint16s adds a zap.Uint16s field to Error
func (e *Error) WithUint16s(key string, nums []uint16) *Error {
	return e.WithField(zap.Uint16s(key, nums))
}

// WithUint32 adds a zap.Uint32 field to Error
func (e *Error) WithUint32(key string, val uint32) *Error { return e.WithField(zap.Uint32(key, val)) }

// WithUint32p adds a zap.Uint32p field to Error
func (e *Error) WithUint32p(key string, val *uint32) *Error {
	return e.WithField(zap.Uint32p(key, val))
}

// WithUint32s adds a zap.Uint32s field to Error
func (e *Error) WithUint32s(key string, nums []uint32) *Error {
	return e.WithField(zap.Uint32s(key, nums))
}

// WithUint64 adds a zap.Uint64 field to Error
func (e *Error) WithUint64(key string, val uint64) *Error { return e.WithField(zap.Uint64(key, val)) }

// WithUint64p adds a zap.Uint64p field to Error
func (e *Error) WithUint64p(key string, val *uint64) *Error {
	return e.WithField(zap.Uint64p(key, val))
}

// WithUint64s adds a zap.Uint64s field to Error
func (e *Error) WithUint64s(key string, nums []uint64) *Error {
	return e.WithField(zap.Uint64s(key, nums))
}

// WithUint8 adds a zap.Uint8 field to Error
func (e *Error) WithUint8(key string, val uint8) *Error { return e.WithField(zap.Uint8(key, val)) }

// WithUint8p adds a zap.Uint8p field to Error
func (e *Error) WithUint8p(key string, val *uint8) *Error { return e.WithField(zap.Uint8p(key, val)) }

// WithUint8s adds a zap.Uint8s field to Error
func (e *Error) WithUint8s(key string, nums []uint8) *Error {
	return e.WithField(zap.Uint8s(key, nums))
}

// WithUintptr adds a zap.Uintptr field to Error
func (e *Error) WithUintptr(key string, val uintptr) *Error {
	return e.WithField(zap.Uintptr(key, val))
}

// WithUintptrp adds a zap.Uintptrp field to Error
func (e *Error) WithUintptrp(key string, val *uintptr) *Error {
	return e.WithField(zap.Uintptrp(key, val))
}

// WithUintptrs adds a zap.Uintptrs field to Error
func (e *Error) WithUintptrs(key string, us []uintptr) *Error {
	return e.WithField(zap.Uintptrs(key, us))
}
