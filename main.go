package main

import (
	"github.com/go-chi/chi" //nolint:goimports
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", routers()) //nolint:errcheck
}

func routers() chi.Router {
	r := chi.NewRouter()

	r.Get("/admin/redirects", redirects)
	r.Get("/admin/redirects/{id}", redirect)
	r.Post("/admin/redirects", newRedirect)
	r.Patch("/admin/redirects/{id}", updateRedirect)
	r.Delete("/admin/redirects/{id}", deleteRedirect)

	r.Get("/redirects", uredirect)
	return r
}

func uredirect(w http.ResponseWriter, r *http.Request) {

}

func deleteRedirect(w http.ResponseWriter, r *http.Request) {

}

func updateRedirect(w http.ResponseWriter, r *http.Request) {

}

func newRedirect(w http.ResponseWriter, r *http.Request) {

}

func redirect(w http.ResponseWriter, r *http.Request) {

}

func redirects(w http.ResponseWriter, r *http.Request) {

}
