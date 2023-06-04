package usecases

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gustavoteixeira8/url-shortener/src/cerrors"
	"github.com/gustavoteixeira8/url-shortener/src/entities"
	"github.com/gustavoteixeira8/url-shortener/src/repositories"
)

type createUrlShortUserCase struct {
	urlShortRepository *repositories.URLShortRepository
}

type CreateUrlShortRequest struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (u *createUrlShortUserCase) Exec(req *CreateUrlShortRequest) (*entities.URLShort, error) {
	if req.URL == "" {
		return nil, cerrors.ErrUrlIsRequired
	}

	// ping na URL para verificar se ele existe

	purl, err := url.ParseRequestURI(req.URL)

	if err != nil {
		return nil, cerrors.ErrInvalidURLFormat
	}

	hostWithProtocol := fmt.Sprintf("https://%s%s", purl.Host, purl.Path)

	resp, err := http.Get(hostWithProtocol)
	if err != nil {
		return nil, cerrors.ErrCantPingUrl
	}

	if resp.StatusCode > 399 {
		return nil, cerrors.ErrCantPingUrl
	}

	nameExists := u.urlShortRepository.ExistsWithName(entities.FormatName(req.Name))
	if nameExists {
		return nil, cerrors.ErrNameAlreadyExists
	}

	urlExists := u.urlShortRepository.ExistsWithURL(req.URL)
	if urlExists {
		return nil, cerrors.ErrUrlAlreadyExists
	}

	urlShort, err := entities.NewURLShort(&entities.URLShort{
		Name: req.Name,
		URL:  req.URL,
	})

	if err != nil {
		return nil, err
	}

	err = u.urlShortRepository.Save(urlShort)

	if err != nil {
		return nil, err
	}

	urlShort, err = u.urlShortRepository.FindByName(urlShort.Name)

	return urlShort, err
}

func NewCreateUrlShortUseCase() *createUrlShortUserCase {
	return &createUrlShortUserCase{
		urlShortRepository: repositories.NewURLShortRepository(),
	}
}
