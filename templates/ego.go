package templates

import (
	"fmt"
	"github.com/boltdb/bolt"
	"io"
	"path/filepath"
)

//line bucket.ego:1
func Bucket(w io.Writer, tx *bolt.Tx, b *bolt.Bucket, keys [][]byte) error {
//line bucket.ego:2
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line bucket.ego:4
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line bucket.ego:5
	if _, err := fmt.Fprintf(w, "<!DOCTYPE html>\n"); err != nil {
		return err
	}
//line bucket.ego:6
	if _, err := fmt.Fprintf(w, "<html lang=\"en\">\n  "); err != nil {
		return err
	}
//line bucket.ego:7
	head(w, tx)
//line bucket.ego:8
	if _, err := fmt.Fprintf(w, "\n\n  "); err != nil {
		return err
	}
//line bucket.ego:9
	if _, err := fmt.Fprintf(w, "<body class=\"bucket\">\n    "); err != nil {
		return err
	}
//line bucket.ego:10
	if _, err := fmt.Fprintf(w, "<div class=\"container\">\n      "); err != nil {
		return err
	}
//line bucket.ego:11
	nav(w, tx)
//line bucket.ego:12
	if _, err := fmt.Fprintf(w, "\n\n      "); err != nil {
		return err
	}
//line bucket.ego:13
	if _, err := fmt.Fprintf(w, "<table class=\"table\">\n        "); err != nil {
		return err
	}
//line bucket.ego:14
	if _, err := fmt.Fprintf(w, "<thead>\n          "); err != nil {
		return err
	}
//line bucket.ego:15
	if _, err := fmt.Fprintf(w, "<tr>\n            "); err != nil {
		return err
	}
//line bucket.ego:16
	if _, err := fmt.Fprintf(w, "<th align=\"left\">Key"); err != nil {
		return err
	}
//line bucket.ego:16
	if _, err := fmt.Fprintf(w, "</th>\n            "); err != nil {
		return err
	}
//line bucket.ego:17
	if _, err := fmt.Fprintf(w, "<th align=\"left\">Value"); err != nil {
		return err
	}
//line bucket.ego:17
	if _, err := fmt.Fprintf(w, "</th>\n          "); err != nil {
		return err
	}
//line bucket.ego:18
	if _, err := fmt.Fprintf(w, "</tr>\n        "); err != nil {
		return err
	}
//line bucket.ego:19
	if _, err := fmt.Fprintf(w, "</thead>\n        "); err != nil {
		return err
	}
//line bucket.ego:20
	if _, err := fmt.Fprintf(w, "<tbody>\n          "); err != nil {
		return err
	}
//line bucket.ego:21
	b.ForEach(func(k, v []byte) error {
//line bucket.ego:22
		if _, err := fmt.Fprintf(w, "\n            "); err != nil {
			return err
		}
//line bucket.ego:22
		if _, err := fmt.Fprintf(w, "<tr>\n              "); err != nil {
			return err
		}
//line bucket.ego:23
		if v == nil {
//line bucket.ego:24
			if _, err := fmt.Fprintf(w, "\n                "); err != nil {
				return err
			}
//line bucket.ego:24
			if _, err := fmt.Fprintf(w, "<td>"); err != nil {
				return err
			}
//line bucket.ego:24
			if _, err := fmt.Fprintf(w, "<a href=\""); err != nil {
				return err
			}
//line bucket.ego:24
			if _, err := fmt.Fprintf(w, "%v", subbucketlink(keys, k)); err != nil {
				return err
			}
//line bucket.ego:24
			if _, err := fmt.Fprintf(w, "\">"); err != nil {
				return err
			}
//line bucket.ego:24
			if _, err := fmt.Fprintf(w, "%v", trunc(tostr(k), 40)); err != nil {
				return err
			}
//line bucket.ego:24
			if _, err := fmt.Fprintf(w, "</a>"); err != nil {
				return err
			}
//line bucket.ego:24
			if _, err := fmt.Fprintf(w, "</td>\n                "); err != nil {
				return err
			}
//line bucket.ego:25
			if _, err := fmt.Fprintf(w, "<td>&lt;bucket&gt;"); err != nil {
				return err
			}
//line bucket.ego:25
			if _, err := fmt.Fprintf(w, "</td>\n              "); err != nil {
				return err
			}
//line bucket.ego:26
		} else {
//line bucket.ego:27
			if _, err := fmt.Fprintf(w, "\n                "); err != nil {
				return err
			}
//line bucket.ego:27
			if _, err := fmt.Fprintf(w, "<td>"); err != nil {
				return err
			}
//line bucket.ego:27
			if _, err := fmt.Fprintf(w, "%v", trunc(tostr(k), 40)); err != nil {
				return err
			}
//line bucket.ego:27
			if _, err := fmt.Fprintf(w, "</td>\n                "); err != nil {
				return err
			}
//line bucket.ego:28
			if _, err := fmt.Fprintf(w, "<td>"); err != nil {
				return err
			}
//line bucket.ego:28
			if _, err := fmt.Fprintf(w, "%v", trunc(tostr(v), 40)); err != nil {
				return err
			}
//line bucket.ego:28
			if _, err := fmt.Fprintf(w, "</td>\n              "); err != nil {
				return err
			}
//line bucket.ego:29
		}
//line bucket.ego:30
		if _, err := fmt.Fprintf(w, "\n            "); err != nil {
			return err
		}
//line bucket.ego:30
		if _, err := fmt.Fprintf(w, "</tr>\n          "); err != nil {
			return err
		}
//line bucket.ego:31
		return nil
	})
//line bucket.ego:32
	if _, err := fmt.Fprintf(w, "\n        "); err != nil {
		return err
	}
//line bucket.ego:32
	if _, err := fmt.Fprintf(w, "</tbody>\n      "); err != nil {
		return err
	}
//line bucket.ego:33
	if _, err := fmt.Fprintf(w, "</table>      \n    "); err != nil {
		return err
	}
//line bucket.ego:34
	if _, err := fmt.Fprintf(w, "</div> "); err != nil {
		return err
	}
//line bucket.ego:34
	if _, err := fmt.Fprintf(w, "<!-- /container -->\n  "); err != nil {
		return err
	}
//line bucket.ego:35
	if _, err := fmt.Fprintf(w, "</body>\n"); err != nil {
		return err
	}
//line bucket.ego:36
	if _, err := fmt.Fprintf(w, "</html>\n"); err != nil {
		return err
	}
	return nil
}

