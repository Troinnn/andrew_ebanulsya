package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func database() (*mongo.Database, error) {
	client, err := mongo.Connect(context.Background())
	if err != nil {
		return nil, err
	}
	return client.Database("redirects"), nil
}

func main() {
	db, err := database()
	if err != nil {
		fmt.Printf("connect to mongodb: %s\n", err.Error())
		return
	}

	linksResource := LinkResource{
		col: *db.Collection("links"),
	}

	http.ListenAndServe(":8080", routers(linksResource))
}

func routers(lr LinkResource) chi.Router {
	r := chi.NewRouter()

	r.Get("/admin/redirects", lr.redirects)
	r.Get("/admin/redirects/{id}", lr.redirect)
	r.Post("/admin/redirects", lr.newRedirect)
	r.Patch("/admin/redirects/{id}", lr.updateRedirect)
	r.Delete("/admin/redirects/{id}", lr.deleteRedirect)

	r.Get("/redirects", lr.uredirect)

	return r
}
