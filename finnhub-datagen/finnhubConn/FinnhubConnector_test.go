package finnhubConn

import (
	"github.com/rs/zerolog/log"
	"testing"
)

func init() {
	InitFinnhubClient()
}

func TestBasicCompanyFinanacials(t *testing.T) {
	log.Print("fetching basic financial from finnhub")

	GetBasicFinancials()
	//fmt.Printf("%+v\n", stockCandles)
}
