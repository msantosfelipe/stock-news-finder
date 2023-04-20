package main

import (
	"fmt"
	"strings"

	"github.com/msantosfelipe/stock-news-finder/providers"
)

func main() {
	providersMap := providers.NewProviders()

	fmt.Println("Type the stock code: (ex: 'ITSA4')")
	var stockInput string
	_, _ = fmt.Scanln(&stockInput)

	stockInput = strings.TrimSpace(stockInput)
	for _, provider := range providersMap {
		fmt.Println("Executing provider", provider.Name)
		provider.Function(stockInput, provider.BaseUrl)
	}
}
