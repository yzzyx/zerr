package zerr

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"time"
)

// WithRequest adds information about a http request to the given error.
// This is a convenience function that performs the same task as calling
//  err.WithField(zerr.FieldRequest("request", r))
// or
//  WrapNoStack(err, zerr.FieldRequest("request", r))
func (e *Error) WithRequest(r *http.Request) *Error {
	return e.WithField(FieldRequest("request", r))
}

// Standard zap-fields
func (e *Error) WithAny(key string, value interface{}) *Error { return e.WithField(zap.Any(key, value)) }
func (e *Error) WithArray(key string, val zapcore.ArrayMarshaler) *Error {
	return e.WithField(zap.Array(key, val))
}
func (e *Error) WithBinary(key string, val []byte) *Error { return e.WithField(zap.Binary(key, val)) }
func (e *Error) WithBool(key string, val bool) *Error     { return e.WithField(zap.Bool(key, val)) }
func (e *Error) WithBools(key string, bs []bool) *Error   { return e.WithField(zap.Bools(key, bs)) }
func (e *Error) WithByteString(key string, val []byte) *Error {
	return e.WithField(zap.ByteString(key, val))
}
func (e *Error) WithByteStrings(key string, bss [][]byte) *Error {
	return e.WithField(zap.ByteStrings(key, bss))
}
func (e *Error) WithComplex128(key string, val complex128) *Error {
	return e.WithField(zap.Complex128(key, val))
}
func (e *Error) WithComplex128s(key string, nums []complex128) *Error {
	return e.WithField(zap.Complex128s(key, nums))
}
func (e *Error) WithComplex64(key string, val complex64) *Error {
	return e.WithField(zap.Complex64(key, val))
}
func (e *Error) WithComplex64s(key string, nums []complex64) *Error {
	return e.WithField(zap.Complex64s(key, nums))
}
func (e *Error) WithDuration(key string, val time.Duration) *Error {
	return e.WithField(zap.Duration(key, val))
}
func (e *Error) WithDurations(key string, ds []time.Duration) *Error {
	return e.WithField(zap.Durations(key, ds))
}
func (e *Error) WithError(err error) *Error                 { return e.WithField(zap.Error(err)) }
func (e *Error) WithErrors(key string, errs []error) *Error { return e.WithField(zap.Errors(key, errs)) }
func (e *Error) WithFloat32(key string, val float32) *Error { return e.WithField(zap.Float32(key, val)) }
func (e *Error) WithFloat32s(key string, nums []float32) *Error {
	return e.WithField(zap.Float32s(key, nums))
}
func (e *Error) WithFloat64(key string, val float64) *Error { return e.WithField(zap.Float64(key, val)) }
func (e *Error) WithFloat64s(key string, nums []float64) *Error {
	return e.WithField(zap.Float64s(key, nums))
}
func (e *Error) WithInt(key string, val int) *Error         { return e.WithField(zap.Int(key, val)) }
func (e *Error) WithInt16(key string, val int16) *Error     { return e.WithField(zap.Int16(key, val)) }
func (e *Error) WithInt16s(key string, nums []int16) *Error { return e.WithField(zap.Int16s(key, nums)) }
func (e *Error) WithInt32(key string, val int32) *Error     { return e.WithField(zap.Int32(key, val)) }
func (e *Error) WithInt32s(key string, nums []int32) *Error { return e.WithField(zap.Int32s(key, nums)) }
func (e *Error) WithInt64(key string, val int64) *Error     { return e.WithField(zap.Int64(key, val)) }
func (e *Error) WithInt64s(key string, nums []int64) *Error { return e.WithField(zap.Int64s(key, nums)) }
func (e *Error) WithInt8(key string, val int8) *Error       { return e.WithField(zap.Int8(key, val)) }
func (e *Error) WithInt8s(key string, nums []int8) *Error   { return e.WithField(zap.Int8s(key, nums)) }
func (e *Error) WithInts(key string, nums []int) *Error     { return e.WithField(zap.Ints(key, nums)) }
func (e *Error) WithNamedError(key string, err error) *Error {
	return e.WithField(zap.NamedError(key, err))
}
func (e *Error) WithNamespace(key string) *Error { return e.WithField(zap.Namespace(key)) }
func (e *Error) WithObject(key string, val zapcore.ObjectMarshaler) *Error {
	return e.WithField(zap.Object(key, val))
}
func (e *Error) WithReflect(key string, val interface{}) *Error {
	return e.WithField(zap.Reflect(key, val))
}
func (e *Error) WithSkip() *Error                         { return e.WithField(zap.Skip()) }
func (e *Error) WithStack(key string) *Error              { return e.WithField(zap.Stack(key)) }
func (e *Error) WithString(key string, val string) *Error { return e.WithField(zap.String(key, val)) }
func (e *Error) WithStringer(key string, val fmt.Stringer) *Error {
	return e.WithField(zap.Stringer(key, val))
}
func (e *Error) WithStrings(key string, ss []string) *Error  { return e.WithField(zap.Strings(key, ss)) }
func (e *Error) WithTime(key string, val time.Time) *Error   { return e.WithField(zap.Time(key, val)) }
func (e *Error) WithTimes(key string, ts []time.Time) *Error { return e.WithField(zap.Times(key, ts)) }
func (e *Error) WithUint(key string, val uint) *Error        { return e.WithField(zap.Uint(key, val)) }
func (e *Error) WithUint16(key string, val uint16) *Error    { return e.WithField(zap.Uint16(key, val)) }
func (e *Error) WithUint16s(key string, nums []uint16) *Error {
	return e.WithField(zap.Uint16s(key, nums))
}
func (e *Error) WithUint32(key string, val uint32) *Error { return e.WithField(zap.Uint32(key, val)) }
func (e *Error) WithUint32s(key string, nums []uint32) *Error {
	return e.WithField(zap.Uint32s(key, nums))
}
func (e *Error) WithUint64(key string, val uint64) *Error { return e.WithField(zap.Uint64(key, val)) }
func (e *Error) WithUint64s(key string, nums []uint64) *Error {
	return e.WithField(zap.Uint64s(key, nums))
}
func (e *Error) WithUint8(key string, val uint8) *Error     { return e.WithField(zap.Uint8(key, val)) }
func (e *Error) WithUint8s(key string, nums []uint8) *Error { return e.WithField(zap.Uint8s(key, nums)) }
func (e *Error) WithUintptr(key string, val uintptr) *Error { return e.WithField(zap.Uintptr(key, val)) }
func (e *Error) WithUintptrs(key string, us []uintptr) *Error {
	return e.WithField(zap.Uintptrs(key, us))
}
func (e *Error) WithUints(key string, nums []uint) *Error { return e.WithField(zap.Uints(key, nums)) }

