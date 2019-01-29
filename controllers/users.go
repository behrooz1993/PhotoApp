package controllers

import (
	"fmt"
	"net/http"

	"github.com/photoApp/views"
)

type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
	}
}

type Users struct {
	NewView *views.View
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	var form SignupForm
	if err := ParseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, form)
}
