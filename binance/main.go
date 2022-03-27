package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/gen2brain/beeep"
)

var currency = []string{"USDT", "BTC"}

type SearchResult struct {
	Data SearchResultData
}

type SearchResultData = []SearchResultItem

type SearchResultItem struct {
	Adv SearchResultItemAdv
}

type SearchResultItemAdv struct {
	Price string
	TradeMethods []SearchResultItemTradeMethod
}

type SearchResultItemTradeMethod struct {
	Identifier string
}

func main() {
	for {
		err := do()
		if err != nil {
			log.Fatalln("do error:", err)
			break
		}

		time.Sleep(60 * time.Second)
	}
}

func do() error {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	url := "https://www.binance.com/"
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(ctx context.Context) error {
			prices := make(map[string]float64)

			for _, cur := range currency {
				price, err := search(ctx, cur)
				if err != nil {
					log.Fatalln("search error:", err)
					return err
				}
				prices[cur] = price
			}

			message := ""
			for cur, price := range prices {
				message += fmt.Sprintf("%s: %.2f\n", cur, price)
			}

			err := notify("Binance", message)
			if err != nil {
				log.Fatalln("notify error:", err)
				return err
			}

			return err
		}),
	)
	if err != nil {
		log.Fatal("chromedp runner error:", err)
	}

	return err
}

func search(ctx context.Context, asset string) (float64, error) {
	data, err := request(ctx, asset)
	if err != nil {
		log.Fatalln("request error:", err)
		return 0, err
	}

	price, err := find(data)
	if err != nil {
		log.Fatalln("find error:", err)
		return 0, err
	}

	return price, err
}

func request(ctx context.Context, asset string) (SearchResultData, error) {
	payload := bytes.NewBuffer([]byte(fmt.Sprintf(`{
		"asset": "%s",
		"fiat": "RUB",
		"page": 1,
		"payTypes": [
			"A-Bank",
			"ABB Bank",
			"Ak Bars Bank",
			"Asaka Bank",
			"Bank Saint-Petersburg",
			"Bank Transfer (Venezuela)",
			"Bank Transfer",
			"BCS Bank",
			"Citibank (Russia)",
			"Credit Europe Bank (Russia)",
			"Faster Payment System (FPS)",
			"Faster Payments",
			"Home Credit Bank (Russia)",
			"Idram",
			"Monobank",
			"MTS-Bank",
			"OTP",
			"Post Bank",
			"Privat Bank",
			"QIWI",
			"Raiffeisen Bank Aval",
			"Raiffeisenbank",
			"Renaissance Credit Bank",
			"Rosbank",
			"Russian Standard Bank",
			"SettlePay",
			"Tinkoff",
			"Transfers with specific bank",
			"UniCredit",
			"Uralsib Bank",
			"Vostochny Bank",
			"ЮMoney",
		],
		"publisherType": null,
		"rows": 10,
		"tradeType": "BUY",
	}`, asset)))
	url := "https://p2p.binance.com/bapi/c2c/v2/friendly/c2c/adv/search"
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("client error:", err)
		return nil, err
	}
	defer resp.Body.Close()

	var res SearchResult
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		log.Fatalln("decoder error:", err)
		return nil, err
	}

	return res.Data, err
}

func find(data SearchResultData) (float64, error) {
	var min float64
	for index, item := range data {
		price, err := strconv.ParseFloat(item.Adv.Price, 8)
		if err != nil {
			log.Fatalln("parse error:", err)
			return 0, err
		}

		if index == 0 || price < min {
			min = price
		}
	}
	return min, nil
}

func notify(title string, message string) error {
	err := beeep.Beep(440.0, 60)
	if err != nil {
		log.Fatalln("beep error:", err)
		return err
	}

	err = beeep.Notify(title, message, "")
	if err != nil {
		log.Fatalln("notify error:", err)
		return err
	}

	return nil
}
