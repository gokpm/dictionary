package main

import (
	"fmt"
	"strings"
)

func lookupAntonyms(in string) (out string, err error) {
	query, err := lookupCorrection(in)
	if err != nil {
		return
	}
	url := fmt.Sprintf("https://www.thesaurus.com/browse/%[1]s", query)
	elements, err := scrap(url, "a.css-15bafsg.eh475bn0")
	if err != nil {
		return
	}
	for _, element := range elements {
		out += "• " + element + "\n"
	}
	out = strings.TrimSpace(out)
	return
}

func lookupCorrection(in string) (out string, err error) {
	url := fmt.Sprintf("https://www.google.com/search?&q=%[1]s", in)
	elements, err := scrap(url, "i")
	if err != nil {
		return
	}
	if len(elements) > 0 {
		out = elements[0]
	} else {
		out = in
	}
	out = strings.TrimSpace(out)
	return
}

func lookupDefinition(in string) (out string, err error) {
	query, err := lookupCorrection(in)
	if err != nil {
		return
	}
	url := fmt.Sprintf("https://gcide.gnu.org.ua/?q=%[1]s&db=gcide&define=1", query)
	elements, err := scrap(url, "pre")
	if err != nil {
		return
	}
	out = "\n"
	for _, element := range elements {
		out += "• " + element + "\n\n"
	}
	out = strings.TrimSpace(out)
	return
}

func lookupSynonyms(in string) (out string, err error) {
	query, err := lookupCorrection(in)
	if err != nil {
		return
	}
	url := fmt.Sprintf("https://www.thesaurus.com/browse/%[1]s", query)
	elements, err := scrap(url, "a.css-1n6g4vv.eh475bn0")
	if err != nil {
		return
	}
	for _, element := range elements {
		out += "• " + element + "\n"
	}
	out = strings.TrimSpace(out)
	return
}

func lookupUsage(in string) (out string, err error) {
	url := fmt.Sprintf("https://sentence.yourdictionary.com/%[1]s", in)
	elements, err := scrap(url, "p.sentence-item__text")
	if err != nil {
		return
	}
	for _, element := range elements {
		out += "• " + element + "\n\n"
	}
	out = strings.TrimSpace(out)
	return
}

func lookupWordOfTheDay() (out string, err error) {
	url := "https://www.dictionary.com/e/word-of-the-day/"
	elements, err := scrap(url, "body > div.wrapper > div.shell.shell--otd > div.otd-wrapper > div.otd-content.wotd-content > div > div:nth-child(1) > div.otd-item-wrapper-content > div.wotd-item > div > div.otd-item-headword__word > h1")
	if err != nil {
		return
	}
	for _, element := range elements {
		out += element + ": "
	}
	elements, err = scrap(url, "body > div.wrapper > div.shell.shell--otd > div.otd-wrapper > div.otd-content.wotd-content > div > div:nth-child(1) > div.otd-item-wrapper-content > div.wotd-item > div > div.otd-item-headword__pos-blocks > div > p:nth-child(2)")
	if err != nil {
		return
	}
	for _, element := range elements {
		out += element + "\n"
	}
	out = strings.TrimSpace(out)
	return
}
