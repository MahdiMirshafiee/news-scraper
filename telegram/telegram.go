package telegram

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/MahdiMirshafiee/news-scraper/scraper"
)

func SendPost(items []scraper.News) error {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	chatId := os.Getenv("TELEGRAM_CHAT_ID")

	if botToken == "" || chatId == "" {
		return fmt.Errorf("missing TELEGRAM_BOT_TOKEN or TELEGRAM_CHAT_ID environment variables")
	}

	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

	var b strings.Builder
	b.WriteString("🔥 Top 10 Hacker News today:\n\n")
	for i, item := range items {
		b.WriteString(fmt.Sprintf("%d. %s\n%s\n\n", i+1, item.Title, item.Link))
	}
	message := b.String()
	resp, err := http.PostForm(apiURL, url.Values{
		"chatId": {chatId},
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