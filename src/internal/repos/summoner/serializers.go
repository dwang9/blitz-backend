package summoner

import "github.com/derek-test/blitz-backend/src/internal/entity"

func serializeSummonerToEntity(s *getSummonerResp) *entity.Summoner {
    return &entity.Summoner{
        ID:            s.ID,
        AccountID:     s.AccountID,
        Puuid:         s.Puuid,
        Name:          s.Name,
        ProfileIconID: s.ProfileIconID,
        RevisionDate:  s.RevisionDate,
        SummonerLevel: s.SummonerLevel,
    }
}

func serializePlayerDataToEntity(s *playerDataResponse) *entity.PlayerData {
    if s.TotalPoints == nil {
        return &entity.PlayerData{}
    }
    return &entity.PlayerData{
        Level:      s.TotalPoints.Level,
        Current:    s.TotalPoints.Current,
        Max:        s.TotalPoints.Max,
        Percentile: s.TotalPoints.Percentile,
    }
}
