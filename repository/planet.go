package repository

import (
	"github.com/yagoazedias/star-wars-planets-api/domain"
	"github.com/yagoazedias/star-wars-planets-api/helpers"
	"gopkg.in/mgo.v2/bson"
)

type Planet struct {}

func (*Planet) Search() ([]domain.Planet, error) {
	var planets []domain.Planet
	var d = domain.Planet{}

	Mongo.Connect()

	c := Mongo.db.C(d.CollectionName())
	err := c.Find(bson.M{}).All(&planets)

	if err != nil {
		return nil, err
	}

	return planets, nil
}

func (*Planet) Lookup(id bson.ObjectId) (*domain.Planet, error) {
	var planet domain.Planet

	Mongo.Connect()

	c := Mongo.db.C(planet.CollectionName())
	err := c.Find(bson.M{"_id": id}).One(&planet)

	if err != nil {
		return nil, err
	}

	return &planet, nil
}

func (*Planet) Create(newPlanet domain.CreatePlanet) (*domain.Planet, error) {
	var planet domain.Planet

	Mongo.Connect()
	c := Mongo.db.C(planet.CollectionName())

	// Check if planet already exist
	count, err := c.Find(newPlanet.Me()).Count()

	if count > 0 {
		return nil, helpers.NewError("Planet already exists")
	}

	err = c.Insert(newPlanet.ToBson())
	err = c.Find(newPlanet.Me()).One(&planet)

	if err != nil {
		return nil, err
	}

	return &planet, nil
}