package user

import (
	"fmt"
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Sign Up Page!! </h1>")
}
