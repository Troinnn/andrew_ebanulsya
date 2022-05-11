package main

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

func load_from_file(col mongo.Collection) {
	data, err := os.ReadFile("links.json")
	if err != nil {
		return
	}
	var links []Link
	err = json.Unmarshal(data, &links)
	if err != nil {
		return
	}

	for _, link := range links {
		link.ID = primitive.NewObjectID().Hex()
		find, err := col.Find(context.Background(), bson.D{
			{
				Key:   "active_link",
				Value: link.ActiveLink,
			},
			{
				Key:   "history_link",
				Value: link.HistoryLink,
			},
		})

		if err != nil || find.RemainingBatchLength() != 0 {
			continue
		}

		col.InsertOne(context.Background(), link)
	}
}
