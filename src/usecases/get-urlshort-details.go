package usecases

import (
	"time"

	"github.com/gustavoteixeira8/url-shortener/src/cache"
	"github.com/gustavoteixeira8/url-shortener/src/cerrors"
	"github.com/gustavoteixeira8/url-shortener/src/entities"
	"github.com/gustavoteixeira8/url-shortener/src/repositories"
)

type getURLShortDetailsUseCase struct {
	cache              *cache.AppCache[entities.URLShort]
	urlShortRepository *repositories.URLShortRepository
}

type GetURLShortDetailsRequest struct {
	Name string `json:"name"`
}

func (u getURLShortDetailsUseCase) Exec(req GetURLShortDetailsRequest) (*entities.URLShort, error) {
	if req.Name == "" {
		return nil, cerrors.ErrNameIsRequired
	}

	var (
		err      error
		urlShort *entities.URLShort
	)

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

func NewGetURLShortDetailsUseCase() *getURLShortDetailsUseCase {
	return &getURLShortDetailsUseCase{
		cache:              cache.NewAppCache[entities.URLShort](),
		urlShortRepository: repositories.NewURLShortRepository(),
	}
}
