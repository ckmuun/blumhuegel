package main

/*
	Main Application to fetch data from Finnhub.io and publish it into apache pulsar
*/

import (
	"finnhub-datagen/cordClient"
	"finnhub-datagen/finnhubConn"
	"finnhub-datagen/pulsarConn"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
)

var symbolsToQuery []string

// init pulsar client
func init() {
	initViperConfig()
	verifyFinnhubApiKey()
	pulsarConn.InitPulsarClientInstance(viper.GetString("PULSAR_URL"))

}

func initViperConfig() {
	viper.SetEnvPrefix("BLUM")
	_ = viper.BindEnv("CORDSVC_URL")
	_ = viper.BindEnv("PULSAR_URL")
	_ = viper.BindEnv("FINNHUB_API_KEY")
}

func verifyFinnhubApiKey() {
	log.Print("verifying presence of Finnhub API Key")
	financials, err := finnhubConn.GetBasicFinancials("MSFT", "margin")

	if err != nil {
		panic("can not fetch from finnhub, probably no or wrong finnhub api key supplied")
	}

	log.Print(financials)
}

func getAndPublishQuote2Pulsar(symbol string) {
	log.Print("getting quote")
	quote := finnhubConn.GetTimeSensitiveData(symbol)
	log.Print("getting pulsar producer")
	producer := *pulsarConn.GetProducer("blumhuegel/findata/quotes")

	log.Print("sending quote to pulsar")
	msgId, err := producer.Send(
		context.Background(),
		&pulsar.ProducerMessage{
			Key:   symbol,
			Value: quote,
		},
	)
	if err != nil {
		log.Err(err)
	}
	log.Print("sent msgId to pulsar: ", msgId)
}

func main() {

	finnhubConn.InitFinnhubClient()

	log.Print("Getting list of stock symbols the service should query for")
	symbolsToQuery = cordClient.GetSymbolShorthandsToQuery()

	for index := range symbolsToQuery {
		symbol := symbolsToQuery[index]
		log.Print("fetching data for symbol ", symbol)
		go getAndPublishQuote2Pulsar(symbol)
	}

	router := setupRouter()
	router.Run(":7078")
}

/*
	TODO extend router functionality to, for example, trigger refreshes of the quotes
	TODO add health / metrics endpoints
*/
func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})

	return router

}

