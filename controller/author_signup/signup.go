package author_signup

import (
	"fmt"
	"net/http"
)

func AuthorSignup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Author Registration Page!!</h1>")
}
