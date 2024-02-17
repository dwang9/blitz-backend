package matches

type getMatchResp struct {
	Info     *getMatchRespInfo     `json:"info"`
	MetaData *getMatchRespMetaData `json:"metadata"`
}
type getMatchRespInfo struct {
	Participants []*getMatchRespParticipant `json:"participants"`
	Teams        []*getMatchRespTeam        `json:"teams"`
	QueueID      int64                      `json:"queueId"`
	GameCreation int64                      `json:"gameCreation"`
}

type getMatchRespMetaData struct {
	Participants []string `json:"participants"`
}

type getMatchRespParticipant struct {
	ParticipantId int64  `json:"participantId"`
	TeamID        int64  `json:"teamId"`
	Assists       int64  `json:"assists"`
	ChampionName  string `json:"championName"`
	Kills         int64  `json:"kills"`
	Deaths        int64  `json:"deaths"`
	GoldEarned    int64  `json:"goldEarned"`
	VisionScore   int64  `json:"visionScore"`
}

type getMatchRespTeam struct {
	TeamID int64 `json:"teamId"`
	Win    bool  `json:"win"`
}
