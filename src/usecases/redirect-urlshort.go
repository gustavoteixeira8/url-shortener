package usecases

import (
	"time"

	"github.com/gustavoteixeira8/url-shortener/src/cache"
	"github.com/gustavoteixeira8/url-shortener/src/cerrors"
	"github.com/gustavoteixeira8/url-shortener/src/entities"
	"github.com/gustavoteixeira8/url-shortener/src/repositories"
	"github.com/sirupsen/logrus"
)

type redirectURLShortUseCases struct {
	cache              *cache.AppCache[entities.URLShort]
	urlShortRepository *repositories.URLShortRepository
}

type RedirectURLShortRequest struct {
	Name string `json:"name"`
}

func (u redirectURLShortUseCases) Exec(req RedirectURLShortRequest) (*entities.URLShort, error) {
	if req.Name == "" {
		return nil, cerrors.ErrNameIsRequired
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
				logrus.Errorf("Error counting click in psql to URL %s (%v)", req.Name, err)
			}
			err = u.cache.Set(req.Name, *urlShort, time.Hour*24)
			if err != nil {
				logrus.Errorf("Error counting click in redis to URL %s (%v)", req.Name, err)
			}
		}
	}

	defer func() {
		go countClickFn()
	}()

	urlShort, err = u.cache.Get(req.Name)
	if err == nil && urlShort != nil {
		return urlShort, nil
	}

	urlShort, err = u.urlShortRepository.FindByName(req.Name)
	if err == nil && urlShort != nil {
		u.cache.Set(req.Name, *urlShort, time.Hour*24)
		return urlShort, nil
	}

	return nil, cerrors.ErrUrlNotFound
}

func NewRedirectURLShortUseCases() *redirectURLShortUseCases {
	return &redirectURLShortUseCases{
		cache:              cache.NewAppCache[entities.URLShort](),
		urlShortRepository: repositories.NewURLShortRepository(),
	}
}
