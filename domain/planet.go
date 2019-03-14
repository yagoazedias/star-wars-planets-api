package domain

import (
	"github.com/yagoazedias/star-wars-planets-api/helpers"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// The representation of a created planet
type Planet struct {
	ID           bson.ObjectId  `bson:"_id"          json:"id,omitempty"`
	Name         string         `bson:"name"         json:"name,omitempty"`
	Weather      string         `bson:"weather"      json:"weather,omitempty"`
	Terrain      string         `bson:"terrain"      json:"terrain,omitempty"`
	Count        int            `bson:"count"        json:"count"`
	CreatedAt    *time.Time     `bson:"created_at"   json:"created_at,omitempty"`
	UpdatedAt    *time.Time     `bson:"updated_at"   json:"updated_at,omitempty"`
	DeletedAt    *time.Time     `bson:"deleted_at"   json:"deleted_at,omitempty"`
}

// The representation of a potential planet
type CreatePlanet struct {
	Name         string         `bson:"name"         json:"name"`
	Weather      string         `bson:"weather"      json:"weather"`
	Terrain  string             `bson:"terrain"  json:"terrain"`
}

// Warning about timezone issues:
// https://stackoverflow.com/questions/44873825/how-to-get-timestamp-of-utc-time-with-golang
func (c *CreatePlanet) ToBson() bson.M {
	return bson.M{
		"name": c.Name,
		"weather": c.Weather,
		"terrain": c.Terrain,
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}
}

func (c *Planet) ToBson() bson.M {
	return bson.M{
		"name": c.Name,
		"weather": c.Weather,
		"terrain": c.Terrain,
		"created_at": c.CreatedAt,
		"updated_at": time.Now(),
	}
}

func (c *CreatePlanet) Me() bson.M {
	return bson.M{
		"name": c.Name,
	}
}

func (c *Planet) Me() bson.M {
	return bson.M{
		"_id": c.ID,
	}
}

func (c *CreatePlanet) IsValid() (bool, error) {
	if c.Name == "" {
		return false, helpers.NewError("Planets must have a name")
	}

	if c.Terrain == "" {
		return false, helpers.NewError("Planets must have a terrain")
	}

	if c.Weather == "" {
		return false, helpers.NewError("Planets must have a weather")
	}

	return true, nil
}

func (c *Planet) IsValid() (bool, error) {
	if c.Name == "" {
		return false, helpers.NewError("Planets must have a name")
	}

	if c.Terrain == "" {
		return false, helpers.NewError("Planets must have a terrain")
	}

	if c.Weather == "" {
		return false, helpers.NewError("Planets must have a weather")
	}

	return true, nil
}

// Same idea from gorm, when "tableName" implements the gorm.Tabler interface
func (*Planet) CollectionName() string {
	return "planets"
}
