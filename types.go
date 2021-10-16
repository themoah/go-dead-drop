package main

type secret struct {
	key  string
	data string
}

type storedSecretResponse struct {
	Key      string `json:"key"`
	Password string `json:"password"`
}

type retrievedDropResponse struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

type statusResponse struct {
	Status string `json:"status"`
}
