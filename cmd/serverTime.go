package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// ResponseTime is used to output a standard structure (marshalled as JSON)
type ResponseTime struct {
	ServerTime int `serverTime:"number"`
}

var timeCmd = &cobra.Command{
	Use:   "serverTime",
	Short: "Get ServerTime from Binance",
	Run: func(cmd *cobra.Command, args []string) {
		resp, er := http.Get("https://api.binance.com/api/v1/time")
		if er != nil {
			fmt.Printf("The HTTP request failed with error %s\n", er)
		} else {
			data, errio := ioutil.ReadAll(resp.Body)
			if errio != nil {
				log.Fatalf("Problems: %s", errio.Error())
			}
			var responseTime ResponseTime
			err := json.Unmarshal(data, &responseTime)
			if err != nil {
				log.Fatalf("Problems unmarshaling response: %s", err.Error())
			}

			fmt.Println(responseTime.ServerTime)
		}
	},
}

func init() {
	RootCmd.AddCommand(timeCmd)
}
