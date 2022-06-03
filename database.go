package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	client *mongo.Client
	db *mongo.Database
	users *mongo.Collection
}

type RecordId interface{}

func OpenDatabase(uri string) (Database, error) {
	var (
		db Database
		err error
	)
	
	ctx := context.TODO()
	opts := options.Client().ApplyURI(uri)

	if db.client, err = mongo.Connect(ctx, opts); err != nil {
		return db, err
	}

	db.db = db.client.Database("wedelve")
	db.users = db.db.Collection("users")

	return db, nil
}

func (db Database) Close() error {
	return db.client.Disconnect(context.TODO())
}

func (db Database) InsertUser(user User) (RecordId, error) {
	var (
		result *mongo.InsertOneResult
		err error
	)

	user.Password = ""
	
	result, err =  db.users.InsertOne(context.TODO(), user)
	
	return result.InsertedID, err
}
