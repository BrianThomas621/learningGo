package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//apiUrl = "https://coinmarketcap-nexuist.rhcloud.com/api/"
//ethApi = apiUrl + "eth"
//btcApi = apiUrl + "btc"

func main() {
	r, _ := http.Get("https://coinmarketcap-nexuist.rhcloud.com/api/eth")
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	var data interface{}
	json.Unmarshal(body, &data)
	// Generate your Go Structures from JSON data by
	// using this link: http://json2struct.mervine.net/
	type ApiReturn struct {
		Change    string `json:"change"`
		MarketCap struct {
			Aud float64 `json:"aud"`
			Btc float64 `json:"btc"`
			Cad float64 `json:"cad"`
			Cny float64 `json:"cny"`
			Eur float64 `json:"eur"`
			Gbp float64 `json:"gbp"`
			Hkd float64 `json:"hkd"`
			Jpy float64 `json:"jpy"`
			Rub float64 `json:"rub"`
			Usd float64 `json:"usd"`
		} `json:"market_cap"`
		Name     string `json:"name"`
		Position string `json:"position"`
		Price    struct {
			Aud float64 `json:"aud"`
			Btc float64 `json:"btc"`
			Cad float64 `json:"cad"`
			Cny float64 `json:"cny"`
			Eur float64 `json:"eur"`
			Gbp float64 `json:"gbp"`
			Hkd float64 `json:"hkd"`
			Jpy float64 `json:"jpy"`
			Rub float64 `json:"rub"`
			Usd float64 `json:"usd"`
		} `json:"price"`
		Supply    string `json:"supply"`
		Symbol    string `json:"symbol"`
		Timestamp string `json:"timestamp"`
		Volume    struct {
			Aud float64 `json:"aud"`
			Btc float64 `json:"btc"`
			Cad float64 `json:"cad"`
			Cny float64 `json:"cny"`
			Eur float64 `json:"eur"`
			Gbp float64 `json:"gbp"`
			Hkd float64 `json:"hkd"`
			Jpy float64 `json:"jpy"`
			Rub float64 `json:"rub"`
			Usd int     `json:"usd"`
		} `json:"volume"`
	}

	var n ApiReturn
	json.Unmarshal(body, &n)
	fmt.Printf("1 ETH = %f USD\n", n.Price.Usd)
}
