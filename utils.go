package main

import (
	"strings"

	"github.com/gocolly/colly"
)

func scrap(url string, selector string) (elements []string, err error) {
	c := colly.NewCollector()
	c.OnHTML(
		selector,
		func(e *colly.HTMLElement) {
			elements = append(elements, strings.TrimSpace(e.Text))
		},
	)
	err = c.Visit(url)
	return
}
