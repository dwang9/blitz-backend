package summoner

type getSummonerResp struct {
	ID            string `json:"id"`
	AccountID     string `json:"account_id"`
	Puuid         string `json:"puuid"`
	Name          string `json:"name"`
	ProfileIconID int64  `json:"profileIconId"`
	RevisionDate  int64  `json:"revisionDate"`
	SummonerLevel int64  `json:"summonerLevel"`
}

type playerDataResponse struct {
	TotalPoints *playerDataResponseTotalPoints `json:"totalPoints"`
}

type playerDataResponseTotalPoints struct {
	Level      string  `json:"level"`
	Current    int64   `json:"current"`
	Max        int64   `json:"max"`
	Percentile float64 `json:"percentile"`
}
