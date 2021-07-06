package controllers

import (
	"fmt"
	"net/http"

	"github.com/Gunzilla89/sandbox/views"
)

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/newUser"),
	}
}

type Users struct {
	NewView *views.View
}

type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

//Get /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

//create/POST
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := ParseForm(r, &form); err != nil {
		panic(err)
	}

	fmt.Fprint(w, "Email is ", form.Email)
	fmt.Fprint(w, "Password is ", form.Password)

}
