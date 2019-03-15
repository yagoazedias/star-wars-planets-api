package repository

import (
	"fmt"
	"github.com/yagoazedias/star-wars-planets-api/domain"
	"github.com/yagoazedias/star-wars-planets-api/helpers"
	"gopkg.in/mgo.v2/bson"
)

type Planet struct {}

func (*Planet) Search(offset int, limit int) ([]domain.Planet, error) {
	var planets []domain.Planet
	var d = domain.Planet{}

	Mongo.Connect()

	c := Mongo.db.C(d.CollectionName())
	err := c.Find(bson.M{}).Skip(offset).Limit(limit).All(&planets)

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

func (*Planet) Update(planet domain.Planet) (*domain.Planet, error) {
	var updatedPlanet domain.Planet

	Mongo.Connect()
	c := Mongo.db.C(planet.CollectionName())
	err := c.Find(planet.Me()).One(&planet)

	if err != nil {
		return nil, helpers.NewError("Planet does not exists")
	}

	updatedPlanet = domain.Planet{
		ID: planet.ID,
		Name: planet.Name,
		Terrain: planet.Terrain,
		Weather: planet.Weather,
	}

	err = c.Update(planet.Me(), updatedPlanet.ToBson())

	if err != nil {
		return nil, err
	}

	return &planet, nil
}

func (*Planet) Delete(id string) error {
	var planet domain.Planet

	Mongo.Connect()
	c := Mongo.db.C(planet.CollectionName())
	err := c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})

	if err != nil {
		fmt.Printf("Error on remove planet %s", id)
		return err
	}

	return nil
}