package auth

import (
	"encoding/json"
	"os"
)

type APIKey struct {
	Name       string `json:"name"`
	PrivateKey string `json:"privateKey"`
}

func LoadAPIKeyFromFile(fileName string) (APIKey, error) {
	var key APIKey
	file, err := os.Open(fileName)
	if err != nil {
		return key, err
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&key)
	return key, err
}
