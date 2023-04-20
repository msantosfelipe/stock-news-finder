package providers

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type ProviderCode string

type Provider struct {
	Code     string
	Name     string
	BaseUrl  string
	Function func(stockInput, baseUrl string)
}

type News struct {
	Title string
	Link  string
	Date  string
}

func NewProviders() map[ProviderCode]Provider {
	providersMap := make(map[ProviderCode]Provider)

	addValorInveste(providersMap)

	return providersMap
}

func getNewsPage(baseUrl string) *goquery.Document {
	resp, err := http.Get(baseUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return doc
}
