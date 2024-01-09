// news_aggregator.go
package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/Anisulh/content_personalization/models"
	"github.com/IBM/sarama"
)

// Struct to parse response from NewsAPI
type NewsAPIResponse struct {
	Articles []struct {
		Source struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"source"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Url         string `json:"url"`
		UrlToImage  string `json:"urlToImage"`
		PublishedAt string `json:"publishedAt"`
		Content     string `json:"content"`
	} `json:"articles"`
}

type MyData struct {
	FirstElement string
	SecondElement []string
}

func CallNLP(text string) (MyData, error) {
	// function that calls the nlp via grpc sending the content and recieving an array ['category', ['keywords']]
	return MyData{FirstElement: "", SecondElement: []string{"", "", ""}}, nil
}

func produceToKafka(topic string, message []byte, brokers []string) error {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
			return err
	}
	defer producer.Close()

	_, _, err = producer.SendMessage(&sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder(message),
	})
	return err
}

func FetchNews() ([]models.Content, error) {

	url := fmt.Sprintf("https://newsapi.org/v2/top-headlines?country=us&apiKey=%v", NewsAPIKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiResp NewsAPIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, err
	}

	var contents []models.Content
	for _, article := range apiResp.Articles {
		val, err := CallNLP(article.Content)
		if err != nil {
			return nil, err
		}
		fmt.Print(val)
		

		// Serialize content for Kafka
		contentBytes, err := json.Marshal(models.Content{
			Title:       article.Title,
			Description: article.Description,
			URL:         article.Url,
			ImageURL:    article.UrlToImage,
			PublishedAt: article.PublishedAt,
			Source:      article.Source.Id,
			Category:    val.FirstElement,
			Keywords:    val.SecondElement,
	})
	if err != nil {
			return nil, err
	}
	// Produce the message to Kafka
	err = produceToKafka("news_topic", contentBytes, []string{"localhost:9092"})
	if err != nil {
			log.Printf("Failed to produce message to Kafka: %s", err)
	}
		
	}
	// return to datastorage function to save to db
	return contents, nil
}


func StartScheduledFetching() {
	ticker := time.NewTicker(24 * time.Hour) // Set the duration according to your needs

	go func() {
			for {
					select {
					case <-ticker.C:
							_, err := FetchNews()
							if err != nil {
									log.Printf("Error fetching news: %s", err)
							} else {
									log.Println("Successfully fetched news")
							}
					}
			}
	}()
}

