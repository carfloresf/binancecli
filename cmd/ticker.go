package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

type Ticker struct {
	Symbol string `symbol:"symbol"`
	Price  string `price:"price"`
}

var tickerCmd = &cobra.Command{
	Use:   "ticker",
	Short: "Get Tickers from Binance",
	Run: func(cmd *cobra.Command, args []string) {
		var resp *http.Response
		var req *http.Request
		var er error

		client := &http.Client{}
		req, _ = http.NewRequest("GET", "https://api.binance.com/api/v3/ticker/price", nil)
		q := req.URL.Query()

		if len(args) == 0 {
			resp, er = client.Do(req)
			if er != nil {
				fmt.Printf("The HTTP request failed with error %s\n", er)
			} else {
				data, _ := ioutil.ReadAll(resp.Body)
				var tickers []Ticker
				json.Unmarshal(data, &tickers)

				fmt.Println("Found", len(tickers), "pairs in Binance")

				for _, ti := range tickers {
					fmt.Println(ti.Symbol + " - " + ti.Price)
				}
			}
		} else if len(args) == 1 {
			q.Add("symbol", args[0])
			req.URL.RawQuery = q.Encode()
			resp, er = client.Do(req)

			if er != nil {
				fmt.Printf("The HTTP request failed with error %s\n", er)
			} else {
				data, _ := ioutil.ReadAll(resp.Body)
				var ticker Ticker
				json.Unmarshal(data, &ticker)
				fmt.Println(ticker.Symbol + " - " + ticker.Price)
			}
		} else {
			fmt.Println("Too many arguments")
		}

	},
}

func init() {
	RootCmd.AddCommand(tickerCmd)
}
