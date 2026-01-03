package wordService

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
)

type Service struct {
	startWords []string
}

func NewService(startWordsPath string) *Service {
	startWords, err := loadStartWords(startWordsPath)
	if err != nil {
		log.Fatalf("Failed to load start words: %v", err)
	}

	log.Printf("Loaded %d start words", len(startWords))

	return &Service{
		startWords: startWords,
	}
}

func loadStartWords(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var startWords []string
	err = json.Unmarshal(data, &startWords)
	if err != nil {
		return nil, err
	}

	return startWords, nil
}

// GetRandomStartWord returns a random start word
func (s *Service) GetRandomStartWord() string {
	return s.startWords[rand.Intn(len(s.startWords))]
}
