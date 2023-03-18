package authorization

import (
	"encoding/json"
	"errors"
	"net/url"
)

// формирование URL с username и паролем
func FormURL(account, password, URL string) string {
	params := url.Values{}
	params.Set("grant_type", "client_credentials")
	params.Set("client_id", account)
	params.Set("client_secret", password)
	return URL + "?" + params.Encode()
}

// Получение токена
func FormToken(body []byte) (string, error) {
	info := make(map[string]any)
	json.Unmarshal(body, &info)
	accessToken, ok := info["access_token"].(string)
	if !ok {
		return "", errors.New("Ошибка при получении токена")
	}
	return "Bearer " + accessToken, nil
}
