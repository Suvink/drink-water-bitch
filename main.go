package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

type GoogleChatMessage struct {
	Text string `json:"text"`
}

type Config struct {
	WebhookURL  string
	UserID      string
	PhrasesFile string
}

func main() {
	config, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	phrases, err := readPhrases(config.PhrasesFile)
	if err != nil {
		log.Fatal(err)
	}

	if len(phrases) == 0 {
		log.Fatal("no phrases found in file")
	}

	message := formatMessage(config.UserID, selectRandomPhrase(phrases))
	if err := sendToGoogleChat(config.WebhookURL, message); err != nil {
		log.Fatal(err)
	}

	log.Println("Message sent successfully")
}

func loadConfig() (*Config, error) {
	webhookURL := os.Getenv("GOOGLE_CHAT_WEBHOOK")
	if webhookURL == "" {
		return nil, fmt.Errorf("GOOGLE_CHAT_WEBHOOK environment variable not set")
	}

	userID := os.Getenv("USER_ID")
	if userID == "" {
		return nil, fmt.Errorf("USER_ID environment variable not set")
	}

	return &Config{
		WebhookURL:  webhookURL,
		UserID:      userID,
		PhrasesFile: "phrases.txt",
	}, nil
}

func readPhrases(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var phrases []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, ". ", 2)
		if len(parts) == 2 {
			phrases = append(phrases, parts[1])
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return phrases, nil
}

func selectRandomPhrase(phrases []string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return phrases[r.Intn(len(phrases))]
}

func formatMessage(userID, phrase string) string {
	return fmt.Sprintf("Yo <users/%s>,\n%s", userID, phrase)
}

func sendToGoogleChat(webhookURL, message string) error {
	payload := GoogleChatMessage{Text: message}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
