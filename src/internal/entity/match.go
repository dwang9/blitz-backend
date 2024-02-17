package entity

type Match struct {
    Participants       []*MatchParticipant
    Teams              []*MatchTeam
    QueueID            int64
    GameCreation       int64
    ParticipantsPuuids []string
}

type MatchParticipant struct {
    ParticipantId int64
    TeamID        int64
    Assists       int64
    ChampionName  string
    Kills         int64
    Deaths        int64
    GoldEarned    int64
    VisionScore   int64
}

type MatchTeam struct {
    TeamID int64
    Win    bool
}
