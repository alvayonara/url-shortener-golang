package store

import "url-shortener-golang/models"

type LinkStore interface {
	Save(link models.Link) error
	Get(code string) (models.Link, error)
}
