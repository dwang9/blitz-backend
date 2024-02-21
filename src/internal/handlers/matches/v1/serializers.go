package v1

import (
    "github.com/derek-test/blitz-backend/src/internal/entity"
)

func serializePlayerMatchToGetMatchHistoryResp(matches []*entity.PlayerMatch) *getMatchInfoResponse {
    return &getMatchInfoResponse{
        Matches: serializEntityToGetMatchInfoMatches(matches),
    }
}
func serializEntityToGetMatchInfoMatches(matches []*entity.PlayerMatch) []*getMatchInfoMatch {
    res := []*getMatchInfoMatch{}
    for i := range matches {
        tempNew := serializEntityToGetMatchInfoMatch(matches[i])
        res = append(res, tempNew)
    }

    return res
}
func serializEntityToGetMatchInfoMatch(m *entity.PlayerMatch) *getMatchInfoMatch {
    return &getMatchInfoMatch{
        MatchDate:          m.MatchDate,
        MatchDateTimestamp: m.MatchDate.UTC().Unix(),
        GameType:           m.GameType,
        ChampionPlayed:     m.ChampionPlayed,
        DidWin:             m.DidWin,
        Kills:              m.Kills,
        Deaths:             m.Deaths,
        Assists:            m.Assists,
        GoldEarned:         m.GoldEarned,
        VisionScored:       m.VisionScored,
    }
}
