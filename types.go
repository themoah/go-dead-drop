package main

//Configuration blab blah
type Configuration struct {
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

//StoredSecret blah blah
type StoredSecret struct {
	basePath string
	filename string
}
