package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"golang.org/x/exp/slices"
)

var validMatchTypeString = []string{"ranked", "normal", "tourney", "tutorial"}

func (c *handler) GetMatchHistory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	summonerNameParam := params["summonerName"]
	if summonerNameParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Invalid summoner name query param!"))
		return
	}

	endTimeParam := 0
	if r.URL.Query().Has("endTime") {
		var err error
		endTimeParam, err = strconv.Atoi(r.URL.Query().Get("endTime"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("500 - endTime query param is not integer"))
			return
		}
		if endTimeParam < 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - endTimeParam has to be between greater than 0"))
			return
		}
	}

	// TO EXPLAIN this needs to be reworked to use queue ids
	// The current implementation will not give the correct results
	matchTypeParam := ""
	if r.URL.Query().Has("matchType") {
		matchTypeParam = r.URL.Query().Get("matchType")
		if !slices.Contains(validMatchTypeString, matchTypeParam) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - matchType has to be [ranked, normal, tourney, tutorial]"))
		}
	}

	countParam := 0
	if r.URL.Query().Has("count") {
		var err error
		countParam, err = strconv.Atoi(r.URL.Query().Get("count"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("500 - count query param is not integer"))
			return
		}
		if countParam < 0 || countParam > 100 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - countParam has to be between 0 and 100"))
			return
		}
	}

	matchHistory, err := c.matchesService.GetPlayerMatchHistory(r.Context(), summonerNameParam, endTimeParam, matchTypeParam, countParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("500 - %s", err.Error())))
		return
	}

	resp := serializePlayerMatchToGetMatchHistoryResp(matchHistory)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
