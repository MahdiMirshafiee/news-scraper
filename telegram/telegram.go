package telegram

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/MahdiMirshafiee/news-scraper/scraper"
)

func SendPost(items []scraper.News) error {
	botToken := os.Getenv("BOT_TOKEN")
	chatID := "@TopTenHackerNews"

	if botToken == "" || chatID == "" {
		log.Fatal("BOT_TOKEN or CHAT_ID not set")
	}

	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

	var b strings.Builder
	b.WriteString("🔥 Top 10 Hacker News today:\n\n")
	for i, item := range items {
		b.WriteString(fmt.Sprintf("%d. %s\n%s\n\n", i+1, item.Title, item.Link))
	}
	b.WriteString("@TopTenHackerNews\n")
	message := b.String()
	resp, err := http.PostForm(apiURL, url.Values{
		"chat_id": {chatID},
		"text":   {message},
	})

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("telegram API returned status: %s", resp.Status)
	}
	return nil
}