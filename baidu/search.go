package baidu

import (
	"strings"
	"time"

	"github.com/karust/openserp/core"
	"github.com/sirupsen/logrus"
)

type Baidu struct {
	core.Browser
	checkTimeout time.Duration
}

func New(browser core.Browser) *Baidu {
	baid := Baidu{Browser: browser}
	baid.checkTimeout = time.Second * 2
	return &baid
}
func (baid *Baidu) Name() string {
	return "baidu"
}

func (baid *Baidu) Search(query core.Query) ([]core.SearchResult, error) {
	logrus.Tracef("Start Baidu search, query: %+v", query)

	searchResults := []core.SearchResult{}

	// Build URL from query struct to open in browser
	url, err := BuildURL(query)
	if err != nil {
		return nil, err
	}

	page := baid.Navigate(url)

	results, err := page.Timeout(baid.Timeout).Search("div.c-container.new-pmd")
	if err != nil {
		return nil, err
	}

	resultElements, err := results.All()
	if err != nil {
		return nil, err
	}

	for i, r := range resultElements {
		// Get URL
		link, err := r.Element("a")
		if err != nil {
			continue
		}
		linkText, err := link.Property("href")
		if err != nil {
			logrus.Error("No `href` tag found")
		}

		// Get title
		title, err := link.Text()
		if err != nil {
			logrus.Error("Cannot extract text from title")
			title = "No title"
		}

		// Get description
		desc, err := r.Text()
		if err != nil {
			desc = ""
		}
		desc = strings.ReplaceAll(desc, title, "")

		gR := core.SearchResult{Rank: i + 1, URL: linkText.String(), Title: title, Description: desc}
		searchResults = append(searchResults, gR)
	}

	if !baid.LeavePageOpen {
		err = page.Close()
		if err != nil {
			logrus.Error(err)
		}
	}

	return searchResults, nil
}