package server

import (
	matchesV1 "github.com/derek-test/blitz-backend/src/internal/handlers/matches/v1"
	summonerV1 "github.com/derek-test/blitz-backend/src/internal/handlers/summoner/v1"
	"github.com/gorilla/mux"
)

func NewHttpServer(
	summonerHandler summonerV1.SummonerHandler,
	matchesHandler matchesV1.MatchesHandler,
) *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/v1/summoner/{summonerName}", summonerHandler.GetSummonerInfo)
	r.HandleFunc("/v1/summoner/{summonerName}/matches", matchesHandler.GetMatchHistory)

	return r
}
