package entities

import (
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gustavoteixeira8/url-shortener/src/entities/validations"
	"github.com/gustavoteixeira8/url-shortener/src/utils"
)

type URLShort struct {
	ID             string    `json:"id" gorm:"primaryKey;type:varchar"`
	Name           string    `json:"name" gorm:"unique;type:varchar"`
	URL            string    `json:"url" gorm:"not null;type:varchar"`
	NumberOfClicks uint      `json:"numberOfClicks" gorm:"type:int"`
	CreatedAt      time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updatedAt" gorm:"autoCreateTime"`
}

func NewURLShort(urlShort *URLShort) (*URLShort, error) {
	err := urlShort.SetName(urlShort.Name)

	if err != nil {
		return nil, err
	}

	err = urlShort.SetURL(urlShort.URL)

	if err != nil {
		return nil, err
	}

	if urlShort.ID == "" {
		urlShort.ID = uuid.NewString()
		urlShort.CreatedAt = time.Now()
		urlShort.UpdatedAt = time.Now()
		urlShort.NumberOfClicks = 0
	}

	return urlShort, nil
}

func (u *URLShort) SetName(name string) error {
	if name != "" {
		if err := validations.ValidateName(name); err != nil {
			return err
		}
	} else {
		name = utils.NewID()
	}

	u.Name = FormatName(name)
	u.UpdatedAt = time.Now()

	return nil
}

func (u *URLShort) SetURL(newurl string) error {
	newurlParsed, err := validations.ValidateURL(newurl)
	if err != nil {
		return err
	}

	u.URL = newurlParsed
	u.UpdatedAt = time.Now()

	return nil
}

func (u *URLShort) AddClick() {
	u.NumberOfClicks += 1
	u.UpdatedAt = time.Now()
}

func FormatName(name string) string {
	name = strings.Trim(name, " ")
	name = strings.ReplaceAll(name, " ", "-")
	name = strings.ReplaceAll(name, "/", "-")
	name = strings.ReplaceAll(name, "\\", "-")
	name = url.PathEscape(name)
	return name
}
