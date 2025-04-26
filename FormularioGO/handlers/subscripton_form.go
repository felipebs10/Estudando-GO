package handlers

import (
	"fmt"
	"net/http"

	"github.com/felipebs10/curso-go/controllers"
)

func SubsriptionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "erro ao fazer parse do Form: ", err)
			return
		}

		err := controllers.Create(r.PostForm.Get("name"), r.PostForm.Get("email"))
		if err != nil {
			fmt.Fprintf(w, "erro ao fazer parse do Form: ", err)
			return
		}
	}

	http.ServeFile(w, r, "handlers/templates/subscripton_form.html")

}
