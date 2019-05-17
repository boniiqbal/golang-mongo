package config

import (
	"os"

	mgo "github.com/globalsign/mgo"
)

func GetMongoDB() (*mgo.Database, error) {
	host := os.Getenv("MONGO_HOST")
	dbName := os.Getenv("DB_NAME")

	session, err := mgo.Dial(host)

	if err != nil {
		return nil, err
	}

	db := session.DB(dbName)

	return db, nil

}
