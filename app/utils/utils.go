package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/Fi44er/ton-backend/utils/response"
)

func getTransaction(hash string) (response.Transaction, error) {
	url := "https://tonapi.io/v2/traces/" + hash

	var data response.Transaction
	res, err := http.Get(url)
	if err != nil {
		return data, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func SearchValidTransaction(boc string) (response.Transaction, error) {
	data, err := getTransaction(boc)
	if err != nil {
		return data, err
	}

	if data.Hash == "" {
		time.Sleep(3 * time.Minute)

		data, err = getTransaction(boc)
		if err != nil {
			return data, err
		}

		if data.Hash == "" {
			return data, errors.New("data is empty after retry")
		}
	}

	return data, nil
}
