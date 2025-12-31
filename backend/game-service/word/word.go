package word

import (
	"encoding/json"
	"log"
	"os"
	"slices"
)

type Service struct {
	wordMap map[string][]string
}

func NewService(wordMapPath string) *Service {
	wordMap, err := loadWordMap(wordMapPath)
	if err != nil {
		log.Fatalf("Failed to load word map: %v", err)
	}

	log.Printf("Loaded %d words into word map", len(wordMap))

	return &Service{
		wordMap: wordMap,
	}
}

func loadWordMap(path string) (map[string][]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var wordMap map[string][]string
	err = json.Unmarshal(data, &wordMap)
	if err != nil {
		return nil, err
	}

	return wordMap, nil
}

// IsValidMove checks if newWord is a valid move from currentWord
func (s *Service) IsValidMove(currentWord, newWord string) bool {
	validMoves, exists := s.wordMap[currentWord]
	if !exists {
		return false
	}
	return slices.Contains(validMoves, newWord)
}

// IsValidWord checks if a word exists in the wordmap
func (s *Service) IsValidWord(word string) bool {
	_, exists := s.wordMap[word]
	return exists
}
