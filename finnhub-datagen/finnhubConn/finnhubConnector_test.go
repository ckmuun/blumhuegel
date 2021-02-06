package finnhubConn

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	InitFinnhubClient()
}

/*
	TODO explore finnhub api and develop more elaborate scenarios
*/
func TestBasicCompanyFinancials(t *testing.T) {
	log.Print("fetching basic financial from finnhub")

	financials, err := GetBasicFinancials()
	fmt.Printf("%+v\n", financials)

	assert.NoError(t, err)

}
