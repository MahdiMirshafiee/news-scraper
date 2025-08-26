package telegram

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"log"
	"strings"

	"github.com/MahdiMirshafiee/news-scraper/scraper"
)

func SendPost(items []scraper.News) error {
	botToken := os.Getenv("BOT_TOKEN")
	chatID := os.Getenv("CHAT_ID")

	if botToken == "" || chatID == "" {
		log.Fatal("BOT_TOKEN or CHAT_ID not set")
	}

	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

	var b strings.Builder
	b.WriteString("🔥 Top 10 Hacker News today:\n\n")
	for i, item := range items {
		b.WriteString(fmt.Sprintf("%d. %s\n%s\n\n", i+1, item.Title, item.Link))
	}
	message := b.String()
	resp, err := http.PostForm(apiURL, url.Values{
		"chatId": {chatID},
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