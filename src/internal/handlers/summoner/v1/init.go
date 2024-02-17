package v1

import (
	"github.com/derek-test/blitz-backend/src/internal/service/summoner"
)

type handler struct {
	summonerService summoner.SummonerService
}

func NewSummonerServiceHandler(summonerService summoner.SummonerService) SummonerHandler {
	return &handler{
		summonerService: summonerService,
	}
}
