package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type LinkResource struct {
	col mongo.Collection
}

func (lr LinkResource) uredirect(w http.ResponseWriter, r *http.Request) {

}

func (lr LinkResource) deleteRedirect(w http.ResponseWriter, r *http.Request) {

}

func (lr LinkResource) updateRedirect(w http.ResponseWriter, r *http.Request) {

}

func (lr LinkResource) newRedirect(w http.ResponseWriter, r *http.Request) {

}

func (lr LinkResource) redirect(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	find := lr.col.FindOne(ctx, bson.D{
		{
			Key:   "_id",
			Value: id,
		},
	}, nil)

	var link Link
	err := find.Decode(&link)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.NoContent(w, r)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, link)
}

func (lr LinkResource) redirects(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	find, err := lr.col.Find(ctx, bson.D{})
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.NoContent(w, r)
		return
	}

	links := make([]Link, 0, 1)

	for find.Next(ctx) {
		var link Link
		err := find.Decode(&link)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.NoContent(w, r)
			return
		}
		links = append(links, link)
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, links)
}
