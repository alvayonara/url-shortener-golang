package shortener

import (
	"math/rand"
	"time"
	"url-shortener-golang/internal/store"
	"url-shortener-golang/models"
)

type Service struct {
	store store.LinkStore
}

func NewService(store store.LinkStore) *Service {
	rand.Seed(time.Now().UnixNano())
	return &Service{store: store}
}

const (
	codeLength = 6
	charset    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func generateCode() string {
	b := make([]byte, codeLength)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func (s *Service) Create(url string) (models.Link, error) {
	for {
		code := generateCode()
		if _, exists := s.store.Get(code); exists {
			continue
		}
		link := models.Link{
			Code: code,
			URL:  url,
		}
		if err := s.store.Save(link); err != nil {
			return models.Link{}, err
		}
		return link, nil
	}
}