//line error.ego:1
func Error(w io.Writer, err error) error {
//line error.ego:2
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line error.ego:3
	if _, err := fmt.Fprintf(w, "<!DOCTYPE html>\n"); err != nil {
		return err
	}
//line error.ego:4
	if _, err := fmt.Fprintf(w, "<html lang=\"en\">\n  "); err != nil {
		return err
	}
//line error.ego:5
	if _, err := fmt.Fprintf(w, "<head>\n    "); err != nil {
		return err
	}
//line error.ego:6
	if _, err := fmt.Fprintf(w, "<meta charset=\"utf-8\">\n    "); err != nil {
		return err
	}
//line error.ego:7
	if _, err := fmt.Fprintf(w, "<title>boltd"); err != nil {
		return err
	}
//line error.ego:7
	if _, err := fmt.Fprintf(w, "</title>\n  "); err != nil {
		return err
	}
//line error.ego:8
	if _, err := fmt.Fprintf(w, "</head>\n\n  "); err != nil {
		return err
	}
//line error.ego:10
	if _, err := fmt.Fprintf(w, "<body class=\"error\">\n    "); err != nil {
		return err
	}
//line error.ego:11
	if _, err := fmt.Fprintf(w, "<div class=\"container\">\n      "); err != nil {
		return err
	}
//line error.ego:12
	if _, err := fmt.Fprintf(w, "<div class=\"header\">\n        "); err != nil {
		return err
	}
//line error.ego:13
	if _, err := fmt.Fprintf(w, "<h3 class=\"text-muted\">Error"); err != nil {
		return err
	}
//line error.ego:13
	if _, err := fmt.Fprintf(w, "</h3>\n      "); err != nil {
		return err
	}
//line error.ego:14
	if _, err := fmt.Fprintf(w, "</div>\n\n      An error has occurred: "); err != nil {
		return err
	}
//line error.ego:16
	if _, err := fmt.Fprintf(w, "%v", err); err != nil {
		return err
	}
//line error.ego:17
	if _, err := fmt.Fprintf(w, "\n    "); err != nil {
		return err
	}
//line error.ego:17
	if _, err := fmt.Fprintf(w, "</div> "); err != nil {
		return err
	}
//line error.ego:17
	if _, err := fmt.Fprintf(w, "<!-- /container -->\n  "); err != nil {
		return err
	}
//line error.ego:18
	if _, err := fmt.Fprintf(w, "</body>\n"); err != nil {
		return err
	}
//line error.ego:19
	if _, err := fmt.Fprintf(w, "</html>\n"); err != nil {
		return err
	}
	return nil
}

//line head.ego:1
func head(w io.Writer, tx *bolt.Tx) error {
//line head.ego:2
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line head.ego:4
	if _, err := fmt.Fprintf(w, "\n"); err != nil {
		return err
	}
//line head.ego:5
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line head.ego:6
	if _, err := fmt.Fprintf(w, "<head>\n  "); err != nil {
		return err
	}
//line head.ego:7
	if _, err := fmt.Fprintf(w, "<meta charset=\"utf-8\">\n  "); err != nil {
		return err
	}
//line head.ego:8
	if _, err := fmt.Fprintf(w, "<title>"); err != nil {
		return err
	}
//line head.ego:8
	if _, err := fmt.Fprintf(w, "%v", filepath.Base(tx.DB().Path())); err != nil {
		return err
	}
//line head.ego:8
	if _, err := fmt.Fprintf(w, "</title>\n"); err != nil {
		return err
	}
//line head.ego:9
	if _, err := fmt.Fprintf(w, "</head>\n"); err != nil {
		return err
	}
	return nil
}

