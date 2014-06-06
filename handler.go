package boltd

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/boltdb/boltd/templates"
)

// NewHandler returns a new root HTTP handler.
func NewHandler(db *bolt.DB) http.Handler {
	h := &handler{db}
	mux := http.NewServeMux()
	mux.HandleFunc("/", h.index)
	mux.HandleFunc("/page", h.page)
	return mux
}

type handler struct {
	db *bolt.DB
}

func (h *handler) index(w http.ResponseWriter, r *http.Request) {
	templates.Index(w)
}

func (h *handler) page(w http.ResponseWriter, r *http.Request) {
	err := h.db.View(func(tx *bolt.Tx) error {
		showUsage := (r.FormValue("usage") == "true")

		// Use the direct page id, if available.
		if r.FormValue("id") != "" {
			id, _ := strconv.Atoi(r.FormValue("id"))
			return templates.Page(w, r, tx, nil, id, showUsage)
		}

		// Otherwise extract the indexes and traverse.
		indexes, err := indexes(r)
		if err != nil {
			return err
		}

		return templates.Page(w, r, tx, indexes, 0, showUsage)
	})
	if err != nil {
		templates.Error(w, err)
	}
}

// parses and returns all indexes from a request.
func indexes(r *http.Request) ([]int, error) {
	var a = []int{0}
	if len(r.FormValue("index")) > 0 {
		for _, s := range strings.Split(r.FormValue("index"), ":") {
			i, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			a = append(a, i)
		}
	}
	return a, nil
}
