package mongo

import "github.com/mongodb/mongo-go-driver/mongo"

var Client, err = mongo.Connect(cotext.Background(), "Insert your MongoDB URI here!", nil)
