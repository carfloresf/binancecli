package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// DayTicker is used to output a standard structure (marshalled as JSON)
type DayTicker struct {
	Symbol             string `symbol:"symbol"`
	PriceChange        string `priceChange:"priceChange"`
	PriceChangePercent string `priceChangePercent:"priceChangePercent"`
	WeightedAvgPrice   string `weightedAvgPrice:"weightedAvgPrice"`
	PrevClosePrice     string `prevClosePrice:"prevClosePrice"`
	LastPrice          string `lastPrice:"lastPrice"`
	LastQty            string `lastQty:"lastQty"`
	BidPrice           string `bidPrice:"bidPrice"`
	AskPrice           string `askPrice:"askPrice"`
	PpenPrice          string `openPrice:"openPrice"`
	HighPrice          string `highPrice:"highPrice"`
	LowPrice           string `lowPrice:"lowPrice"`
	Volume             string `volume:"volume"`
	QuoteVolume        string `quoteVolume:"quoteVolume"`
	PpenTime           string `openTime:"openTime"`
	CloseTime          string `closeTime:"closeTime"`
	FristID            string `fristId:"fristId"`
	LastID             string `lastId:"lastId"`
	Count              string `count:"count"`
}

var dayTickerCmd = &cobra.Command{
	Use:   "dayTicker",
	Short: "Get Day Ticker from Binance for a symbol",
	Run: func(cmd *cobra.Command, args []string) {
		var resp *http.Response
		var req *http.Request
		var errr error
		var er1 error

		client := &http.Client{}
		req, errr = http.NewRequest("GET", "https://api.binance.com/api/v1/ticker/24hr", nil)
		if errr != nil {
			log.Fatalf("Problems: %s", errr.Error())
		}

		q := req.URL.Query()

		q.Add("symbol", args[0])
		req.URL.RawQuery = q.Encode()
		resp, er1 = client.Do(req)

		if er1 != nil {
			log.Fatalf("The HTTP request failed with error %s\n", er1)
		} else {
			data, er2 := ioutil.ReadAll(resp.Body)
			if er2 != nil {
				log.Fatalf("Problems during unmarshalling: %s", er2.Error())
			}

			var ticker DayTicker
			err := json.Unmarshal(data, &ticker)
			if err != nil {
				log.Fatalf("Problems during unmarshalling: %s", err.Error())
			}
			fmt.Printf("%+v\n", ticker)
		}
	},
}

func init() {
	RootCmd.AddCommand(dayTickerCmd)
}
