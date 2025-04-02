# XSS
```
// exploit
http://localhost:8080/?input=<script>alert('XSS')</script>
```

How to find in code
```
Semgrep

CodeQL
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
