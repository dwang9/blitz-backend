package matches

import (
	"context"

	"github.com/derek-test/blitz-backend/src/internal/entity"
)

type MatchService interface {
	GetPlayerMatchHistory(ctx context.Context, summonerName string, endTimestamp int, matchType string, count int) ([]*entity.PlayerMatch, error)
}