// zap bool pointer fields, not available in current uber-zap version
//func (e *Error) WithBoolp(key string, val *bool) *Error { return e.WithField(zap.Boolp(key, val)) }
//func (e *Error) WithComplex128p(key string, val *complex128) *Error { return e.WithField(zap.Complex128p(key, val)) }
//func (e *Error) WithComplex64p(key string, val *complex64) *Error { return e.WithField(zap.Complex64p(key, val)) }
//func (e *Error) WithDurationp(key string, val *time.Duration) *Error { return e.WithField(zap.Durationp(key, val)) }
//func (e *Error) WithFloat32p(key string, val *float32) *Error { return e.WithField(zap.Float32p(key, val)) }
//func (e *Error) WithFloat64p(key string, val *float64) *Error { return e.WithField(zap.Float64p(key, val)) }
//func (e *Error) WithInt16p(key string, val *int16) *Error { return e.WithField(zap.Int16p(key, val)) }
//func (e *Error) WithInt32p(key string, val *int32) *Error { return e.WithField(zap.Int32p(key, val)) }
//func (e *Error) WithInt64p(key string, val *int64) *Error { return e.WithField(zap.Int64p(key, val)) }
//func (e *Error) WithInt8p(key string, val *int8) *Error { return e.WithField(zap.Int8p(key, val)) }
//func (e *Error) WithIntp(key string, val *int) *Error { return e.WithField(zap.Intp(key, val)) }
//func (e *Error) WithStringp(key string, val *string) *Error { return e.WithField(zap.Stringp(key, val)) }
//func (e *Error) WithTimep(key string, val *time.Time) *Error { return e.WithField(zap.Timep(key, val)) }
//func (e *Error) WithUint16p(key string, val *uint16) *Error { return e.WithField(zap.Uint16p(key, val)) }
//func (e *Error) WithUint32p(key string, val *uint32) *Error { return e.WithField(zap.Uint32p(key, val)) }
//func (e *Error) WithUint64p(key string, val *uint64) *Error { return e.WithField(zap.Uint64p(key, val)) }
//func (e *Error) WithUint8p(key string, val *uint8) *Error { return e.WithField(zap.Uint8p(key, val)) }
//func (e *Error) WithUintp(key string, val *uint) *Error { return e.WithField(zap.Uintp(key, val)) }
//func (e *Error) WithUintptrp(key string, val *uintptr) *Error { return e.WithField(zap.Uintptrp(key, val)) }
