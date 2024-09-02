package controllers

import (
	"log"
	"net/http"
	"reakgo/models"
)

func AddForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		} else {
			// Logic
			name := r.FormValue("name")
			address := r.FormValue("address")
			email := r.FormValue("email")
			password := r.FormValue("password")

			log.Println(name, address, email, password)
			models.FormAddView{}.Add(name, address, email, password)
		}
	}
	Helper.RenderTemplate(w, r, "addForm", nil)
}

func ViewForm(w http.ResponseWriter, r *http.Request) {
	result, err := models.FormAddView{}.View()
	if err != nil {
		log.Println(err)
	}
	Helper.RenderTemplate(w, r, "viewForm", result)
}
