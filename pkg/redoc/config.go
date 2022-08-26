package redoc

type XLogo struct {
	URL             string `json:"url"`
	BackgroundColor string `json:"backgroundColor"`
	AltText         string `json:"altText"`
}
type Config struct {
	Src   string
	Logo  XLogo
	Theme string
}
