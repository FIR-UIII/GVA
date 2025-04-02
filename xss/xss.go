package xss

import (
	// "html"
    "fmt"
    "net/http"
)

func XSS() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        unsafeInput := r.URL.Query().Get("input") // Get user input from query parameter
		// safeInput := html.EscapeString(unsafeInput) // Escape special characters
		response := fmt.Sprintf("<h1>Hello, %s!</h1>", unsafeInput) // Reflect input directly
        w.Header().Set("Content-Type", "text/html")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(response)) // Vulnerable to XSS
    })

    fmt.Println("Server running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}