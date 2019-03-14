package client

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/yagoazedias/star-wars-planets-api/domain"
	"github.com/yagoazedias/star-wars-planets-api/environment"
	"net/http"
	"time"
)

type Swapi struct {}

type SwapiResponse struct {
	Results  []SwapiPlanet  `bson:"results"         json:"results,omitempty"`
}

type SwapiPlanet struct {
	Name string 			`bson:"name"            json:"name,omitempty"`
	Films []string 			`bson:"films"           json:"films,omitempty"`
}

func count(s SwapiResponse) int {

	if len(s.Results) > 1 || len(s.Results) == 0 {
		return 0
	}

	return len(s.Results[0].Films)
}

func (*Swapi) GetPlanetAttendance(planet *domain.Planet) (int, error) {

	var swapiResponse SwapiResponse
	val, err := Cache.db.Get(environment.STARTSHIPS_REDIS_KEY).Result()

	if err != nil {
		resp, _ := http.Get(fmt.Sprintf("%s?search=%s", environment.SWAPI_URL, planet.Name))
		defer resp.Body.Close()

		err := json.NewDecoder(resp.Body).Decode(&swapiResponse)

		if err != nil {
			return 0, err
		}

		set, err := Cache.db.Set(environment.STARTSHIPS_REDIS_KEY, []byte(string(json.Marshal(swapiResponse))), 10*time.Second).Result()

		if err != nil {
			fmt.Printf("Was not possible to cache: %s", err)
		} else {
			fmt.Printf("Done: %s", set)
		}

	} else {
		fmt.Printf("%s", val)
	}

	return count(swapiResponse), nil
}