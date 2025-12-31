package wordService

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
)

type Service struct {
	wordKeys []string
}

func NewService(wordMapPath string) *Service {
	wordMap, err := loadWordMap(wordMapPath)
	if err != nil {
		log.Fatalf("Failed to load word map: %v", err)
	}

	// Pre-compute keys for random selection (we only need the keys, not the values)
	keys := make([]string, 0, len(wordMap))
	for k := range wordMap {
		keys = append(keys, k)
	}

	log.Printf("Loaded %d words for start word selection", len(keys))

	return &Service{
		wordKeys: keys,
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

// GetRandomStartWord returns a random word from the wordmap
func (s *Service) GetRandomStartWord() string {
	return s.wordKeys[rand.Intn(len(s.wordKeys))]
}
