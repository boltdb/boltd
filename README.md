boltd
=====

This server allows you to view a Bolt database in a web browser. Because Bolt
can only be used by a single process, you cannot use the browser while another
application is using the database file.


## Getting Started

To install boltd, use the `go get` command:

```sh
$ go get github.com/boltdb/boltd/...
```

And then run the `boltd` binary by passing in the path to your database:

```sh
$ boltd path/to/my.db
```

