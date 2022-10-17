package services

import (
	"errors"
	"fetch-app/models"
	"fmt"
	"log"
	"strconv"

	"github.com/patrickmn/go-cache"
)

func FetchList(c *cache.Cache) ([]*models.Commodity, error) {
	commodities, err := requestCommodity()
	if err != nil {
		return nil, err
	}

	if commodities == nil {
		return nil, errors.New("empty commodities")
	}

	return commodities, nil
}

func FetchListUSD(c *cache.Cache) ([]*models.Commodity, error) {
	commodities, err := requestCommodity()
	if err != nil {
		return nil, err
	}

	if commodities == nil {
		return nil, errors.New("empty commodities")
	}

	rate := float64(0)
	for _, cmdty := range commodities {
		if cmdty == nil || cmdty.Price == nil {
			continue
		}

		amnt, err := strconv.ParseFloat(*cmdty.Price, 64)
		if err != nil {
			return nil, err
		}

		if rate == 0 {
			cr, err := requestCurrency("usd", "idr", amnt)
			if err != nil {
				log.Println(err.Error())
				return nil, err
			}

			if cr == nil {
				continue
			}
			rate = cr.Info.Rate
		}

		prcUSD := fmt.Sprintf("%f", amnt*rate)
		cmdty.PriceUSD = &prcUSD
	}

	return commodities, nil
}

func rate(c *cache.Cache, from, to string, amnt float64) (*models.CurrencyResponse, error) {
	data, exist := getCache(c, fmt.Sprintf("%s-%s", from, to))
	if exist {
		cr := data.(*models.CurrencyResponse)
		return cr, nil
	}

	cr, err := requestCurrency("usd", "idr", amnt)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return cr, nil
}
