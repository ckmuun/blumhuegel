package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"sort"
)

var usSymbolsFull = make(map[string]AssetSymbol)

var symbols = make(map[string][]AssetSymbol)

// TODO add configuration for file-based port settings, add general means for dynamic configuration
/*
	TODO add support for symbol fetching from
*/
func main() {

	viper.SetDefault("finnhub-api-key", "no-finnhub-api-key-loaded")
	viper.SetEnvPrefix("BLUM")
	_ = viper.BindEnv("FINNHUB_API_KEY")

	log.Println("fetching usExchangeShorthand symbols")
	usSymbolsFull = GetSymbolsAsShorthandMap("US")

	usExchangeShorthand := "us"
	symbols[usExchangeShorthand] = GetFinnhubSymbolsAsArr(usExchangeShorthand)

	router := setupRouter()
	router.Run(":7077")
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})

	router.GET("/short/all", func(c *gin.Context) {
		c.JSON(200, getAssetSymbolShorthands(usSymbolsFull))
	})

	// random demo symbols
	router.GET("/demo", func(c *gin.Context) {
		c.JSON(200, getDemoSymbols())
	})

	// TODO currently this 'real' endpoint also just returns the demo symbols. The magic needs to be implemented that makes this work and return proper symbols.
	router.GET("/selection", func(c *gin.Context) {
		c.JSON(200, getDemoSymbols())
	})

	// exchange is the country code for the requested stock exchange, e.g. "US"
	router.GET("/symbols/:exchange/:field/:fieldValue", func(c *gin.Context) {
		exchange := c.Param("exchange")
		field := c.Param("field")
		fieldValue := c.Param("fieldValue")

		// TODO add validation middleware for these parameters, e.g. allowed exchanges and struct atts / struct values

		subset, err := getStockSymbolSubset(exchange, field, fieldValue)

		// todo return the 400 if the query params are validated as junk , return a 500 if err != nil.
		if nil != err {
			c.JSON(500, err)
		}

		c.JSON(200, subset)
	})

	return router
}

func getDemoSymbols() []string {
	demoSymbols := make([]string, 11)

	demoSymbols[0] = "AAPL"
	demoSymbols[1] = "IBM"
	demoSymbols[2] = "MMM"
	demoSymbols[3] = "GME"
	demoSymbols[4] = "PLXXF"
	demoSymbols[5] = "ASOMF"
	demoSymbols[6] = "IBDD"
	demoSymbols[7] = "HECCU"
	demoSymbols[8] = "NPTVF"
	demoSymbols[9] = "SMAWF"
	demoSymbols[10] = "FDIS"

	return demoSymbols
}

/*

	This function crafts the specific response for a SQL WHERE-like request.
	Basically, the client is asking "give me all stock symbols with this value in this field"
	Think of it as a (very) poor man's query engine.
	It uses a magic index integer that codes the field of AssetSymbol to look at.
	The plan is to move this stuff into postgres with full SQL syntax, that's why this implemenation
	is very rudimentary

	type AssetSymbol struct {
		Currency      string `json:"currency"`
		Description   string `json:"description"`
		DisplaySymbol string `json:"displaySymbol"`
		Figi          string `json:"figi"`
		Mic           string `json:"mic"`
		Symbol        string `json:"symbol"`
		AssetType     string `json:"type"`
}
*/
func getStockSymbolSubset(exchange string, field string, fieldValue string) ([]AssetSymbol, error) {

	excSymbols := symbols[exchange]

	log.Println("selection criteria are: ")
	log.Println("exchange: ", exchange)
	log.Println("fieldToMatch: ", field)
	log.Println("fieldValueToMatch: ", fieldValue)

	switch field {
	case "currency":
		log.Println("select where currency ==", fieldValue)
		return populateReturnSymbolArr(excSymbols, 0, fieldValue), nil
	case "description":
		log.Println("0select where description == ", fieldValue)
		return populateReturnSymbolArr(excSymbols, 1, fieldValue), nil
	case "displaySymbol":
		log.Println("select where displaySymbol == ", fieldValue)
		return populateReturnSymbolArr(excSymbols, 2, fieldValue), nil
	case "figi":
		log.Println("select where figi ==", fieldValue)
		return populateReturnSymbolArr(excSymbols, 3, fieldValue), nil
	case "mic":
		log.Println("select where mic == ", fieldValue)
		return populateReturnSymbolArr(excSymbols, 4, fieldValue), nil
	case "symbol":
		log.Println("select where symbol == ", fieldValue)
		return populateReturnSymbolArr(excSymbols, 5, fieldValue), nil
	case "type":
		log.Println("select where type == ", fieldValue)
		return populateReturnSymbolArr(excSymbols, 6, fieldValue), nil
	}

	return nil, errors.New("can not fulfill asset symbol query with given values")
}

func populateReturnSymbolArr(symbols []AssetSymbol, magicIndex int8, fieldValue string) (subset []AssetSymbol) {

	for index := range symbols {
		symbol := symbols[index]
		if compareAssetSymbolField(magicIndex, symbol, fieldValue) {
			subset = append(subset, symbol)
		}
	}
	return subset
}

/*
	TODO this implementation is not really elegant, with the magic number.
	FIXME add some sort of enum here or make this not-hardcoded or just put the data into postgres
	The @magicIndex is hardcoded value to specify which field of AssetSymbol should be compared to the desired
	field value.

*/
func compareAssetSymbolField(magicIndex int8, symbol AssetSymbol, fieldValue string) bool {

	if magicIndex == 0 && symbol.Currency == fieldValue {
		return true
	}
	if magicIndex == 1 && symbol.Description == fieldValue {
		return true
	}
	if magicIndex == 2 && symbol.DisplaySymbol == fieldValue {
		return true
	}
	if magicIndex == 3 && symbol.Figi == fieldValue {
		return true
	}
	if magicIndex == 4 && symbol.Mic == fieldValue {
		return true
	}
	if magicIndex == 5 && symbol.Symbol == fieldValue {
		return true
	}
	if magicIndex == 6 && symbol.AssetType == fieldValue {
		return true
	}
	return false
}

/*
	This just takes the map keys (the stock symbol shorthands) and puts them into an array.
*/
func getAssetSymbolShorthands(fullSymbols map[string]AssetSymbol) []string {
	log.Println("converting full symbols map to array containing shorthands")
	log.Println("number of symbols:")
	log.Println(len(fullSymbols))

	keys := make([]string, len(fullSymbols))

	i := 0
	for k := range fullSymbols {
		keys[i] = k
		i++
	}
	// Everything in Go is evaluated lazily, but this runs in-place.
	log.Println("sorting shorthands")
	sort.Strings(keys)
	return keys
}
