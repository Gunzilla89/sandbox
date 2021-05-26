package main

import (
	"fmt"
	"net/http"

	"github.com/Gunzilla89/sandbox/views"
	"github.com/gorilla/mux"
)

var homeView *views.View
var contactView *views.View
var faqView *views.View
var signupView *views.View

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	checkAndRender(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	checkAndRender(contactView.Render(w, nil))
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	checkAndRender(faqView.Render(w, nil))
}

func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	checkAndRender(signupView.Render(w, nil))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Oops!</h1>")
}

func checkAndRender(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	faqView = views.NewView("bootstrap", "views/faq.gohtml")
	signupView = views.NewView("bootstrap", "views/signup.gohtml")

	var notFoundHandle http.Handler = http.HandlerFunc(notFound)
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.HandleFunc("/signup", signup)
	r.NotFoundHandler = notFoundHandle

	http.ListenAndServe(":3000", r)
}
