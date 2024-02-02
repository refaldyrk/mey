package model

type ResponseSend struct {
	Data struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Code string `json:"code"`
	} `json:"data"`
}

type ResponseGet struct {
	Data string `json:"data"`
	Name string `json:"name"`
}
