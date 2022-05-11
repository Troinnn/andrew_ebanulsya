package main

import (
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

}

func (lr LinkResource) redirects(w http.ResponseWriter, r *http.Request) {

}
