package finnhubConn

import (
	"github.com/Finnhub-Stock-API/finnhub-go"
	"github.com/rs/zerolog/log"
	"sync"
)

var once sync.Once
var finnhubClientSingleton *finnhub.DefaultApiService

func GetFinnhubClient() *finnhub.DefaultApiService {

	log.Print("creating Pulsar Client")
	once.Do(func() {
		finnhubClientSingleton = finnhub.NewAPIClient(finnhub.NewConfiguration()).DefaultApi
	})
	return finnhubClientSingleton

}
