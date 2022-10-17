package services

import (
	"encoding/json"
	"fetch-app/models"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func callRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}

func newRequestCommodity() (*http.Request, error) {
	r, err := http.NewRequest(
		http.MethodGet,
		urlRequestCommodity(),
		nil,
	)

	return r, err
}

func urlRequestCommodity() string {
	return "https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list"
}

func requestCommodity() ([]*models.Commodity, error) {
	req, err := newRequestCommodity()
	if err != nil {
		return nil, err
	}

	res, err := callRequest(req)
	if err != nil {
		return nil, err
	}

	cmds := []*models.Commodity{}
	err = json.Unmarshal(res, &cmds)
	if err != nil {
		return nil, err
	}

	return cmds, nil
}

func newRequestCurrency(to, from string, amount float64) (*http.Request, error) {
	r, err := http.NewRequest(
		http.MethodGet,
		urlRequestCurrency(to, from, amount),
		nil,
	)
	r.Header.Set("apikey", os.Getenv("API_KEY"))

	return r, err
}

func urlRequestCurrency(to, from string, amount float64) string {
	return fmt.Sprintf(
		"https://api.apilayer.com/fixer/convert?to=%s&from=%s&amount=%f",
		to,
		from,
		amount,
	)
}

func requestCurrency(to, from string, amount float64) (*models.CurrencyResponse, error) {
	req, err := newRequestCurrency(to, from, amount)
	if err != nil {
		return nil, err
	}

	res, err := callRequest(req)
	if err != nil {
		return nil, err
	}

	cr := models.CurrencyResponse{}
	err = json.Unmarshal(res, &cr)
	if err != nil {
		return nil, err
	}

	return &cr, nil
}
