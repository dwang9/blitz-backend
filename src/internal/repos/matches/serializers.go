package matches

import (
    "github.com/derek-test/blitz-backend/src/internal/entity"
)

func serializeMatchToEntity(m *getMatchResp) *entity.Match {
    return &entity.Match{
        ParticipantsPuuids: m.MetaData.Participants,
        Participants:       serializeMatchParticipantsToEntity(m.Info.Participants),
        Teams:              serializeMatchTeamsToEntity(m.Info.Teams),
        GameCreation:       m.Info.GameCreation,
        QueueID:            m.Info.QueueID,
    }
}
func serializeMatchParticipantsToEntity(p []*getMatchRespParticipant) []*entity.MatchParticipant {
    res := []*entity.MatchParticipant{}
    for i := range p {
        tempNew := serializeMatchParticipantToEntity(p[i])
        res = append(res, tempNew)
    }

    return res
}

func serializeMatchParticipantToEntity(p *getMatchRespParticipant) *entity.MatchParticipant {
    return &entity.MatchParticipant{
        ParticipantId: p.ParticipantId,
        TeamID:        p.TeamID,
        Assists:       p.Assists,
        ChampionName:  p.ChampionName,
        Kills:         p.Kills,
        Deaths:        p.Deaths,
        GoldEarned:    p.GoldEarned,
        VisionScore:   p.VisionScore,
    }
}

func serializeMatchTeamsToEntity(p []*getMatchRespTeam) []*entity.MatchTeam {
    res := []*entity.MatchTeam{}
    for i := range p {
        tempNew := serializeMatchTeamToEntity(p[i])
        res = append(res, tempNew)
    }

    return res
}

func serializeMatchTeamToEntity(p *getMatchRespTeam) *entity.MatchTeam {
    return &entity.MatchTeam{
        TeamID: p.TeamID,
        Win:    p.Win,
    }
}
