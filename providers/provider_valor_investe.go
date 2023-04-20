package providers

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/msantosfelipe/stock-news-finder/config"
)

const (
	ValorInvesteCode ProviderCode = "ValorInveste"
	valorInvesteName string       = "Valor Investe"
)

func addValorInveste(providersMap map[ProviderCode]Provider) {
	providersMap[ValorInvesteCode] = Provider{
		Code:     string(ValorInvesteCode),
		Name:     valorInvesteName,
		BaseUrl:  config.ENV.UrlValorInveste,
		Function: getFromValorInveste,
	}
}

func getFromValorInveste(stockInput, baseUrl string) {
	baseUrl = fmt.Sprintf("%s/busca?q=%s&page=1", baseUrl, stockInput)
	doc := getNewsPage(baseUrl)
	findFromValorInveste(doc)
}

func findFromValorInveste(doc *goquery.Document) {
	doc.Find(".widget--info__text-container").Each(func(i int, s *goquery.Selection) {
		var link string
		link, ok := s.Find("a").Attr("href")
		if !ok {
			link = "error"
		}
		title := strings.TrimSpace(s.Find(".widget--info__title").Text())
		date := strings.TrimSpace(s.Find(".widget--info__meta").Text())
		var news = News{
			Link:  link,
			Title: title,
			Date:  date,
		}
		fmt.Printf("%d. %s\n%s\n%s\n\n", i+1, news.Title, news.Link, news.Date)
	})
}
