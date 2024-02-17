package v1

type getSummonerInfoResponse struct {
    ID           string `json:"id"`
    SummonerName string `json:"summoner_name"`
    ProfileIcon  string `json:"profile_icon"`
    Rank         string `json:"rank"`
    MasteryScore int64  `json:"mastery_score"`
}
