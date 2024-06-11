// news_aggregator.go
package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Struct to parse response from NewsAPI
type NewsAPIResponse struct {
	Articles []Article `json:"articles"`
}
type Article struct {
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
}

func FetchNews(NewsAPIKey string) (NewsAPIResponse, error) {
	log.Println("Fetching news")
	url := fmt.Sprintf("https://newsapi.org/v2/top-headlines?country=us&apiKey=%v", NewsAPIKey)
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return NewsAPIResponse{}, err
	}
	defer resp.Body.Close()

	// Check HTTP Response
	if resp.StatusCode != http.StatusOK {
		return NewsAPIResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return NewsAPIResponse{}, err
	}

	var apiResp NewsAPIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return NewsAPIResponse{}, err
	}

	return apiResp, nil
}

// type Content struct {
// 	ID          uint       `json:"ID"`
// 	CreatedAt   time.Time  `json:"CreatedAt"`
// 	UpdatedAt   time.Time  `json:"UpdatedAt"`
// 	DeletedAt   *time.Time `json:"DeletedAt"`
// 	Title       string     `json:"Title"`
// 	Description string     `json:"Description"`
// 	Content     string     `json:"Content"`
// 	URL         string     `json:"URL"`
// 	ImageURL    string     `json:"ImageURL"`
// 	PublishedAt string     `json:"PublishedAt"`
// 	Source      string     `json:"Source"`
// 	Keywords    []string   `json:"Keywords"`
// 	Category    string     `json:"Category"`
// }

// const KafkaTopicNewsInput = "news-input"
// const KafkaConsumerGroup = "news-consumer-group"

// func FetchNews(db *gorm.DB, NewsAPIKey string) error {
// 	log.Println("Fetching news")
// 	url := fmt.Sprintf("https://newsapi.org/v2/top-headlines?country=us&apiKey=%v", NewsAPIKey)
// 	client := &http.Client{
// 		Timeout: 10 * time.Second,
// 	}
// 	resp, err := client.Get(url)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	// Check HTTP Response
// 	if resp.StatusCode != http.StatusOK {
// 		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
// 	}

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return err
// 	}

// 	var apiResp NewsAPIResponse
// 	if err := json.Unmarshal(body, &apiResp); err != nil {
// 		return err
// 	}

// 	// Kafka writer initialization (outside the loop)
// 	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
// 	if err != nil {
// 		return err
// 	}
// 	defer producer.Close()

// 	topic := "news-input"
// 	log.Println("Looping through articles")

// 	for _, article := range apiResp.Articles {
// 		content := models.Content{
// 			Title:       article.Title,
// 			Description: article.Description,
// 			Content:     article.Content,
// 			URL:         article.Url,
// 			ImageURL:    article.UrlToImage,
// 			PublishedAt: article.PublishedAt,
// 			Source:      article.Source.Name,
// 		}

// 		// Check if the content already exists in the db
// 		var existingContent models.Content
// 		log.Printf("Checking for URL: %s", article.Url)
// 		result := db.Model(&models.Content{}).Where("url = ?", article.Url).First(&existingContent)

// 		if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 			log.Printf("Error querying db: %v", result.Error)
// 			continue // Skip this article
// 		}

// 		// If the content does not exist, send it to Kafka
// 		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 			contentBytes, err := json.Marshal(content)
// 			if err != nil {
// 				return err
// 			}

// 			// Kafka production code...
// 			if err := producer.Produce(&kafka.Message{
// 				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
// 				Value:          contentBytes,
// 			}, nil); err != nil {
// 				log.Printf("Error producing Kafka message: %v", err)
// 			}
// 		} else {
// 			log.Println("Article exists in db")
// 		}
// 	}

// 	producer.Flush(15 * 1000)
// 	return nil
// }

// func StartScheduledFetching(db *gorm.DB, NewsAPIKey string) {
// 	ticker := time.NewTicker(24 * time.Hour)
// 	defer ticker.Stop()

