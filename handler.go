package boltd

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"path"

	"github.com/boltdb/bolt"
	"github.com/boltdb/boltd/assets"
	"github.com/boltdb/boltd/templates"
	"github.com/gorilla/mux"
)

var (
	ErrNotFound = errors.New("not found")
)

// NewHandler returns a new root HTTP handler.
func NewHandler(db *bolt.DB) http.Handler {
	h := &handler{db}
	r := mux.NewRouter()
	r.HandleFunc("/", h.index)
	r.HandleFunc("/assets/{filename}", h.assets)
	r.HandleFunc("/bucket", h.bucket)
	return r
}

type handler struct {
	db *bolt.DB
}

func (h *handler) index(w http.ResponseWriter, r *http.Request) {
	h.db.View(func(tx *bolt.Tx) error {
		return templates.Index(w, tx)
	})
}

func (h *handler) assets(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]
	b, err := assets.Asset(filename)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	switch path.Ext(filename) {
	case ".css":
		w.Header().Set("Content-Type", "text/css")
	case ".js":
		w.Header().Set("Content-Type", "application/javascript")
	}
	w.Write(b)
}

func (h *handler) bucket(w http.ResponseWriter, r *http.Request) {
	err := h.db.View(func(tx *bolt.Tx) error {
		// Extract the keys.
		keys, err := keys(r)
		if err != nil {
			return err
		}

		// Traverse to the appropriate bucket.
		b, err := bucket(tx, keys)
		if err != nil {
			return err
		}
		return templates.Bucket(w, tx, b, keys)
	})
	if err != nil {
		Error(w, err)
	}
}

func Error(w http.ResponseWriter, err error) {
	switch err {
	case ErrNotFound:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	templates.Error(w, err)
}

// parses and returns all keys from a request.
func keys(r *http.Request) ([][]byte, error) {
	r.ParseForm()
	var keys [][]byte
	for i, k := range r.Form["key"] {
		key, err := base64.StdEncoding.DecodeString(k)
		if err != nil {
			return nil, fmt.Errorf("key(%d): %s: %s", i, k, err)
		}
		keys = append(keys, key)
	}
	return keys, nil
}

// traverses the bucket hierarchy and returns the last bucket.
func bucket(tx *bolt.Tx, keys [][]byte) (*bolt.Bucket, error) {
	if len(keys) == 0 {
		return nil, ErrNotFound
	}

	var b = tx.Bucket([]byte(keys[0]))
	if b == nil {
		return nil, ErrNotFound
	}

	for _, k := range keys[1:] {
		if b = b.Bucket([]byte(k)); b == nil {
			return nil, ErrNotFound
		}
	}

	return b, nil
}
