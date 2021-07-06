package controllers

import "github.com/Gunzilla89/sandbox/views"

//http handlers for the Handlers in our main.go
func NewStatic() *Static {
	return &Static{
		Home: views.NewView(
			"bootstrap", "static/home"),
		Contact: views.NewView(
			"bootstrap", "static/contact"),
		Faq: views.NewView(
			"bootstrap", "static/faq"),
	}
}

//static interface made up of views
type Static struct {
	Home    *views.View
	Contact *views.View
	Faq     *views.View
}
