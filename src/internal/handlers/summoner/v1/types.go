package v1

type getSummonerInfoResponse struct {
    ID           string `json:"id"`
    Level        int64  `json:"level"`
    SummonerName string `json:"summoner_name"`
    ProfileIcon  string `json:"profile_icon"`
    Rank         string `json:"rank"`
}
