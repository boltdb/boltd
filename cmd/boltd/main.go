package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/boltdb/bolt"
	"github.com/boltdb/boltd"
)

func main() {
	log.SetFlags(0)
	var (
		addr = flag.String("addr", ":9000", "bind address")
	)
	flag.Parse()

	// Validate parameters.
	var path = flag.Arg(0)
	if path == "" {
		log.Fatal("path required")
	}

	// Open the database.
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Enable logging.
	log.SetFlags(log.LstdFlags)

	// Setup the HTTP handlers.
	http.Handle("/", boltd.NewHandler(db))

	// Start the HTTP server.
	go func() { log.Fatal(http.ListenAndServe(*addr, nil)) }()

	fmt.Printf("Listening on http://localhost%s\n", *addr)
	select {}
}
