package hackernews

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type News struct {
	Title string
	Link  string
}

func FetchTopNews() ([]News, error) {
	url := "https://news.ycombinator.com/"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; NewsScraper/1.0; +https://example.com)")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var list []News

	doc.Find("tr.athing .titleline > a").EachWithBreak(func(i int, s *goquery.Selection) bool {
		title := s.Text()
		link, _ := s.Attr("href")
		if title != "" && link != "" {
			list = append(list, News{Title: title, Link: link})
		}
		return len(list) < 10
	})

	return list, nil
}
