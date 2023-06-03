package usecases

import (
	"errors"
	"time"

	"github.com/gustavoteixeira8/url-shortener/src/cache"
	"github.com/gustavoteixeira8/url-shortener/src/entities"
	"github.com/gustavoteixeira8/url-shortener/src/repositories"
)

type getURLShortDetailsUseCase struct {
	cache              *cache.AppCache[entities.URLShort]
	urlShortRepository *repositories.URLShortRepository
}

type GetURLShortDetailsRequest struct {
	ID string `json:"id"`
}

func (u getURLShortDetailsUseCase) Exec(req GetURLShortDetailsRequest) (*entities.URLShort, error) {
	if req.ID == "" {
		return nil, errors.New("id cannot be empty")
	}

	var (
		err      error
		urlShort *entities.URLShort
	)

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

func NewGetURLShortDetailsUseCase() *getURLShortDetailsUseCase {
	return &getURLShortDetailsUseCase{
		cache:              cache.NewAppCache[entities.URLShort](),
		urlShortRepository: repositories.NewURLShortRepository(),
	}
}
