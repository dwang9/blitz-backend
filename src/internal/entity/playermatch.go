package entity

import "time"

type PlayerMatch struct {
    MatchDate      *time.Time
    GameType       string
    ChampionPlayed string
    DidWin         bool
    Kills          int64
    Deaths         int64
    Assists        int64
    GoldEarned     int64
    VisionScored   int64
}
