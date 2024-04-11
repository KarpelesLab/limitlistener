# limitlistener

This library makes it easy to limit the number of live threads for a given listener. This
can be useful in a number of cases to ensure not too many connections will be processed
at the same time rather than letting go create new goroutines like there is no tomorrow.

## Usage

```go
l, err := net.Listen("tcp", ":12345")
if err != nil {
    // ...
}

// limit to 128 concurrent processes
l = limitlistener.New(l, 128)
```

This can be used for a `http.Server` or anything that takes a `net.Listener` and will
apply the limit as standard as possible. Multiple threads calling Accept() are also
supported.
