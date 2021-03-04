package cordClient

import (
	"github.com/magiconair/properties/assert"
	"github.com/rs/zerolog/log"
	"testing"
)

func TestGetDemoSymbols(t *testing.T) {
	symbols := GetDemoSymbols()

	log.Print(symbols)

	assert.Equal(t, len(symbols), 11)
}

func TestGetSymbolShorthandsToQuery(t *testing.T) {

}
