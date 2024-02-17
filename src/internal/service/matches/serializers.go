package matches

import (
    "time"

    "github.com/derek-test/blitz-backend/src/internal/entity"
)

func serializeMatchToPlayerMatch(m *entity.Match, p *entity.MatchParticipant, t *entity.MatchTeam, gameType string) *entity.PlayerMatch {
    gameCreation := time.Unix(m.GameCreation/1000, 0)
    return &entity.PlayerMatch{
        MatchDate:      &gameCreation,
        GameType:       gameType,
        ChampionPlayed: p.ChampionName,
        DidWin:         t.Win,
        Kills:          p.Kills,
        Deaths:         p.Deaths,
        Assists:        p.Assists,
        GoldEarned:     p.GoldEarned,
        VisionScored:   p.VisionScore,
    }
}
