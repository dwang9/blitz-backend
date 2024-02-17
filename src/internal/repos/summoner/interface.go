package summoner

import (
	"context"

	"github.com/derek-test/blitz-backend/src/internal/entity"
)

//go:generate mockgen -source interface.go -destination mocks/mock.go -mock_names SummonerRepo=MockSummonerRepo
type SummonerRepo interface {
	GetSummonerInfo(ctx context.Context, summonerName string) (*entity.Summoner, error)
	GetPlayerData(ctx context.Context, puuuid string) (*entity.PlayerData, error)
}