/*
	//Stock candles


	// Example with required parameters
	news, _, err := finnhubClient.CompanyNews(auth, "AAPL", "2020-05-01", "2020-05-01")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", news)

	// Example with required and optional parameters
	investorsOwnershipOpts := &finnhub.InvestorsOwnershipOpts{Limit: optional.NewInt64(10)}
	ownerships, _, err := finnhubClient.InvestorsOwnership(auth, "AAPL", investorsOwnershipOpts)
	fmt.Printf("%+v\n", ownerships)

	// Aggregate Indicator
	aggregateIndicator, _, err := finnhubClient.AggregateIndicator(auth, "AAPL", "D")
	fmt.Printf("%+v\n", aggregateIndicator)



	// Company earnings
	earningsSurprises, _, err := finnhubClient.CompanyEarnings(auth, "AAPL", nil)
	fmt.Printf("%+v\n", earningsSurprises)

	// Company EPS estimates
	epsEstimate, _, err := finnhubClient.CompanyEpsEstimates(auth, "AAPL", nil)
	fmt.Printf("%+v\n", epsEstimate)

	// Company executive
	executive, _, err := finnhubClient.CompanyExecutive(auth, "AAPL")
	fmt.Printf("%+v\n", executive)

	// Company peers
	peers, _, err := finnhubClient.CompanyPeers(auth, "AAPL")
	fmt.Printf("%+v\n", peers)

	// Company profile
	profile, _, err := finnhubClient.CompanyProfile(auth, &finnhub.CompanyProfileOpts{Symbol: optional.NewString("AAPL")})
	fmt.Printf("%+v\n", profile)
	profileISIN, _, err := finnhubClient.CompanyProfile(auth, &finnhub.CompanyProfileOpts{Isin: optional.NewString("US0378331005")})
	fmt.Printf("%+v\n", profileISIN)
	profileCusip, _, err := finnhubClient.CompanyProfile(auth, &finnhub.CompanyProfileOpts{Cusip: optional.NewString("037833100")})
	fmt.Printf("%+v\n", profileCusip)

	// Company profile2
	profile2, _, err := finnhubClient.CompanyProfile2(auth, &finnhub.CompanyProfile2Opts{Symbol: optional.NewString("AAPL")})
	fmt.Printf("%+v\n", profile2)

	// Revenue Estimates
	revenueEstimates, _, err := finnhubClient.CompanyRevenueEstimates(auth, "AAPL", nil)
	fmt.Printf("%+v\n", revenueEstimates)

	// List country
	countries, _, err := finnhubClient.Country(auth)
	fmt.Printf("%+v\n", countries)

	// Covid-19
	covid19, _, err := finnhubClient.Covid19(auth)
	fmt.Printf("%+v\n", covid19)

	// Crypto candles
	cryptoCandles, _, err := finnhubClient.CryptoCandles(auth, "BINANCE:BTCUSDT", "D", 1590988249, 1591852249)
	fmt.Printf("%+v\n", cryptoCandles)

	// Crypto exchanges
	cryptoExchange, _, err := finnhubClient.CryptoExchanges(auth)
	fmt.Printf("%+v\n", cryptoExchange)

	// Crypto symbols
	cryptoSymbol, _, err := finnhubClient.CryptoSymbols(auth, "BINANCE")
	fmt.Printf("%+v\n", cryptoSymbol[0:5])

	// Earnings calendar
	earningsCalendar, _, err := finnhubClient.EarningsCalendar(auth, &finnhub.EarningsCalendarOpts{
		From: optional.NewString("2020-06-12"), To: optional.NewString("2020-06-20")})
	fmt.Printf("%+v\n", earningsCalendar)

	// Economic code
	economicCode, _, err := finnhubClient.EconomicCode(auth)
	fmt.Printf("%+v\n", economicCode)

	// Economic data
	economicData, _, err := finnhubClient.EconomicData(auth, "MA-USA-656880")
	fmt.Printf("%+v\n", economicData)

	// Filings
	filings, _, err := finnhubClient.Filings(auth, &finnhub.FilingsOpts{Symbol: optional.NewString("AAPL")})
	fmt.Printf("%+v\n", filings)

	// Financials
	financials, _, err := finnhubClient.Financials(auth, "AAPL", "bs", "annual")
	fmt.Printf("%+v\n", financials)

	// Financials Reported
	financialsReported, _, err := finnhubClient.FinancialsReported(auth, &finnhub.FinancialsReportedOpts{Symbol: optional.NewString("AAPL")})
	fmt.Printf("%+v\n", financialsReported)

	// Forex candles
	forexCandles, _, err := finnhubClient.ForexCandles(auth, "OANDA:EUR_USD", "D", 1590988249, 1591852249)
	fmt.Printf("%+v\n", forexCandles)

	// Forex exchanges
	forexExchanges, _, err := finnhubClient.ForexExchanges(auth)
	fmt.Printf("%+v\n", forexExchanges)
	// Forex rates
	forexRates, _, err := finnhubClient.ForexRates(auth, nil)
	fmt.Printf("%+v\n", forexRates)

	// Forex symbols
	forexSymbols, _, err := finnhubClient.ForexSymbols(auth, "OANDA")
	fmt.Printf("%+v\n", forexSymbols)

	// Fund ownership
	fundOwnership, _, err := finnhubClient.FundOwnership(auth, "AAPL", nil)
	fmt.Printf("%+v\n", fundOwnership)

	// General news
	generalNews, _, err := finnhubClient.GeneralNews(auth, "general", nil)
	fmt.Printf("%+v\n", generalNews)

	// Ipo calendar
	ipoCalendar, _, err := finnhubClient.IpoCalendar(auth, "2020-01-01", "2020-06-15")
	fmt.Printf("%+v\n", ipoCalendar)

	// Major development
	majorDevelopment, _, err := finnhubClient.MajorDevelopments(auth, "AAPL", nil)
	fmt.Printf("%+v\n", majorDevelopment)

	// News sentiment
	newsSentiment, _, err := finnhubClient.NewsSentiment(auth, "AAPL")
	fmt.Printf("%+v\n", newsSentiment)

	// Pattern recognition
	patterns, _, err := finnhubClient.PatternRecognition(auth, "AAPL", "D")
	fmt.Printf("%+v\n", patterns)

	// Price target
	priceTarget, _, err := finnhubClient.PriceTarget(auth, "AAPL")
	fmt.Printf("%+v\n", priceTarget)

	// Quote
	quote, _, err := finnhubClient.Quote(auth, "AAPL")
	fmt.Printf("%+v\n", quote)

	// Recommendation trends
	recommendationTrend, _, err := finnhubClient.RecommendationTrends(auth, "AAPL")
	fmt.Printf("%+v\n", recommendationTrend)

	// Stock dividens
	dividends, _, err := finnhubClient.StockDividends(auth, "KO", "2019-01-01", "2020-06-30")
	fmt.Printf("%+v\n", dividends)

	// Splits
	splits, _, err := finnhubClient.StockSplits(auth, "AAPL", "2000-01-01", "2020-06-15")
	fmt.Printf("%+v\n", splits)

	// Stock symbols
	stockSymbols, _, err := finnhubClient.StockSymbols(auth, "US")
	fmt.Printf("%+v\n", stockSymbols[0:5])

	// Support resistance
	supportResitance, _, err := finnhubClient.SupportResistance(auth, "AAPL", "D")
	fmt.Printf("%+v\n", supportResitance)

	// Technical indicator
	technicalIndicator, _, err := finnhubClient.TechnicalIndicator(auth, "AAPL", "D", 1583098857, 1584308457, "sma", &finnhub.TechnicalIndicatorOpts{
		IndicatorFields: map[string]interface{}{
			"timeperiod": 3,
		},
	})
	fmt.Printf("%+v\n", technicalIndicator)

	// Transcripts
	transcripts, _, err := finnhubClient.Transcripts(auth, "AAPL_162777")
	fmt.Printf("%+v\n", transcripts)

	// Transcripts list
	transcriptsList, _, err := finnhubClient.TranscriptsList(auth, "AAPL")
	fmt.Printf("%+v\n", transcriptsList)

	// Upgrade/downgrade
	upgradeDowngrade, _, err := finnhubClient.UpgradeDowngrade(auth, &finnhub.UpgradeDowngradeOpts{Symbol: optional.NewString("BYND")})
	fmt.Printf("%+v\n", upgradeDowngrade)

	// Tick Data
	tickData, _, err := finnhubClient.StockTick(auth, "AAPL", "2020-03-25", 500, 0)
	fmt.Printf("%+v\n", tickData)

	// Indices Constituents
	indicesConstData, _, err := finnhubClient.IndicesConstituents(auth, "^GSPC")
	fmt.Printf("%+v\n", indicesConstData)

	// Indices Historical Constituents
	indicesHistoricalConstData, _, err := finnhubClient.IndicesHistoricalConstituents(auth, "^GSPC")
	fmt.Printf("%+v\n", indicesHistoricalConstData)

	// ETFs Profile
	etfsProfileData, _, err := finnhubClient.EtfsProfile(auth, "SPY")
	fmt.Printf("%+v\n", etfsProfileData)

	// ETFs Holdings
	etfsHoldingsData, _, err := finnhubClient.EtfsHoldings(auth, "SPY")
	fmt.Printf("%+v\n", etfsHoldingsData)

	// ETFs Industry Exposure
	etfsIndustryExposureData, _, err := finnhubClient.EtfsIndustryExposure(auth, "SPY")
	fmt.Printf("%+v\n", etfsIndustryExposureData)

	// ETFs Country Exposure
	etfsCountryExposureData, _, err := finnhubClient.EtfsCountryExposure(auth, "SPY")
	fmt.Printf("%+v\n", etfsCountryExposureData)


}

*/
