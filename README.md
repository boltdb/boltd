boltd ![Project status](http://img.shields.io/status/alpha.png?color=yellow)
=====

This server allows you introspect Bolt database in a web browser. The boltd tool
gives you access to low-level page information and b-tree structures so you can
better understand how Bolt is laying out your data.


## Getting Started

To install boltd, use the `go get` command:

```sh
$ go get github.com/boltdb/boltd/...
```

And then run the `boltd` binary by passing in the path to your database:

```sh
$ boltd path/to/my.db
```


## HTTP Integration

You can also use boltd as an `http.Handler` in your own application. To use it,
simply add the handler to your muxer:

```go
http.Handle("/introspect", http.StripPrefix("/introspect", boltd.NewHandler(mydb)))
```
