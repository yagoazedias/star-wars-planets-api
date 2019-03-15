package client

import (
	"encoding/json"
	"fmt"
	"github.com/yagoazedias/star-wars-planets-api/domain"
	"github.com/yagoazedias/star-wars-planets-api/environment"
	"net/http"
	"strings"
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

	searchName := strings.Replace(planet.Name, " ", "%20", -1)

	resp, _ := http.Get(fmt.Sprintf("%s?search=%s", environment.SWAPI_URL, searchName))
	defer resp.Body.Close()
	err := json.NewDecoder(resp.Body).Decode(&swapiResponse)

	if err != nil {
		return 0, err
	}

	return count(swapiResponse), nil
}