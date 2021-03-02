package finnhubConn

import (
	"github.com/Finnhub-Stock-API/finnhub-go"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"sync"
)

var once sync.Once
var finnhubClient *finnhub.DefaultApiService

func InitFinnhubClient() *finnhub.DefaultApiService {

	log.Print("creating Pulsar Client")
	once.Do(func() {
		finnhubClient = finnhub.NewAPIClient(finnhub.NewConfiguration()).DefaultApi
	})
	return finnhubClient
}

func GetBasicFinancials(symbol string, metric string) (finnhub.BasicFinancials, error) {
	// Basic financials

	basicFinancials, _, err := finnhubClient.CompanyBasicFinancials(getFinnhubAuth(), symbol, "margin")
	return basicFinancials, err
}

func getFinnhubAuth() context.Context {
	apikey := viper.Get("API_KEY")
	auth := context.WithValue(context.Background(), finnhub.ContextAPIKey, finnhub.APIKey{
		Key: apikey.(string),
	})
	return auth

}

// OLD file-based configuration of API key
/*
func getFinnhubAuth() (error, context.Context) {
	file, err := os.Open("../api-key.txt")

	if err != nil {
		log.Fatal()
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal()
		}
	}()

	scanner := bufio.NewScanner(file)

	auth := context.WithValue(context.Background(), finnhub.ContextAPIKey, finnhub.APIKey{
		Key: scanner.Text(),
	})
	return err, auth
}

*/
