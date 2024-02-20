package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (c *handler) GetSummonerInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	summonerNameParam := params["summonerName"]
	if summonerNameParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Invalid summoner name query param!"))
		return
	}

	summonerInfo, playerData, err := c.summonerService.GetSummonerInfo(r.Context(), summonerNameParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("500 - %s", err.Error())))
		return
	}

	resp := &getSummonerInfoResponse{
		ID:           summonerInfo.ID,
		SummonerName: summonerInfo.Name,
		ProfileIcon:  summonerInfo.GetProfileIconImageURL(),
		Rank:         playerData.Level,
		Level:        summonerInfo.SummonerLevel,
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
