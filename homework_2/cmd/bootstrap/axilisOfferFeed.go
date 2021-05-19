package bootstrap

import (
	stdhttp "net/http"

	"code-cadets-2021/homework_2/internal/infrastructure/http"
)

func AxilisOfferFeed() *http.AxilisOfferFeed {
	httpClient := &stdhttp.Client{}
	return http.NewAxilisOfferFeed(httpClient)
}
