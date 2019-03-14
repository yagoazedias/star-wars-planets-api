package repository

import (
	"fmt"
	"github.com/yagoazedias/star-wars-planets-api/environment"
	"gopkg.in/mgo.v2"
	"log"
	"os"
)

type MongoDB struct {
	Host   string
	Database string
	Password string
	Port string
	User string
	db *mgo.Database
}

func (m *MongoDB) ConnectionUrl() string {
	return fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/%s",
			m.User, m.Password, m.Host, m.Port, m.Database,
	)
}

func (m *MongoDB) Connect() {
	session, err := mgo.Dial(m.ConnectionUrl())
	if err != nil {
		log.Fatal(err)
	}
	m.db = session.DB(m.Database)
}

func getFromEnvIfExists() MongoDB {

	var host = os.Getenv("DB_HOST")
	var database = os.Getenv("DB_NAME")
	var password = os.Getenv("DB_PASSWORD")
	var user = os.Getenv("DB_USER")

	_, envExist := os.LookupEnv("DB_HOST")
	_, envDatabase := os.LookupEnv("DB_NAME")
	_, envPassword := os.LookupEnv("DB_PASSWORD")
	_, envUser := os.LookupEnv("DB_USER")
	_, envPort := os.LookupEnv("DB_PORT")

	if !envExist || !envDatabase || !envPassword || !envUser || !envPort {
		return MongoDB {
			Host: environment.DATABASE_HOST,
			Database: environment.DATABASE_NAME,
			Password: environment.DATABASE_PASSWORD,
			User: environment.DATABASE_USER,
			Port: environment.DATABASE_PORT,
		}
	}

	newMongoDB := MongoDB{Host: host, Database: database, Password: password, User: user}

	return newMongoDB
}

var Mongo = getFromEnvIfExists()
