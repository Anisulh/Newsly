# Newsly - Personalized News Recommendation Engine

![Go Version](https://img.shields.io/badge/Go-1.16-blue.svg)
![Fiber Framework](https://img.shields.io/badge/fiber-v2.x-brightgreen.svg)
![Kafka](https://img.shields.io/badge/Kafka-2.8.0-blue.svg)
![NLP with BERT](https://img.shields.io/badge/NLP-BERT-orange.svg)
![TailwindCSS](https://img.shields.io/badge/TailwindCSS-2.2.x-06B6D4.svg)

Welcome to the repository for Newsly, a news recommendation engine inspired by TikTok's "For You" page. Newsly leverages a backend written in Go, utilizing the Fiber framework for its lightweight and efficient server capabilities, Templ for templating, and HTMX for dynamic content updates.

If you haven't already, check out the Newsly NLP Service: [https://github.com/Anisulh/NewlyNLP](https://github.com/Anisulh/NewlyNLP)

## Project Overview

Newsly aims to deliver a seamless and highly personalized news reading experience. By integrating advanced NLP techniques using a custom-tuned BERT model, the system categorizes and tags incoming news articles which are then recommended based on user interactions such as likes, dislikes, and saves. Here’s how it works:

- **News Fetching**: A cron job regularly fetches news articles from various sources via the News API.
- **Processing via Kafka**: These articles are sent to an NLP service through Kafka, where they are analyzed and categorized.
- **User Interaction Tracking**: User preferences are dynamically updated based on their interactions with the articles.
- **Continuous Learning**: The recommendation algorithm adjusts to user preferences, providing increasingly accurate article suggestions.
- **Infinite Scrolling**: New articles load as the user scrolls, ensuring an endless stream of personalized content.

## Technologies

- **Backend**: Written in Go with the Fiber framework for optimal performance.
- **Frontend**: TailwindCSS for styling and HTMX for partial HTML updates without reloading.
- **NLP**: Utilizes a fine-tuned BERT model for natural language processing to categorize and tag news articles.
- **Messaging**: Apache Kafka handles message passing between the news fetching cron job and the NLP service.
- **Authentication**: JWT-based authentication secures user accounts and preferences.

## Getting Started

Follow these instructions to get Newsly up and running on your local machine for development and testing purposes.

### Prerequisites

- Go 1.16+
- Docker (for running Kafka)
- Access to a Kafka instance
- Node.js and npm (for TailwindCSS)

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/anisulh/newsly.git
   ```

2. **Navigate to the project directory:**

   ```bash
   cd newsly
   ```

3. **Install frontend dependencies:**

   ```bash
   npm install
   ```

4. **Start the Kafka and Zookeeper services:**

   ```bash
   docker-compose up -d
   ```

5. **Run the backend server:**

   ```bash
   go run main.go
   ```

6. **Visit the local app:**
   
   Open your browser and go to `http://localhost:3000` to see Newsly in action.

## Contributing

Your contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## Contact

The best way to get in touch with my is to contact me via email: anisulhaque9391@gmail.com
