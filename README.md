# tut-01

## Interfaces

Some pointers on the usability of interfaces in golang

### httpServ.go

* `httpServ.go` makes use of the interface in order to write the http response from the server to the network socket.

* This is possible because `fmt.Fprint` allows `io.Writer` to be an interface instead of a strict method.

* `io.Writer` has just the `Write` method as criteria in it to be acceptable for further processing by `fmt.Fprintf`.

### proxyServ.go 
* `proxyServ.go` makes use of the interface in order to pass the `rp` value to the `http.Server` struct.

* This is possible because `http.Server` allows `Handler` to be a Handler interface.

* `Handler` interface has just the `ServeHTTP` method as criteria in it to be acceptable for further processing by `http.Server.ListenAndServe`.
