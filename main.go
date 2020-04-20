package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my awesome site</h1>")
}

func contact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch please send an email to <a href=\"mailto:brennon.loveless@gmail.com\">brennon.loveless@gmail.com</a>")
}

func faq(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "This is the FAQ page. Lets answer some questions")
}

func handle404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "Hot damn... couldn't find that one")
}

func main() {
	r := httprouter.New()
	r.NotFound = http.HandlerFunc(handle404)

	r.GET("/", home)
	r.GET("/contact", contact)
	r.GET("/faq", faq)

	http.ListenAndServe(":3000", r)
}
