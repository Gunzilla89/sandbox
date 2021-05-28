package controllers

import (
	"fmt"
	"net/http"

	"github.com/Gunzilla89/sandbox/views"
)

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/newUser.gohtml"),
	}
}

type Users struct {
	NewView *views.View
}

//Get /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

//create/POST
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Temp response")
}
