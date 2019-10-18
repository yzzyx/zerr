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

A shorthand is available for including a stacktrace:
```go
// Add stacktrace to error
err = zerr.WrapStack(err)

// WrapStack can also take extra fields, just like Wrap()
err = zerr.WrapStack(err, zap.Int("int-field", 15))
```

Note that `WrapStack` will not add additional stacktraces if one was already included in the error.
If additional stacktraces should be included, it must be specified explicitly, by calling `zerr.Wrap`
with a field created with `zap.Stack()`

Errors can be wrapped multiple times. All added fields, regardless of level, will be extracted.

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
        logger.Error("error calling broken", zerr.Fields(err))
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