// 	// Run the task once immediately and then schedule it to run periodically
// 	fetchAndLogNews(db, NewsAPIKey)

// 	for range ticker.C {
// 			fetchAndLogNews(db, NewsAPIKey)
// 	}
// }

// func fetchAndLogNews(db *gorm.DB, NewsAPIKey string) {
// 	err := FetchNews(db, NewsAPIKey)
// 	if err != nil {
// 			log.Printf("Error fetching news: %s", err)
// 	} else {
// 			log.Println("Successfully fetched news")
// 	}
// }

// func StoreNews(db *gorm.DB) {
// 	log.Println("Starting to consume messages from news-output topic...")

// 	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
// 		"bootstrap.servers": "localhost:9092",
// 		"group.id":          "news-consumer-group",
// 		"auto.offset.reset": "earliest",
// 	})

// 	if err != nil {
// 		log.Fatalf("Failed to create consumer: %v", err)
// 	}
// 	defer consumer.Close()

// 	consumer.SubscribeTopics([]string{"news-output"}, nil)

// 	for {
// 		msg, err := consumer.ReadMessage(-1)
// 		if err != nil {
// 			log.Printf("Error while reading message: %v", err)
// 			continue
// 		}

// 		var content Content
// 		err = json.Unmarshal(msg.Value, &content)
// 		if err != nil {
// 			log.Printf("Error unmarshaling message: %v", err)
// 			continue
// 		}

// 		// Optimized Keyword Handling
// 		existingKeywords := make(map[string]models.Keyword)
// 		db.Find(&existingKeywords)
// 		keywordModels := make([]models.Keyword, 0)
// 		for _, keywordStr := range content.Keywords {
// 			keyword, exists := existingKeywords[keywordStr]
// 			if !exists {
// 				keyword = models.Keyword{Name: keywordStr}
// 				db.Create(&keyword)
// 			}
// 			keywordModels = append(keywordModels, keyword)
// 		}

// 		// Create a new instance of models.Content and populate it
// 		modelsContent := models.Content{
// 			Title:       content.Title,
// 			Description: content.Description,
// 			Content:     content.Content,
// 			URL:         content.URL,
// 			ImageURL:    content.ImageURL,
// 			PublishedAt: content.PublishedAt,
// 			Source:      content.Source,
// 			Category:    models.Category(content.Category), // Convert to your Category type
// 			Keywords:    keywordModels,
// 		}

// 		// Use Transactions for DB operations
// 		tx := db.Begin()
// 		defer func() {
// 			if r := recover(); r != nil {
// 				tx.Rollback()
// 			}
// 		}()
// 		if err := tx.Error; err != nil {
// 			log.Printf("Error starting transaction: %v", err)
// 			return
// 		}

// 		// Check if the content already exists in the db
// 		var existingContent models.Content
// 		if err := tx.Where("url = ?", content.URL).First(&existingContent).Error; err != nil {
// 			if errors.Is(err, gorm.ErrRecordNotFound) {
// 				// URL does not exist, safe to insert
// 				if err := tx.Create(&modelsContent).Error; err != nil {
// 					tx.Rollback()
// 					log.Printf("Error saving content to db: %v", err)
// 					return
// 				}
// 			} else {
// 				// Handle other possible errors
// 				tx.Rollback()
// 				log.Printf("Error querying for existing content: %v", err)
// 				return
// 			}
// 		} else {
// 			// URL already exists, handle accordingly
// 			log.Println("Content with this URL already exists")
// 			// Assuming content contains the new data
// 			if err := tx.Model(&existingContent).Updates(modelsContent).Error; err != nil {
// 				tx.Rollback()
// 				log.Printf("Error updating existing content: %v", err)
// 				return
// 			}
// 		}

// 		if err := tx.Commit().Error; err != nil {
// 			log.Printf("Transaction commit error: %v", err)
// 		} else {
// 			log.Printf("Content processed successfully: %v", modelsContent.ID)
// 		}
// 	}
// }
