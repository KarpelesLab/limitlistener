[![GoDoc](https://godoc.org/github.com/KarpelesLab/limitlistener?status.svg)](https://godoc.org/github.com/KarpelesLab/limitlistener)

# limitlistener

This library makes it easy to limit the number of live threads for a given listener. This
can be useful in a number of cases to ensure not too many connections will be processed
at the same time rather than letting go create new goroutines like there is no tomorrow.

Any call to `Accept()` when the limit has been reached will wait for any of the currently
running connections to close before accepting any new connection.

## Usage

```go
l, err := net.Listen("tcp", ":12345")
if err != nil {
    // ...
}

// limit to 128 concurrent processes
l = limitlistener.New(l, 128)

for {
    c, err := l.Accept()
    if err != nil {
        return err
    }

    // safe to start a goroutine here, will be limited to 128 routines
    go handleClient(c)
}

handleClient(c net.Conn) {
    defer c.Close()

    // ...
}
```

This can be used for a `http.Server` or anything that takes a `net.Listener` and will
apply the limit as standard as possible. Multiple threads calling Accept() are also
supported.

If using this with `http.Server` be careful to set a `ReadTimeout` or you may end
blocked just with idle connections.
