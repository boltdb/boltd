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
	head(w)
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
	nav(w, tx, "explorer")
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
	if _, err := fmt.Fprintf(w, "<th>Key"); err != nil {
		return err
	}
//line bucket.ego:16
	if _, err := fmt.Fprintf(w, "</th>\n            "); err != nil {
		return err
	}
//line bucket.ego:17
	if _, err := fmt.Fprintf(w, "<th>Value"); err != nil {
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
	head(w)
//line error.ego:6
	if _, err := fmt.Fprintf(w, "\n\n  "); err != nil {
		return err
	}
//line error.ego:7
	if _, err := fmt.Fprintf(w, "<body class=\"error\">\n    "); err != nil {
		return err
	}
//line error.ego:8
	if _, err := fmt.Fprintf(w, "<div class=\"container\">\n      "); err != nil {
		return err
	}
//line error.ego:9
	if _, err := fmt.Fprintf(w, "<div class=\"header\">\n        "); err != nil {
		return err
	}
//line error.ego:10
	if _, err := fmt.Fprintf(w, "<h3 class=\"text-muted\">Error"); err != nil {
		return err
	}
//line error.ego:10
	if _, err := fmt.Fprintf(w, "</h3>\n      "); err != nil {
		return err
	}
//line error.ego:11
	if _, err := fmt.Fprintf(w, "</div>\n\n      An error has occurred: "); err != nil {
		return err
	}
//line error.ego:13
	if _, err := fmt.Fprintf(w, "%v", err); err != nil {
		return err
	}
//line error.ego:14
	if _, err := fmt.Fprintf(w, "\n    "); err != nil {
		return err
	}
//line error.ego:14
	if _, err := fmt.Fprintf(w, "</div> "); err != nil {
		return err
	}
//line error.ego:14
	if _, err := fmt.Fprintf(w, "<!-- /container -->\n  "); err != nil {
		return err
	}
//line error.ego:15
	if _, err := fmt.Fprintf(w, "</body>\n"); err != nil {
		return err
	}
//line error.ego:16
	if _, err := fmt.Fprintf(w, "</html>\n"); err != nil {
		return err
	}
	return nil
}

