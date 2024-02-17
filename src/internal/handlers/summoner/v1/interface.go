package v1

import "net/http"

type SummonerHandler interface {
    GetSummonerInfo(w http.ResponseWriter, r *http.Request)
}
