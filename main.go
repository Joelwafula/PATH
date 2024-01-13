package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// func handlerFunc(w http.ResponseWriter, r *http.Request) {
// 	w.Header.Set("Content-Type", "text/html")

// 	if r.URL.Path == "/" {
// 		fmt.Fprint(w, "<h1>welcome to this site</h1>")

// 	} else if r.URL.Path == "/contact" {
// 		fmt.Fprint(w, "to get in touch please send an email"+
// 			"to <a href=\"mailto:support@lenslocked.com\">"+
// 			"support@lenslocked.com</a>.")

// 	} else {
// 		w.WriteHeader(http.StatusNotFound)

// 		fmt.Fprint(w, "<h1>We could not find the page you "+
// 			"were looking for :(</h1>"+
// 			"<p>Please email us if you keep being sent to an "+
// 			"invalid page.</p>")
// 	}

// }

// lets break the above into functions
// we have also installed the gorrila mux packae,used to implement routes
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	fmt.Fprint(w, "<h1>welcome to our home page</h1>")

}
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "to get in touch please send an email"+
		"to <a href=\"mailto:support@lenslocked.com\">"+
		"support@lenslocked.com</a>.")
}

func FAQ(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")

	fmt.Fprint(w, "<h1>This is the FAQ page</h1>")

}
func PageNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "<h1>page is not found</h1>")

}

//now we are going to use the gorila mux in the main function

func main() {

	r := mux.NewRouter()
	notFoundPage := http.HandlerFunc(PageNotFound)

	r.NotFoundHandler = notFoundPage

	r.HandleFunc("/faq", FAQ)

	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)

	http.ListenAndServe(":3000", r)

}
