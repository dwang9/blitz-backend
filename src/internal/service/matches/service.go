package matches

import (
	"context"
	"fmt"

	"github.com/derek-test/blitz-backend/src/internal/entity"
	"golang.org/x/time/rate"
)

var queueIdToGameTypeMap = map[int64]string{
	400: "Summoner's Rift - 5v5 Draft Pick games",
	420: "Summoner's Rift - 5v5 Ranked Solo games",
	430: "Summoner's Rift - 5v5 Blind pick games",
	440: "Summoner's Rift - 5v5 Ranked flex games",
	450: "Howling Abyss - 5v5 ARAM games",
}

func (s *service) GetPlayerMatchHistory(ctx context.Context, summonerName string, endTimestamp int, matchType string, count int) ([]*entity.PlayerMatch, error) {
	// TOEXPLAIN: this is a spot where we can maybe use caching to make more efficient
	// I think we can either store a link between summoner name and puuid in a cache or
	// in a db and that way this call wont be necessary
	summoner, err := s.summonerRepo.GetSummonerInfo(ctx, summonerName)
	if err != nil {
		return nil, err
	}

	matchHistoryIDs, err := s.matchRepo.GetMatchHistory(ctx, summoner.Puuid, int64(endTimestamp), matchType, count)
	if err != nil {
		return nil, err
	}

	limiter := rate.NewLimiter(20, 1)

	matchHistories := []*entity.Match{}
	for _, matchHistoryID := range matchHistoryIDs {
		limiter.Wait(ctx)

		matchHistory, err := s.matchRepo.GetMatch(ctx, matchHistoryID)
		if err != nil {
			return nil, err
		}
		matchHistories = append(matchHistories, matchHistory)
	}

	playerMatches := []*entity.PlayerMatch{}
	for i := range matchHistories {
		matchHistory := matchHistories[i]
		playerMatch := mapMatchToPlayerMatch(summoner.Puuid, matchHistory)
		playerMatches = append(playerMatches, playerMatch)
	}

	return playerMatches, nil
}

func queueIDToGameType(queueId int64) string {
	if val, ok := queueIdToGameTypeMap[queueId]; ok {
		return val
	}

	return ""
}

func mapMatchToPlayerMatch(puuid string, match *entity.Match) *entity.PlayerMatch {

	participant := getMatchRespParticipantByPuuid(puuid, match)
	team := getMatchRespTeamByTeamID(participant.TeamID, match)

	playerMatch := serializeMatchToPlayerMatch(match, participant, team, queueIDToGameType(match.QueueID))

	return playerMatch
}

func getParticipantId(puuid string, participants []string) int {
	for i, participantId := range participants {
		if participantId == puuid {
			return i + 1
		}

	}

	return -1
}

func getMatchRespParticipantByPuuid(puuid string, resp *entity.Match) *entity.MatchParticipant {
	if resp != nil {

		participantId := getParticipantId(puuid, resp.ParticipantsPuuids)
		for i := range resp.Participants {
			participant := resp.Participants[i]
			if participant.ParticipantId == int64(participantId) {
				return participant
			}

		}
	}

	fmt.Println("Participant with id: " + fmt.Sprintf("%d", puuid) + " is not in participants field")

	return nil
}

func getMatchRespTeamByTeamID(teamId int64, resp *entity.Match) *entity.MatchTeam {
	if resp != nil {
		for i := range resp.Teams {
			team := resp.Teams[i]
			if team.TeamID == teamId {
				return team
			}

		}
	}

	fmt.Println("Team with id: " + fmt.Sprintf("%d", teamId) + " is not in Teams field")

	return nil
}
