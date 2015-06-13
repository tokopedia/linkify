package linkify

import "fmt"

func Example() {
	input := `
	Check out this link to http://google.com
You can also email support@example.com to view more.

Some more links: fsf.org http://www.gnu.org/licenses/gpl-3.0.en.html 127.0.0.1 localhost:80 github.com/trending //reddit.com/r/golang mailto:r@golang.org some.nonexistent.host.name flibustahezeous3.onion
`
	for _, l := range Links(input) {
		fmt.Printf("Schema: %-8s  URL: %s\n", l.Schema, input[l.Start:l.End])
	}

	// Output:
	// Schema: http:     URL: http://google.com
	// Schema: mailto:   URL: support@example.com
	// Schema:           URL: fsf.org
	// Schema: http:     URL: http://www.gnu.org/licenses/gpl-3.0.en.html
	// Schema:           URL: 127.0.0.1
	// Schema:           URL: localhost:80
	// Schema:           URL: github.com/trending
	// Schema: //        URL: //reddit.com/r/golang
	// Schema: mailto:   URL: mailto:r@golang.org
	// Schema:           URL: some.nonexistent.host.name
	// Schema:           URL: flibustahezeous3.onion
}
