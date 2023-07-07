package dto

import "net/http"

type AddShortenDto struct {
	Url       string `json:"url" validate:"http_url,required"`
	CustomUrl string `json:"custom_url"`
	Exp       int32  `json:"exp"`
}

func (a *AddShortenDto) Bind(r *http.Request) error {
	return validate.Struct(a)
}
