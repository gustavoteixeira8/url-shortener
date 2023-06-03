package usecases

import (
	"errors"
	"time"

	"github.com/gustavoteixeira8/url-shortener/src/cache"
	"github.com/gustavoteixeira8/url-shortener/src/entities"
	"github.com/gustavoteixeira8/url-shortener/src/repositories"
	"github.com/sirupsen/logrus"
)

type redirectURLShortUseCases struct {
	cache              *cache.AppCache[entities.URLShort]
	urlShortRepository *repositories.URLShortRepository
}

type RedirectURLShortRequest struct {
	ID string `json:"id"`
}

func (u redirectURLShortUseCases) Exec(req RedirectURLShortRequest) (*entities.URLShort, error) {
	if req.ID == "" {
		return nil, errors.New("id cannot be empty")
	}

	var (
		err      error
		urlShort *entities.URLShort
	)

	countClickFn := func() {
		if urlShort != nil && err == nil {
			urlShort.AddClick()
			err = u.urlShortRepository.Save(urlShort)
			if err != nil {
				logrus.Errorf("Error counting click in psql to URL %s (%v)", req.ID, err)
			}
			err = u.cache.Set(req.ID, *urlShort, time.Hour*24)
			if err != nil {
				logrus.Errorf("Error counting click in redis to URL %s (%v)", req.ID, err)
			}
		}
	}

	defer func() {
		go countClickFn()
	}()

	urlShort, err = u.cache.Get(req.ID)
	if err == nil && urlShort != nil {
		return urlShort, nil
	}

	urlShort, err = u.urlShortRepository.FindByID(req.ID)
	if err == nil && urlShort != nil {
		u.cache.Set(req.ID, *urlShort, time.Hour*24)
		return urlShort, nil
	}

	return nil, errors.New("url not found in our database")
}

func NewRedirectURLShortUseCases() *redirectURLShortUseCases {
	return &redirectURLShortUseCases{
		cache:              cache.NewAppCache[entities.URLShort](),
		urlShortRepository: repositories.NewURLShortRepository(),
	}
}