//line head.ego:1
func head(w io.Writer) error {
//line head.ego:2
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line head.ego:4
	if _, err := fmt.Fprintf(w, "\n\n"); err != nil {
		return err
	}
//line head.ego:5
	if _, err := fmt.Fprintf(w, "<head>\n  "); err != nil {
		return err
	}
//line head.ego:6
	if _, err := fmt.Fprintf(w, "<meta charset=\"utf-8\">\n  "); err != nil {
		return err
	}
//line head.ego:7
	if _, err := fmt.Fprintf(w, "<meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\">\n  "); err != nil {
		return err
	}
//line head.ego:8
	if _, err := fmt.Fprintf(w, "<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n  "); err != nil {
		return err
	}
//line head.ego:9
	if _, err := fmt.Fprintf(w, "<title>boltd"); err != nil {
		return err
	}
//line head.ego:9
	if _, err := fmt.Fprintf(w, "</title>\n  "); err != nil {
		return err
	}
//line head.ego:10
	if _, err := fmt.Fprintf(w, "<link href=\"/assets/bootstrap.min.css\" rel=\"stylesheet\">\n  "); err != nil {
		return err
	}
//line head.ego:11
	if _, err := fmt.Fprintf(w, "<link href=\"/assets/application.css\" rel=\"stylesheet\">\n  "); err != nil {
		return err
	}
//line head.ego:12
	if _, err := fmt.Fprintf(w, "<script src=\"/assets/jquery-2.1.0.min.js\">"); err != nil {
		return err
	}
//line head.ego:12
	if _, err := fmt.Fprintf(w, "</script>\n  "); err != nil {
		return err
	}
//line head.ego:13
	if _, err := fmt.Fprintf(w, "<script src=\"/assets/bootstrap.min.js\">"); err != nil {
		return err
	}
//line head.ego:13
	if _, err := fmt.Fprintf(w, "</script>\n"); err != nil {
		return err
	}
//line head.ego:14
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
	head(w)
//line index.ego:8
	if _, err := fmt.Fprintf(w, "\n\n  "); err != nil {
		return err
	}
//line index.ego:9
	if _, err := fmt.Fprintf(w, "<body class=\"index\">\n    "); err != nil {
		return err
	}
//line index.ego:10
	if _, err := fmt.Fprintf(w, "<div class=\"container\">\n      "); err != nil {
		return err
	}
//line index.ego:11
	nav(w, tx, "explorer")
//line index.ego:12
	if _, err := fmt.Fprintf(w, "\n\n      "); err != nil {
		return err
	}
//line index.ego:13
	if _, err := fmt.Fprintf(w, "<table class=\"table\">\n        "); err != nil {
		return err
	}
//line index.ego:14
	if _, err := fmt.Fprintf(w, "<thead>\n          "); err != nil {
		return err
	}
//line index.ego:15
	if _, err := fmt.Fprintf(w, "<tr>\n            "); err != nil {
		return err
	}
//line index.ego:16
	if _, err := fmt.Fprintf(w, "<th>Bucket Name"); err != nil {
		return err
	}
//line index.ego:16
	if _, err := fmt.Fprintf(w, "</th>\n          "); err != nil {
		return err
	}
//line index.ego:17
	if _, err := fmt.Fprintf(w, "</tr>\n        "); err != nil {
		return err
	}
//line index.ego:18
	if _, err := fmt.Fprintf(w, "</thead>\n        "); err != nil {
		return err
	}
//line index.ego:19
	if _, err := fmt.Fprintf(w, "<tbody>\n          "); err != nil {
		return err
	}
//line index.ego:20
	tx.ForEach(func(name []byte, b *bolt.Bucket) error {
//line index.ego:21
		if _, err := fmt.Fprintf(w, "\n            "); err != nil {
			return err
		}
//line index.ego:21
		if _, err := fmt.Fprintf(w, "<tr>\n              "); err != nil {
			return err
		}
//line index.ego:22
		if _, err := fmt.Fprintf(w, "<td>"); err != nil {
			return err
		}
//line index.ego:22
		if _, err := fmt.Fprintf(w, "<a href=\""); err != nil {
			return err
		}
//line index.ego:22
		if _, err := fmt.Fprintf(w, "%v", bucketlink([][]byte{name})); err != nil {
			return err
		}
//line index.ego:22
		if _, err := fmt.Fprintf(w, "\">"); err != nil {
			return err
		}
//line index.ego:22
		if _, err := fmt.Fprintf(w, "%v", trunc(tostr(name), 40)); err != nil {
			return err
		}
//line index.ego:22
		if _, err := fmt.Fprintf(w, "</a>"); err != nil {
			return err
		}
//line index.ego:22
		if _, err := fmt.Fprintf(w, "</td>\n            "); err != nil {
			return err
		}
//line index.ego:23
		if _, err := fmt.Fprintf(w, "</tr>\n          "); err != nil {
			return err
		}
//line index.ego:24
		return nil
	})
//line index.ego:25
	if _, err := fmt.Fprintf(w, "\n        "); err != nil {
		return err
	}
//line index.ego:25
	if _, err := fmt.Fprintf(w, "</tbody>\n      "); err != nil {
		return err
	}
//line index.ego:26
	if _, err := fmt.Fprintf(w, "</table>      \n    "); err != nil {
		return err
	}
//line index.ego:27
	if _, err := fmt.Fprintf(w, "</div> "); err != nil {
		return err
	}
//line index.ego:27
	if _, err := fmt.Fprintf(w, "<!-- /container -->\n  "); err != nil {
		return err
	}
//line index.ego:28
	if _, err := fmt.Fprintf(w, "</body>\n"); err != nil {
		return err
	}
//line index.ego:29
	if _, err := fmt.Fprintf(w, "</html>\n"); err != nil {
		return err
	}
	return nil
}

//line nav.ego:1
func nav(w io.Writer, tx *bolt.Tx, section string) error {
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
	if _, err := fmt.Fprintf(w, "<div class=\"header\">\n  "); err != nil {
		return err
	}
//line nav.ego:7
	if _, err := fmt.Fprintf(w, "<ul class=\"nav nav-pills pull-right\">\n    "); err != nil {
		return err
	}
//line nav.ego:8
	if _, err := fmt.Fprintf(w, "<li "); err != nil {
		return err
	}
//line nav.ego:8
	if section == "explorer" {
//line nav.ego:8
		if _, err := fmt.Fprintf(w, "class=\"active\""); err != nil {
			return err
		}
//line nav.ego:8
	}
//line nav.ego:8
	if _, err := fmt.Fprintf(w, ">"); err != nil {
		return err
	}
//line nav.ego:8
	if _, err := fmt.Fprintf(w, "<a href=\"/\">Explorer"); err != nil {
		return err
	}
//line nav.ego:8
	if _, err := fmt.Fprintf(w, "</a>"); err != nil {
		return err
	}
//line nav.ego:8
	if _, err := fmt.Fprintf(w, "</li>\n  "); err != nil {
		return err
	}
//line nav.ego:9
	if _, err := fmt.Fprintf(w, "</ul>\n  "); err != nil {
		return err
	}
//line nav.ego:10
	if _, err := fmt.Fprintf(w, "<h3 class=\"text-muted\">"); err != nil {
		return err
	}
//line nav.ego:10
	if _, err := fmt.Fprintf(w, "%v", filepath.Base(tx.DB().Path())); err != nil {
		return err
	}
//line nav.ego:10
	if _, err := fmt.Fprintf(w, "</h3>\n"); err != nil {
		return err
	}
//line nav.ego:11
	if _, err := fmt.Fprintf(w, "</div>\n"); err != nil {
		return err
	}
	return nil
}
