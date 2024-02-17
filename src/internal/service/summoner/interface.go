package summoner

import (
	"context"

	"github.com/derek-test/blitz-backend/src/internal/entity"
)

type SummonerService interface {
	GetSummonerInfo(ctx context.Context, summonerName string) (*entity.Summoner, *entity.PlayerData, error)
}
