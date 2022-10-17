package services

import (
	"errors"
	"fetch-app/models"
	"fmt"
	"log"
	"sort"
	"strconv"
	"time"
)

func Aggregate() ([]*models.Aggregate, error) {
	commodities, err := requestCommodity()
	if err != nil {
		return nil, err
	}

	if commodities == nil {
		return nil, errors.New("empty commodities")
	}

	m := make(map[string]models.Aggregate)
	for _, cmdty := range commodities {
		if cmdty == nil || cmdty.ProvincialArea == nil || cmdty.Timestampp == nil {
			continue
		}

		t := stringTimestamppToTime(*cmdty.Timestampp)
		y, w := t.ISOWeek()

		price, err := strconv.ParseFloat(*cmdty.Price, 64)
		if err != nil {
			return nil, err
		}

		size, err := strconv.ParseFloat(*cmdty.Size, 64)
		if err != nil {
			return nil, err
		}

		key := fmt.Sprintf("%s-%d-%d", *cmdty.ProvincialArea, y, w)
		if v, exist := m[key]; exist {
			if price < v.MinPrice {
				v.MinPrice = price
			}
			if price > v.MaxPrice {
				v.MaxPrice = price
			}
			v.TotalPrice += price
			v.ListPrice = append(v.ListPrice, price)

			if size < v.MinSize {
				v.MinPrice = size
			}
			if size > v.MaxSize {
				v.MaxPrice = size
			}
			v.TotalSize += size
			v.ListSize = append(v.ListSize, size)

			continue
		}

		a := models.Aggregate{
			ProvincialArea: *cmdty.ProvincialArea,
			Year:           y,
			Week:           w,
			MinPrice:       price,
			MaxPrice:       price,
			TotalPrice:     price,
			MinSize:        size,
			MaxSize:        size,
			TotalSize:      size,
		}
		a.ListPrice = append(a.ListPrice, price)
		a.ListSize = append(a.ListSize, size)

		m[key] = a
	}

	out := []*models.Aggregate{}
	for _, v := range m {
		v.MeanPrice = mean(v.TotalPrice, v.ListPrice)
		v.MedianPrice = median(v.ListPrice)
		v.MeanSize = mean(v.TotalSize, v.ListSize)
		v.MedianSize = median(v.ListSize)
		o := v
		out = append(out, &o)
	}

	return out, nil
}

func median(data []float64) float64 {
	dataCopy := make([]float64, len(data))
	copy(dataCopy, data)

	sort.Float64s(dataCopy)

	median := float64(0)
	l := len(dataCopy)
	if l == 0 {
		return 0
	} else if l%2 == 0 {
		median = (dataCopy[l/2-1] + dataCopy[l/2]) / 2
	} else {
		median = dataCopy[l/2]
	}

	return median
}

func mean(total float64, data []float64) float64 {
	return total / float64(len(data))
}

func stringTimestamppToTime(timestampp string) time.Time {
	i, err := strconv.ParseFloat(timestampp, 64)
	if err != nil {
		log.Println("error parsing timestampp: ", err.Error())
	}
	return time.Unix(int64(i), 0)
}
