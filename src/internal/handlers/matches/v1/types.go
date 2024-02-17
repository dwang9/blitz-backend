package v1

import "time"

type getMatchInfoResponse struct {
    Matches []*getMatchInfoMatch `json:"matches"`
}
type getMatchInfoMatch struct {
    MatchDate          *time.Time `json:"match_date"`
    MatchDateTimestamp int64      `json:"match_date_timestamp"`
    GameType           string     `json:"game_type"`
    ChampionPlayed     string     `json:"champion_played"`
    DidWin             bool       `json:"did_win"`
    Kills              int64      `json:"kills"`
    Deaths             int64      `json:"deaths"`
    Assists            int64      `json:"assists"`
    GoldEarned         int64      `json:"gold_earned"`
    VisionScored       int64      `json:"vision_scored"`
}
