package main

import (
	"fmt"
	"net/http"

	"github.com/Gunzilla89/sandbox/controllers"
	"github.com/gorilla/mux"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Oops!</h1>")
}

func main() {

	//static views
	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers()

	//mux router
	r := mux.NewRouter()

	//not found handler
	var notFoundHandle http.Handler = http.HandlerFunc(notFound)

	//router REST handlers
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.Handle("/faq", staticC.Faq).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	r.NotFoundHandler = notFoundHandle

	http.ListenAndServe(":3000", r)
}
