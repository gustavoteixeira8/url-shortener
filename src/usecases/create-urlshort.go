package usecases

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

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
		return nil, errors.New("url is required")
	}

	// ping na URL para verificar se ele existe

	purl, err := url.ParseRequestURI(req.URL)

	if err != nil {
		return nil, err
	}

	hostWithProtocol := fmt.Sprintf("https://%s", purl.Host)

	resp, err := http.Get(hostWithProtocol)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 399 {
		return nil, fmt.Errorf("something when wrong with this url (%s)", req.URL)
	}

	nameExists := u.urlShortRepository.ExistsWithName(req.Name)
	if nameExists {
		return nil, errors.New("name already exists")
	}

	urlExists := u.urlShortRepository.ExistsWithURL(req.URL)
	if urlExists {
		return nil, errors.New("url already exists")
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
