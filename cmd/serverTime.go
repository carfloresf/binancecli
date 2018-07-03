package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

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
			data, _ := ioutil.ReadAll(resp.Body)

			var responseTime ResponseTime
			json.Unmarshal(data, &responseTime)

			fmt.Println(responseTime.ServerTime)
		}
	},
}

func init() {
	RootCmd.AddCommand(timeCmd)
}
