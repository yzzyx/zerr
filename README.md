zerr
====

Error wrapper that allows adding [zap](https://github.com/uber-go/zap) fields to an error, and
fetching them at a later stage

Installation
------------

To install, run:

```bash
go get gitlab.com/yzzyx/zerr
```

Wrapping errors
---------------

To wrap an error with additional field:
```go
// Add int field to error
err = zerr.Wrap(err, zap.Int("int-field", 15))
```

You can add any number of fields the the error:
```go
// Add multiple fields to error
err = zerr.Wrap(err, zap.Int("int-field", 15), zap.String("query", query), zap.Any("obj", obj))
```

By default, a stacktrace is added when an error is wrapped.
To avoid this behaviour call `zerr.WrapNoStack()` instead. This allows for a specific error to be wrapped without a stacktrace, regardless of the
  setting *DefaultAddStacktrace*.

```go
// Do not include a stacktrace
err = zerr.WrapNoStack(err)

// WrapNoStack can also take extra fields, just like Wrap()
err = zerr.WrapNoStack(err, zap.Int("int-field", 15), zap.String("query", query))
```

Note that `Wrap` will not add additional stacktraces if one was already included in the error.
If additional stacktraces should be included, it must be specified explicitly, by calling `zerr.Wrap`
with a field created with `zap.Stack()`

Errors can be wrapped multiple times. All added fields, regardless of level, will be extracted.

Sugared wrapping
----------------

For ease of use, sugared versions are availble of the Wrap()-functions, which expects an error followed by a list
of alternating string keynames and values to be passed as arguments.

```go
err := zerr.Sugar(err, "fieldname1", intvalue1, "fieldname2", stringvalue2)

// which is equal to:
err = zerr.Wrap(err, zap.Int("fieldname1", intvalue1), zap.String("fieldname2", stringvalue2))
```

A corresponding function `zerr.SugarNoStack` is available to wrap an error without a stack trace

Logging HTTP requests
---------------------

The zerr package contains a wrapper Field for conveniently logging HTTP requests.

```go
zerr.Wrap(err, zerr.FieldRequest("request", r))
```

Adding fields to errors
-----------------------

It is possible to create a new error with additional fields with the method `WithField()`

```go
ze1 := zerr.Wrap(err)
ze2 := ze.WithField(zap.Int("test", 1))
```

Using with zap
--------------

Using zerr allows capturing errors with additional context information deep down in the call stack,
and the returning this error back up to a level were logging can take place,
while still having access to the context and stacktrace to where the error actually occurred.

```go

func broken(fname string) error {
    _, err := os.Open(fname)
    if err != nil {
        return zerr.Wrap(err, zap.String("filename", fname))
    }
    // do something else
    return nil
}   

func main() {
    
    logger := zap.NewDevelopment()
    
    err = broken("/this-file-does-not-exist")
    if err != nil {
    	// Log error to logger, using the intial error as message
    	zerr.Wrap(err).LogError(logger)
    	// Or, if we want to add additional fields:
    	zerr.Wrap(err, zap.Int("extra", 15)).LogError(logger)
    	// Logging with different level
    	zerr.Wrap(err, zap.Int("extra", 15)).LogInfo(logger)
    	// Using chanining
    	zerr.Wrap(err).WithRequest(r).LogError(logger)
    	
    	// Or, to call logger with a specific message:
        logger.Error("error calling broken", zerr.Fields(err)...)
    }
    
}
```


Reading errors
--------------

In order to extract the field data from errors, use the function `zerr.Fields`.

```go
// Extract any additional fields from error and log
if err != nil {
    zap.Error("something went wrong", zerr.Fields(err)...)
}
```
