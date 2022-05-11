package main

type Link struct {
	ActiveLink  string `json:"active_link" bson:"active_link"`
	HistoryLink string `json:"history_link" bson:"history_link"`
}
