package main

import (
	"fmt"
	"http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header.Set("Content-Type", "text/html")

	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>welcome to this site</h1>")

	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "to get in touch please send an email"+
			"to <a href=\"mailto:support@lenslocked.com\">"+
			"support@lenslocked.com</a>.")

	}

}
