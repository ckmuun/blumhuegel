package finnhubConn

import (
	"bufio"
	"fmt"
	"github.com/Finnhub-Stock-API/finnhub-go"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"os"
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

func GetBasicFinancials() (finnhub.BasicFinancials, error) {
	// Basic financials
	err, auth := getFinnhubAuth()
	if nil != err {
		panic("can not authenticate to finnhub api")
	}

	basicFinancials, _, err := finnhubClient.CompanyBasicFinancials(auth, "MSFT", "margin")
	fmt.Printf("%+v\n", basicFinancials)
	return basicFinancials, err
}