//line index.ego:1
func Index(w io.Writer, tx *bolt.Tx) error {
//line index.ego:2
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line index.ego:4
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line index.ego:5
	if _, err := fmt.Fprintf(w, "<!DOCTYPE html>\n"); err != nil {
		return err
	}
//line index.ego:6
	if _, err := fmt.Fprintf(w, "<html lang=\"en\">\n  "); err != nil {
		return err
	}
//line index.ego:7
	head(w, tx)
//line index.ego:8
	if _, err := fmt.Fprintf(w, "\n\n  "); err != nil {
		return err
	}
//line index.ego:9
	if _, err := fmt.Fprintf(w, "<body>\n    "); err != nil {
		return err
	}
//line index.ego:10
	nav(w, tx)
//line index.ego:11
	if _, err := fmt.Fprintf(w, "\n\n    "); err != nil {
		return err
	}
//line index.ego:12
	if _, err := fmt.Fprintf(w, "<table>\n      "); err != nil {
		return err
	}
//line index.ego:13
	if _, err := fmt.Fprintf(w, "<thead>\n        "); err != nil {
		return err
	}
//line index.ego:14
	if _, err := fmt.Fprintf(w, "<tr>\n          "); err != nil {
		return err
	}
//line index.ego:15
	if _, err := fmt.Fprintf(w, "<th>Bucket Name"); err != nil {
		return err
	}
//line index.ego:15
	if _, err := fmt.Fprintf(w, "</th>\n        "); err != nil {
		return err
	}
//line index.ego:16
	if _, err := fmt.Fprintf(w, "</tr>\n      "); err != nil {
		return err
	}
//line index.ego:17
	if _, err := fmt.Fprintf(w, "</thead>\n      "); err != nil {
		return err
	}
//line index.ego:18
	if _, err := fmt.Fprintf(w, "<tbody>\n        "); err != nil {
		return err
	}
//line index.ego:19
	tx.ForEach(func(name []byte, b *bolt.Bucket) error {
//line index.ego:20
		if _, err := fmt.Fprintf(w, "\n          "); err != nil {
			return err
		}
//line index.ego:20
		if _, err := fmt.Fprintf(w, "<tr>\n            "); err != nil {
			return err
		}
//line index.ego:21
		if _, err := fmt.Fprintf(w, "<td>"); err != nil {
			return err
		}
//line index.ego:21
		if _, err := fmt.Fprintf(w, "<a href=\""); err != nil {
			return err
		}
//line index.ego:21
		if _, err := fmt.Fprintf(w, "%v", bucketlink([][]byte{name})); err != nil {
			return err
		}
//line index.ego:21
		if _, err := fmt.Fprintf(w, "\">"); err != nil {
			return err
		}
//line index.ego:21
		if _, err := fmt.Fprintf(w, "%v", trunc(tostr(name), 40)); err != nil {
			return err
		}
//line index.ego:21
		if _, err := fmt.Fprintf(w, "</a>"); err != nil {
			return err
		}
//line index.ego:21
		if _, err := fmt.Fprintf(w, "</td>\n          "); err != nil {
			return err
		}
//line index.ego:22
		if _, err := fmt.Fprintf(w, "</tr>\n        "); err != nil {
			return err
		}
//line index.ego:23
		return nil
	})
//line index.ego:24
	if _, err := fmt.Fprintf(w, "\n      "); err != nil {
		return err
	}
//line index.ego:24
	if _, err := fmt.Fprintf(w, "</tbody>\n    "); err != nil {
		return err
	}
//line index.ego:25
	if _, err := fmt.Fprintf(w, "</table>      \n  "); err != nil {
		return err
	}
//line index.ego:26
	if _, err := fmt.Fprintf(w, "</body>\n"); err != nil {
		return err
	}
//line index.ego:27
	if _, err := fmt.Fprintf(w, "</html>\n"); err != nil {
		return err
	}
	return nil
}

//line nav.ego:1
func nav(w io.Writer, tx *bolt.Tx) error {
//line nav.ego:2
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line nav.ego:4
	if _, err := fmt.Fprintf(w, "\n"); err != nil {
		return err
	}
//line nav.ego:5
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line nav.ego:6
	if _, err := fmt.Fprintf(w, "<h1>"); err != nil {
		return err
	}
//line nav.ego:6
	if _, err := fmt.Fprintf(w, "%v", filepath.Base(tx.DB().Path())); err != nil {
		return err
	}
//line nav.ego:6
	if _, err := fmt.Fprintf(w, "</h1>\n"); err != nil {
		return err
	}
	return nil
}
