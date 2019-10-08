package userhomepage

import (
	"app"
	"net/http"
	"templates"
)

func UserHomePageHandler(w http.ResponseWriter, r *http.Request) {

	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templates.RenderTemplate(w, "userhomepage", session.Values["profile"])
}
