package zerr

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io/ioutil"
	"net/http"
	"net/url"
)

// StringArray is an array of strings that implements zapcore.ArrayMarshaler
type StringArray []string

// MarshalLogArray encodes an array of strings with a zapcore.ArrayEncoder
func (a StringArray) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	for _, s := range a {
		enc.AppendString(s)
	}
	return nil
}

// URLValues is a wrapper around url.Values that implements zapcore.ObjectMarshaler
type URLValues url.Values

// MarshalLogObject encodes url values with a zapcore.ObjectEncoder
func (values URLValues) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	for k, v := range values {
		err = enc.AddArray(k, StringArray(v))
		if err != nil {
			return err
		}
	}
	return nil
}

// Header is a wrapper around http.Header that implements zapcore.ObjectMarshaler
type Header http.Header

// MarshalLogObject encodes headers with a zapcore.ObjectEncoder
func (h Header) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	for k, v := range h {
		err = enc.AddArray(k, StringArray(v))
		if err != nil {
			return err
		}
	}
	return nil
}

// Request is a wrapper around http.Request that implements zapcore.ObjectMarshaler
type Request struct {
	*http.Request
}

// MarshalLogObject encodes request with a zapcore.ObjectEncoder
func (r *Request) MarshalLogObject(enc zapcore.ObjectEncoder) error {

	enc.AddString("Method", r.Method)
	if r.URL != nil {
		enc.AddString("Url", r.URL.String())
	}

	err := enc.AddObject("Header", Header(r.Header))
	if err != nil {
		return err
	}

	if len(r.TransferEncoding) > 0 {
		err = enc.AddArray("TransferEncoding", StringArray(r.TransferEncoding))
		if err != nil {
			return err
		}
	}

	enc.AddString("Host", r.Host)
	enc.AddString("RemoteAddr", r.RemoteAddr)

	if r.GetBody != nil {
		if bodyCopy, err := r.GetBody(); err == nil && bodyCopy != nil {
			body, err := ioutil.ReadAll(bodyCopy)
			if err == nil {
				enc.AddByteString("Body", body)
			}
		}
	}

	return nil
}

// FieldRequest converts a http request to a zap request, which is compatible with the zap ObjectMarshaler interface
func FieldRequest(key string, r *http.Request) zap.Field {
	return zap.Object(key, &Request{r})
}
