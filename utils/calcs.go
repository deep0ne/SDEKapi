package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"

	auth "github.com/deep0ne/SDEKapi/authorization"
	"github.com/deep0ne/SDEKapi/data"
)

var (
	FailedNewRequest error = errors.New("Не удалось сформировать запрос")
	FailedDoRequest  error = errors.New("Не удалось сделать запрос")
	FailedReadBody   error = errors.New("Не удалось прочитать тело запроса")

	TestMode *bool
	username string
	password string
	CalcURL  string
)

// формирование body для POST запроса из задания
func FormReqBody(addrFrom, addrTo string, size data.Size) ([]byte, error) {
	reqMap := make(map[string]any, 3)
	reqMap["from_location"] = map[string]string{"address": addrFrom}
	reqMap["to_location"] = map[string]string{"address": addrTo}
	reqMap["packages"] = []map[string]int{{"weight": size.Weight, "height": size.Height, "width": size.Width, "length": size.Length}}
	reqBody, err := json.Marshal(reqMap)
	if err != nil {
		return nil, errors.New("Не удалось сформировать тело запроса")
	}
	return reqBody, nil
}

// Чтобы не делать два запроса (получение токена и данных) друг за другом, вынес в функцию
func GetBody(url, token string, client *http.Client, body io.Reader) ([]byte, error) {
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Fatal(FailedNewRequest)
	}

	if token != "" {
		request.Header.Add("Authorization", token)
		request.Header.Set("Content-Type", "application/json")
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(FailedDoRequest)
	}

	defer response.Body.Close()
	b, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, FailedReadBody
	}
	return b, nil
}

// функция расчёта
func Calculate(addrFrom, addrTo string, size data.Size) ([]data.PriceSending, error) {
	var (
		priceSendings data.PriceSendings
		info          []data.PriceSending
	)

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	authURL := auth.FormURL(username, password, data.ApiURL)
	body, err := GetBody(authURL, "", &client, nil)
	if err != nil {
		return nil, err
	}

	token, err := auth.FormToken(body)
	if err != nil {
		return nil, err
	}

	reqBody, err := FormReqBody(addrFrom, addrTo, size)
	if err != nil {
		return nil, err
	}

	body, err = GetBody(CalcURL, token, &client, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &priceSendings)
	if err != nil {
		return nil, errors.New("Не удалось получить информацию")
	}

	for _, price := range priceSendings.TariffCodes {
		info = append(info, price)
	}
	return info, nil
}

func SetUsername(account string) {
	username = account
}

func SetPassword(pass string) {
	password = pass
}
