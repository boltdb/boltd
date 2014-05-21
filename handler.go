package boltd

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/boltdb/boltd/templates"
	"github.com/gorilla/mux"
)

// NewHandler returns a new root HTTP handler.
func NewHandler(db *bolt.DB) http.Handler {
	h := &handler{db}
	r := mux.NewRouter()
	r.HandleFunc("/", h.index)
	r.HandleFunc("/page", h.page)
	return r
}

type handler struct {
	db *bolt.DB
}

func (h *handler) index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "page", http.StatusFound)
}

func (h *handler) page(w http.ResponseWriter, r *http.Request) {
	err := h.db.View(func(tx *bolt.Tx) error {
		// Extract the indexes.
		indexes, err := indexes(r)
		if err != nil {
			return err
		}
		return templates.Page(w, tx, indexes)
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
