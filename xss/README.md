# XSS
```
// exploit
http://localhost:8080/?input=<script>alert('XSS')</script>
```

How to find in code
```bash
# Semgrep
$ semgrep --config="r/go.lang.security.injection.raw-html-format.raw-html-format"
$ semgrep --config="r/go.lang.security.audit.xss"

# CodeQL
# use rule https://codeql.github.com/codeql-query-help/go/go-reflected-xss/

$ codeql database create gva_codeql_db --language=go
$ run > xss.ql
```

How to avoid:
```go
// option 1
import (
	"html" // use html.EscapeString
    "fmt"
    "net/http" // <= delete
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        unsafeInput := r.URL.Query().Get("input") 
		safeInput := html.EscapeString(unsafeInput) // Escape special characters
		response := fmt.Sprintf("<h1>Hello, %s!</h1>", safeInput) // Reflect input directly
        w.Header().Set("Content-Type", "text/html")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(response)) // Protected to XSS


// option 2
import "html/template"

tmpl := template.Must(template.New("example").Parse("<h1>Hello, {{.}}</h1>"))
tmpl.Execute(w, unsafeInput) // Automatically escapes user input
```
