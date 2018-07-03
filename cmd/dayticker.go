package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

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
		var er error

		client := &http.Client{}
		req, _ = http.NewRequest("GET", "https://api.binance.com/api/v1/ticker/24hr", nil)
		q := req.URL.Query()
		resp, er = client.Do(req)
		q.Add("symbol", args[0])
		req.URL.RawQuery = q.Encode()
		resp, er = client.Do(req)

		if er != nil {
			fmt.Printf("The HTTP request failed with error %s\n", er)
		} else {
			data, _ := ioutil.ReadAll(resp.Body)
			var ticker DayTicker
			json.Unmarshal(data, &ticker)
			fmt.Printf("%+v\n", ticker)
		}
	},
}

func init() {
	RootCmd.AddCommand(dayTickerCmd)
}
