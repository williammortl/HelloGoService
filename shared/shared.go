package shared

import (
	"fmt"
	"net/http"
)

// GetIP gets a requests IP address by reading off the forwarded-for
// header (for proxies) and falls back to use the remote address.
func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

// ErrorHTML generates an HTML error message
func ErrorHTML(title string, message string) string {
	return fmt.Sprintf(`
	<html>
	<title>%v</title>
	<body>
		%v <br />
	</body>
	</html>`, title, message)
}
