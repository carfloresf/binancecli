package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// Ticker is used to output a standard structure (marshalled as JSON)
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
		var errdo error

		client := &http.Client{}
		req, errq := http.NewRequest("GET", "https://api.binance.com/api/v3/ticker/price", nil)
		if errq != nil {
			log.Fatalf("Problems: %s", errq)
		}
		q := req.URL.Query()

		if len(args) == 0 {
			resp, errdo = client.Do(req)
			if errdo != nil {
				fmt.Printf("The HTTP request failed with error %s\n", errdo)
			} else {
				data, errio := ioutil.ReadAll(resp.Body)
				if errio != nil {
					log.Fatalf("Problems: %s", errio)
				}
				var tickers []Ticker
				erru := json.Unmarshal(data, &tickers)
				if erru != nil {
					log.Fatalf("Problems: %s", erru.Error())
				}

				fmt.Println("Found", len(tickers), "pairs in Binance")

				for _, ti := range tickers {
					fmt.Println(ti.Symbol + " - " + ti.Price)
				}
			}
		} else if len(args) == 1 {
			q.Add("symbol", args[0])
			req.URL.RawQuery = q.Encode()
			resp, errdo = client.Do(req)
			if errdo != nil {
				log.Fatalf("Problems: %s", errdo.Error())
			} else {
				data, errio := ioutil.ReadAll(resp.Body)
				if errio != nil {
					log.Fatalf("Problems: %s", errio)
				}
				var ticker Ticker
				erru := json.Unmarshal(data, &ticker)
				if erru != nil {
					log.Fatalf("Problems unmarshaling: %s", erru.Error())
				}
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
