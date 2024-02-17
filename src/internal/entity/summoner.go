package entity

import (
	"fmt"
)

type Summoner struct {
	ID            string
	AccountID     string
	Puuid         string
	Name          string
	ProfileIconID int64
	RevisionDate  int64
	SummonerLevel int64
}

func (s *Summoner) GetProfileIconImageURL() string {
	return fmt.Sprintf("https://ddragon.leagueoflegends.com/cdn/14.3.1/img/profileicon/%d.png", s.ProfileIconID)
}
