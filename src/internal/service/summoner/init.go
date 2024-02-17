package summoner

import (
	"github.com/derek-test/blitz-backend/src/internal/cache"
	"github.com/derek-test/blitz-backend/src/internal/repos/summoner"
)

type service struct {
	repos summoner.SummonerRepo
	cache cache.Cache
}

func NewSummonerService(repos summoner.SummonerRepo, cache cache.Cache) SummonerService {
	return &service{
		repos: repos,
		cache: cache,
	}
}
