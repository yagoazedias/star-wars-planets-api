package repository

import (
	"github.com/yagoazedias/star-wars-planets-api/config"
	"gopkg.in/mgo.v2"
	"log"
	"os"
)

var db *mgo.Database

type Mongo struct {
	Host   string
	Database string
	Password string
	User string
}

func (m *Mongo) Connect() {
	session, err := mgo.Dial(m.Host)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func getFromEnvIfExists() Mongo {

	var host = os.Getenv("DB_HOST")
	var database = os.Getenv("DB_NAME")
	var password = os.Getenv("DB_PASSWORD")
	var user = os.Getenv("DB_USER")

	_, hostExist := os.LookupEnv("DB_HOST")
	_, hostDatabase := os.LookupEnv("DB_NAME")
	_, hostPassword := os.LookupEnv("DB_PASSWORD")
	_, hostUser := os.LookupEnv("DB_USER")

	if !hostExist || !hostDatabase || !hostPassword || !hostUser {
		return Mongo {
			Host: config.DATABASE_HOST,
			Database: config.DATABASE_NAME,
			Password: config.DATABASE_PASSWORD,
			User: config.DATABASE_USER,
		}
	}

	return Mongo{Host: host, Database: database, Password: password, User: user}
}

var mongo = getFromEnvIfExists()
