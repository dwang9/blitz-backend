package summoner

import (
	"context"

	"github.com/derek-test/blitz-backend/src/internal/entity"
)

func (s *service) GetSummonerInfo(ctx context.Context, summonerName string) (*entity.Summoner, *entity.PlayerData, error) {
	summoner, err := s.repos.GetSummonerInfo(ctx, summonerName)
	if err != nil {
		return nil, nil, err
	}
	playerData, err := s.repos.GetPlayerData(ctx, summoner.Puuid)
	if err != nil {
		return nil, nil, err
	}

	return summoner, playerData, nil
}
