package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	cache "github.com/derek-test/blitz-backend/src/internal/cache/dragonfly"
	"github.com/derek-test/blitz-backend/src/internal/config"
	"github.com/derek-test/blitz-backend/src/internal/connection/dragonfly"
	matchesV1 "github.com/derek-test/blitz-backend/src/internal/handlers/matches/v1"
	summonerV1 "github.com/derek-test/blitz-backend/src/internal/handlers/summoner/v1"
	"github.com/derek-test/blitz-backend/src/internal/repos/matches"
	"github.com/derek-test/blitz-backend/src/internal/repos/summoner"
	"github.com/derek-test/blitz-backend/src/internal/server"
	matchesService "github.com/derek-test/blitz-backend/src/internal/service/matches"
	summonerService "github.com/derek-test/blitz-backend/src/internal/service/summoner"
)

func main() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM)

	cfg := config.GetServiceConfig()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	cacheConn := dragonfly.GetDragonflyConnection(cfg)

	dragonflyCache := cache.NewDragonflyCache(cacheConn)

	summonerRepo := summoner.NewSummonerRepo(cfg)
	matchesRepo := matches.NewMatchesRepo(cfg)

	summonerService := summonerService.NewSummonerService(summonerRepo, dragonflyCache)
	matchesService := matchesService.NewMatchService(matchesRepo, summonerRepo, dragonflyCache)
	summonerServiceHandler := summonerV1.NewSummonerServiceHandler(summonerService)
	matchesServiceHandler := matchesV1.NewMatchHandler(matchesService)

	router := server.NewHttpServer(summonerServiceHandler, matchesServiceHandler)

	if err := http.ListenAndServe(":8080", router); err != nil && err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("unable to start http server")
	}

	log.Info().Msg("work done; exiting")

	os.Exit(0)
}
