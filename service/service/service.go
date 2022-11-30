package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"pricing-system-alert-service/domain"
	"time"
)

type CoinMarketCapApiListingResp struct {
	Status struct {
		TimeStamp time.Time `json:"timestamp"`
		Error     string    `json:"error_message"`
	} `json:"status"`
	Data []struct {
		Symbol string `json:"symbol"`
		Quote  struct {
			USD struct {
				Price float64 `json:"price"`
			} `json:"USD"`
		} `json:"quote"`
	} `json:"data"`
}

func FetchAndSaveData(db domain.DataBase) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
	if err != nil {
		return err
	}

	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "4")
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", "08b53f4d-0a5d-452e-b67f-3af8af0d035c")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request to: %s,\nreturted unexpected status: %s", req.URL.String(), resp.Status)
	}

	respData := CoinMarketCapApiListingResp{}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return err
	}

	if len(respData.Status.Error) != 0 {
		return fmt.Errorf("error responce from %s,\nmessage: %s", req.URL.String(), respData.Status.Error)
	}

	db.AddNoteBTC(&domain.PriceNote{
		TimeStamp: respData.Status.TimeStamp,
		Currency:  respData.Data[0].Symbol,          // BTC always on index 0
		Price:     respData.Data[0].Quote.USD.Price, // BTC always on index 0
	})

	db.AddNoteETH(&domain.PriceNote{
		TimeStamp: respData.Status.TimeStamp,
		Currency:  respData.Data[1].Symbol,          // ETH always on index 1
		Price:     respData.Data[1].Quote.USD.Price, // ETH always on index 1
	})

	db.AddNoteBNB(&domain.PriceNote{
		TimeStamp: respData.Status.TimeStamp,
		Currency:  respData.Data[3].Symbol,          // BNB always on index 3
		Price:     respData.Data[3].Quote.USD.Price, // BNB always on index 3
	})

	return nil
}

func ReadAndSendLastData(alertChan chan *domain.PriceNote, db domain.DataBase) error {
	btcVal, err := db.GetLastNoteBTC()
	if err != nil {
		return err
	}
	alertChan <- btcVal

	ethVal, err := db.GetLastNoteETH()
	if err != nil {
		return err
	}
	alertChan <- ethVal

	bnbVal, err := db.GetLastNoteBNB()
	if err != nil {
		return err
	}
	alertChan <- bnbVal

	return nil
}
