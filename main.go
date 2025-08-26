package main

import (
	"log"

	"github.com/MahdiMirshafiee/news-scraper/scraper"
	"github.com/MahdiMirshafiee/news-scraper/telegram"
	"github.com/joho/godotenv"
)
func main (){
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env not found; continuing with environment variables")
	}

	news,err := scraper.FetchTopNews()
	if err != nil {
		log.Fatalf("Error fetching news: %v", err)
	}

	log.Printf("Fetched %d news items", len(news))

	if len(news) == 0 {
		log.Println("No news items found. Exiting.")
		return
	}
	if err := telegram.SendPost(news); err != nil {
		log.Fatalf("Error sending news to Telegram: %v", err)
	}
	log.Println("News successfully sent to Telegram ✅")
}